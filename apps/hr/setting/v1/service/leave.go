package service

import (
	"context"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/paginator"
	"time"
)

func Leave(ctx context.Context, page ...*paginator.Pagination) ILeave {
	if len(page) != 0 {
		return &leave{ctx: ctx, page: page[0]}
	}
	return &leave{ctx: ctx}
}

type (
	ILeave interface {
		// GetByKeyword 根據keyword, status獲取對應假別
		GetByKeyword(in model.GetByKeywordLeaveReq) (out model.GetByKeywordLeaveRes, err error)
		// GetByID 根據ID獲取假別
		GetByID(in model.GetByIDLeaveReq) (out model.GetByIDLeaveRes, err error)
		// Insert 新增假別
		Insert(in model.PostLeaveReq) (err error)
		// Update 更新假別
		Update(in model.PutLeaveReq) (err error)
		// Delete 刪除假別
		Delete(in model.DeleteLeaveReq) error
		// SetStatus 設定狀態
		SetStatus(in model.SetStatusLeaveReq) (err error)
		// GetLeaveGroup 獲取假別群組列表
		GetLeaveGroup(in model.GetLeaveGroupReq) (out []*model.GetLeaveGroupRes, err error)
		// InsertGroup 新增群組
		InsertGroup(in model.PostLeaveGroupReq) error
		// DeleteGroup 刪除群組
		DeleteGroup(in model.DeleteLeaveGroupReq) error
		// SetLeaveGroupName 設置群組名稱
		SetLeaveGroupName(in model.SetLeaveGroupNameReq) (err error)
		// SetLeaveGroupEmployee 設置群組的員工
		SetLeaveGroupEmployee(in model.SetLeaveGroupEmployeeReq) (err error)
		// SetLeaveGroupCond 設置群組的條件
		SetLeaveGroupCond(in model.SetLeaveGroupCondReq) error
		// ResetEmployeeAvailable 重新計算請假可用天數
		ResetEmployeeAvailable(in model.ResetEmployeeAvailableReq) error
	}

	leave struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword, status獲取對應假別
func (l *leave) GetByKeyword(in model.GetByKeywordLeaveReq) (out model.GetByKeywordLeaveRes, err error) {
	var (
		qLeave  = query.Leave
		qdLeave = qLeave.WithContext(dbcache.WithCtx(l.ctx))
		conds   = []gen.Condition{qdLeave.Where(qLeave.Name.Like("%" + in.Keyword + "%")).Or(qLeave.Code.Like("%" + in.Keyword + "%"))}
		leaves  []*types.Leave
	)
	if in.Status != "" {
		conds = append(conds, qLeave.Status.Is(gbconv.Bool(in.Status)))
	}

	leaves, err = qdLeave.Scopes(l.page.Where(conds...)).Preload(
		qLeave.LeaveGroup,
		qLeave.LeaveGroup.Employee,
		qLeave.LeaveGroup.LeaveGroupCondition,
	).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, leaves); err != nil {
		return
	}

	out.Total = l.page.Total

	return out, err
}

// GetByID 根據ID獲取假別
func (l *leave) GetByID(in model.GetByIDLeaveReq) (out model.GetByIDLeaveRes, err error) {
	var (
		qLeave = query.Leave
	)

	out.Leave, err = qLeave.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeave.ID.Eq(in.ID)).Preload(
		qLeave.LeaveGroup,
		qLeave.LeaveGroup.Employee,
		qLeave.LeaveGroup.LeaveGroupCondition,
	).First()
	return
}

// Insert 新增假別
func (l *leave) Insert(in model.PostLeaveReq) (err error) {
	var (
		qLeave = query.Leave
	)
	in.LeaveGroup = nil
	return qLeave.WithContext(dbcache.WithCtx(l.ctx)).Create(in.Leave)
}

// Update 更新假別
func (l *leave) Update(in model.PutLeaveReq) (err error) {
	var (
		qLeave = query.Leave
	)
	in.Leave.ID = in.ID
	in.LeaveGroup = nil

	return qLeave.WithContext(dbcache.WithCtx(l.ctx)).Save(in.Leave)
}

// Delete 刪除假別
func (l *leave) Delete(in model.DeleteLeaveReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qLeave          = tx.Leave
			qLeaveGroup     = tx.LeaveGroup
			qLeaveGroupCond = tx.LeaveGroupCondition
			qM2M            = tx.LeaveGroupEmployee
			qSignOffSetting = tx.LeaveSignOffSetting
			err             error
		)
		// 刪除假別
		_, err = qLeave.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeave.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 查詢有關連的group id
		var groupIDs = make([]uint, 0)
		err = qLeaveGroup.WithContext(dbcache.WithCtx(l.ctx)).Select(qLeaveGroup.ID).Where(qLeaveGroup.LeaveID.Eq(in.ID)).Scan(&groupIDs)
		if err != nil {
			return err
		}

		// 刪除群組
		_, err = qLeaveGroup.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeaveGroup.LeaveID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除條件
		_, err = qLeaveGroupCond.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeaveGroupCond.LeaveGroupID.In(groupIDs...)).Delete()
		if err != nil {
			return err
		}

		// 刪除多對多表
		_, err = qM2M.WithContext(dbcache.WithCtx(l.ctx)).Where(qM2M.LeaveGroupID.In(groupIDs...)).Delete()
		if err != nil {
			return err
		}

		// 刪除簽核設定表
		_, err = qSignOffSetting.WithContext(dbcache.WithCtx(l.ctx)).Where(qSignOffSetting.LeaveID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetStatus 設定狀態
func (l *leave) SetStatus(in model.SetStatusLeaveReq) (err error) {
	_, err = query.Leave.WithContext(dbcache.WithCtx(l.ctx)).Where(query.Leave.ID.Eq(in.ID)).Update(query.Leave.Status, in.Status)
	return err
}

// GetLeaveGroup 獲取假別群組列表
func (l *leave) GetLeaveGroup(in model.GetLeaveGroupReq) (out []*model.GetLeaveGroupRes, err error) {
	var (
		qLeaveGroup = query.LeaveGroup
		leaveGroups []*types.LeaveGroup
	)
	leaveGroups, err = qLeaveGroup.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeaveGroup.LeaveID.Eq(in.LeaveID)).
		Preload(field.Associations).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out, leaveGroups); err != nil {
		return
	}

	return out, err
}

// InsertGroup 新增群組
func (l *leave) InsertGroup(in model.PostLeaveGroupReq) error {
	var (
		qLeaveGroup = query.LeaveGroup
		err         error
	)

	// 避免從postman或swagger近來的資料污染
	in.Leave = nil
	in.Employee = nil
	in.LeaveGroupCondition = nil

	// 創建leave group
	err = qLeaveGroup.WithContext(dbcache.WithCtx(l.ctx)).Create(in.LeaveGroup)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup 刪除群組
func (l *leave) DeleteGroup(in model.DeleteLeaveGroupReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qLeaveGroup         = tx.LeaveGroup
			qLeaveGroupCond     = tx.LeaveGroupCondition
			qLeaveGroupEmployee = tx.LeaveGroupEmployee
			err                 error
		)
		// 刪除群組
		_, err = qLeaveGroup.WithContext(dbcache.WithCtx(l.ctx)).Unscoped().Where(qLeaveGroup.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除條件表
		_, err = qLeaveGroupCond.WithContext(dbcache.WithCtx(l.ctx)).Unscoped().Where(qLeaveGroupCond.LeaveGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除員工表
		_, err = qLeaveGroupEmployee.WithContext(dbcache.WithCtx(l.ctx)).Unscoped().Where(qLeaveGroupEmployee.LeaveGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetLeaveGroupName 設置群組名稱
func (l *leave) SetLeaveGroupName(in model.SetLeaveGroupNameReq) (err error) {
	_, err = query.LeaveGroup.WithContext(dbcache.WithCtx(l.ctx)).Where(query.LeaveGroup.ID.Eq(in.ID)).Update(query.Leave.Name, in.Name)
	return err
}

// SetLeaveGroupEmployee 設置群組的員工
func (l *leave) SetLeaveGroupEmployee(in model.SetLeaveGroupEmployeeReq) (err error) {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			groupEmployee = make([]*types.LeaveGroupEmployee, 0)
			qM2M          = tx.LeaveGroupEmployee
		)

		for _, employeeID := range in.EmployeeID {
			groupEmployee = append(groupEmployee, &types.LeaveGroupEmployee{LeaveGroupID: in.ID, EmployeeID: employeeID})
		}

		// 先刪除原有員工
		_, err = qM2M.WithContext(dbcache.WithCtx(l.ctx)).Where(qM2M.LeaveGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 建立新的
		err = qM2M.WithContext(dbcache.WithCtx(l.ctx)).Create(groupEmployee...)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetLeaveGroupCond 設置群組的條件
func (l *leave) SetLeaveGroupCond(in model.SetLeaveGroupCondReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qLeaveGroupCond = tx.LeaveGroupCondition
			err             error
		)

		// 刪除原有的條件
		_, err = qLeaveGroupCond.WithContext(dbcache.WithCtx(l.ctx)).Unscoped().Where(qLeaveGroupCond.LeaveGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 遍歷
		for index, cond := range in.Condition {
			cond.ID = 0 // 將ID都變為0, ID將由資料庫自動遞增
			cond.LeaveGroupID = in.ID
			cond.Level = uint(index) + 1
		}

		// 建立新的條件
		err = qLeaveGroupCond.WithContext(dbcache.WithCtx(l.ctx)).Create(in.Condition...)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// ResetEmployeeAvailable 重新計算請假可用天數
func (l *leave) ResetEmployeeAvailable(in model.ResetEmployeeAvailableReq) error {
	var (
		qEmployee = query.Employee
		qLeave    = query.Leave
		employees []*types.Employee
		leave     *types.Leave
		err       error
	)

	// query employee
	employees, err = qEmployee.WithContext(dbcache.WithCtx(l.ctx)).Where(qEmployee.ID.In(in.EmployeeID...)).Find()
	if err != nil {
		return err
	}

	// leave query
	leave, err = qLeave.WithContext(dbcache.WithCtx(l.ctx)).
		Preload(qLeave.LeaveGroup).
		Preload(qLeave.LeaveGroup.Employee).
		Preload(qLeave.LeaveGroup.LeaveGroupCondition).
		Where(qLeave.ID.Eq(in.LeaveID)).First()
	if err != nil {
		return err
	}

	// 開啟事務
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qCorrect = tx.LeaveCorrect
			corrects = make([]*types.LeaveCorrect, 0)
			err      error
		)

		for _, employee := range employees {
			var correct []*types.LeaveCorrect
			var temp *types.LeaveCorrect
			switch leave.Cycle {
			case enum.Annual:
				temp, err = l.annual(employee, leave, in.Year)
				correct = append(correct, temp)
			case enum.Calendar:
				temp, err = l.calendar(employee, leave, in.Year)
				correct = append(correct, temp)
			case enum.CalendarTwice:
				correct, err = l.calendarTwice(employee, leave, in.Year)
			case enum.Default:
				correct = append(correct, &types.LeaveCorrect{
					Effective:  driver.NewDate(fmt.Sprintf("%d-01-01", in.Year)),
					Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", in.Year)),
					EmployeeID: employee.ID,
					LeaveID:    in.LeaveID,
					Available:  float64(leave.Default),
				})
			}
			if err != nil {
				return err
			}

			// 刪除原有資料
			for _, c := range correct {
				_, err = qCorrect.WithContext(dbcache.WithCtx(l.ctx)).
					Where(qCorrect.EmployeeID.Eq(employee.ID), qCorrect.LeaveID.Eq(in.LeaveID), qCorrect.Effective.Eq(c.Effective), qCorrect.Expired.Eq(c.Expired)).
					Unscoped().Delete()
				if err != nil {
					return err
				}
			}

			corrects = append(corrects, correct...)
		}

		// 建立
		err = qCorrect.WithContext(dbcache.WithCtx(l.ctx)).Create(corrects...)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// 週年制計算
func (l *leave) annual(employee *types.Employee, leave *types.Leave, year int) (*types.LeaveCorrect, error) {
	var (
		result        = new(types.LeaveCorrect)
		currDate      = time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)
		years, months int
		conditions    []*types.LeaveGroupCondition
		currCondIndex = -1 // 套用的條件index
	)

	result.LeaveID = leave.ID
	result.EmployeeID = employee.ID

	// 計算年資
	years, months, _ = utils.CalcSeniority(employee.HireDate.Time(), currDate)

	// 搜尋套用到哪個群組
	for _, g := range leave.LeaveGroup {
		var exist = false
		for _, e := range g.Employee {
			if e.ID == employee.ID {
				exist = true
				break
			}
		}

		if exist {
			conditions = g.LeaveGroupCondition
			break
		}
	}

	// 判斷條件
	for index, cond := range conditions {
		var isChange = false

		switch cond.ConditionType {
		case enum.ConditionMonth:
			// 年資滿<#>個月
			if uint(years*12+months) >= cond.ConditionNum {
				currCondIndex = index
				isChange = true
			}
		case enum.ConditionYear:
			// 年資滿<#>年
			if uint(years) >= cond.ConditionNum {
				currCondIndex = index
				isChange = true
			}
		}

		if !isChange {
			// 如果沒達成條件 直接退出
			break
		}
	}

	// 帶入預設生效時間和過期時間
	var effDay = driver.NewDate(fmt.Sprintf("%d-%d-%d", year, employee.HireDate.Month(), employee.HireDate.Day()))
	result.Effective = effDay
	result.Expired = effDay.AddDate(1, 0, 0).AddDate(0, 0, -1)

	if currCondIndex == -1 || conditions == nil {
		// 如果連一關都沒有達成或是找不到群組，直接帶入預設天數
		result.Available = float64(leave.Default)
		return result, nil
	}

	// 帶入套用的條件
	if conditions[currCondIndex].ConditionType == enum.ConditionMonth && years == 0 {
		// 如果套用的條件為月和年資不滿一年
		// 從滿條件的那一天開始套用
		result.Effective = effDay.AddDate(0, int(conditions[currCondIndex].ConditionNum), 0)
	}

	result.Available = float64(conditions[currCondIndex].Result)

	return result, nil
}

// 曆年制計算
func (l *leave) calendar(employee *types.Employee, leave *types.Leave, year int) (*types.LeaveCorrect, error) {
	var (
		result        = new(types.LeaveCorrect)
		years, months int
		conditions    []*types.LeaveGroupCondition
		currDate      = time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)
		currCondIndex = -1 // 套用的條件index
	)

	result.LeaveID = leave.ID
	result.EmployeeID = employee.ID

	// 計算年資
	years, months, _ = utils.CalcSeniority(employee.HireDate.Time(), currDate)

	// 搜尋套用到哪個群組
	for _, g := range leave.LeaveGroup {
		var exist = false
		for _, e := range g.Employee {
			if e.ID == employee.ID {
				exist = true
				break
			}
		}

		if exist {
			conditions = g.LeaveGroupCondition
			break
		}
	}

	// 判斷條件
	for index, cond := range conditions {
		var isChange = false

		switch cond.ConditionType {
		case enum.ConditionMonth:
			// 年資滿<#>個月
			if uint(years*12+months) >= cond.ConditionNum {
				currCondIndex = index
				isChange = true
			}
		case enum.ConditionYear:
			// 年資滿<#>年
			if uint(years) >= cond.ConditionNum {
				currCondIndex = index
				isChange = true
			}
		}

		if !isChange {
			// 如果沒達成條件 直接退出
			break
		}
	}

	// 寫入預設生效時間和過期時間
	result.Effective = driver.NewDate(fmt.Sprintf("%d-01-01", year))
	result.Expired = driver.NewDate(fmt.Sprintf("%d-12-31", year))

	if currCondIndex == -1 || conditions == nil {
		// 如果連一關都沒有達成或是找不到群組，直接帶入預設天數
		result.Available = float64(leave.Default)
		return result, nil
	}

	var (
		availableDay    = float64(conditions[currCondIndex].Result)
		hireMonth       = float64(employee.HireDate.Month() - 1)
		hireDay         = float64(employee.HireDate.Day() - 1)
		daysByMonth     = utils.GetDay(employee.HireDate.Year(), employee.HireDate.Month())
		conditionNum    = float64(conditions[currCondIndex].ConditionNum)
		dayMonthPercent = (hireMonth + hireDay/daysByMonth) / 12
	)

	if years == 0 {
		// 無論如何，到今天未滿第一個條件，直接回傳0
		if _, m, _ := utils.CalcSeniority(employee.HireDate.Time(), currDate); float64(m) < conditionNum {
			result.Available = 0
		} else {
			// 未滿一年
			var effDay = driver.NewDate(fmt.Sprintf("%d-%d-%d", year, employee.HireDate.Month(), employee.HireDate.Day()))
			result.Effective = effDay.AddDate(0, int(conditions[currCondIndex].ConditionNum), 0)
			result.Available = availableDay - (hireMonth+hireDay/daysByMonth)/conditionNum*availableDay
		}
	} else if years > 0 {
		var (
			lastAvailableDay float64
			lastConditionNum float64
		)

		if years == 1 {
			if len(conditions) > 1 {
				lastAvailableDay = float64(conditions[currCondIndex-1].Result)
				lastConditionNum = float64(conditions[currCondIndex-1].ConditionNum)
			} else {
				lastAvailableDay = float64(leave.Default)
			}

			if conditions[currCondIndex-1].ConditionType == enum.ConditionYear {
				lastConditionNum *= 12
			}

			if float64(months) >= lastConditionNum {
				if len(conditions) == 1 {
					// 滿一年，加上去年有滿足第一個條件，但是條件只有一個的情況下，直接帶入可用天數
					result.Available = availableDay
				} else {
					// 滿一年，加上去年有滿足第一個條件
					result.Available = (hireMonth+hireDay/daysByMonth)/lastConditionNum*lastAvailableDay + (availableDay - dayMonthPercent*availableDay)
				}
			} else {
				// 滿一年，但去年沒有滿足第一個條件
				result.Available = lastAvailableDay + (availableDay - dayMonthPercent*availableDay)
			}
		} else {
			// 超過一年
			if years-1 >= int(conditions[currCondIndex].ConditionNum) {
				// 如果前一年也符合條件
				// 前一年的套用天數會等於今天套用天數
				lastAvailableDay = availableDay
			} else {
				lastAvailableDay = float64(conditions[currCondIndex-1].Result)
			}
			result.Available = dayMonthPercent*lastAvailableDay + (availableDay - dayMonthPercent*availableDay)
		}
	}

	return result, nil
}

// 曆年制計算(拆兩次給假)
func (l *leave) calendarTwice(employee *types.Employee, leave *types.Leave, year int) ([]*types.LeaveCorrect, error) {
	var (
		result        = make([]*types.LeaveCorrect, 0)
		years, months int
		conditions    []*types.LeaveGroupCondition
		currDate      = time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)
		currCondIndex = -1 // 套用的條件index
	)

	// 計算年資
	years, months, _ = utils.CalcSeniority(employee.HireDate.Time(), currDate)

	// 搜尋套用到哪個群組
	for _, g := range leave.LeaveGroup {
		var exist = false
		for _, e := range g.Employee {
			if e.ID == employee.ID {
				exist = true
				break
			}
		}

		if exist {
			conditions = g.LeaveGroupCondition
			break
		}
	}

	// 判斷條件
	for index, cond := range conditions {
		var isChange = false

		switch cond.ConditionType {
		case enum.ConditionMonth:
			// 年資滿<#>個月
			if uint(years*12+months) >= cond.ConditionNum {
				currCondIndex = index
				isChange = true
			}
		case enum.ConditionYear:
			// 年資滿<#>年
			if uint(years) >= cond.ConditionNum {
				currCondIndex = index
				isChange = true
			}
		}

		if !isChange {
			// 如果沒達成條件 直接退出
			break
		}
	}

	if currCondIndex == -1 || conditions == nil {
		// 如果連一關都沒有達成或是找不到群組，直接帶入預設天數
		result = append(result, &types.LeaveCorrect{
			LeaveID:    leave.ID,
			EmployeeID: employee.ID,
			Effective:  driver.NewDate(fmt.Sprintf("%d-01-01", year)),
			Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
			Available:  float64(leave.Default),
		})
		return result, nil
	}

	var (
		availableDay    = float64(conditions[currCondIndex].Result)
		hireMonth       = float64(employee.HireDate.Month() - 1)
		hireDay         = float64(employee.HireDate.Day() - 1)
		daysByMonth     = utils.GetDay(employee.HireDate.Year(), employee.HireDate.Month())
		conditionNum    = float64(conditions[currCondIndex].ConditionNum)
		dayMonthPercent = (hireMonth + hireDay/daysByMonth) / 12
	)

	if years == 0 {
		// 無論如何，到今天未滿第一個條件，直接回傳0
		if _, m, _ := utils.CalcSeniority(employee.HireDate.Time(), currDate); float64(m) < conditionNum {
			result = append(result, &types.LeaveCorrect{
				LeaveID:    leave.ID,
				EmployeeID: employee.ID,
				Effective:  driver.NewDate(fmt.Sprintf("%d-01-01", year)),
				Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
				Available:  0,
			})
		} else {
			// 未滿一年
			var effDay = driver.NewDate(fmt.Sprintf("%d-%d-%d", year, employee.HireDate.Month(), employee.HireDate.Day()))
			result = append(result, &types.LeaveCorrect{
				LeaveID:    leave.ID,
				EmployeeID: employee.ID,
				Effective:  effDay.AddDate(0, int(conditions[currCondIndex].ConditionNum), 0),
				Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
				Available:  availableDay - (hireMonth+hireDay/daysByMonth)/conditionNum*availableDay,
			})
		}
	} else if years > 0 {
		var (
			lastAvailableDay float64
			lastConditionNum float64
		)

		if years == 1 {
			if len(conditions) > 1 {
				lastAvailableDay = float64(conditions[currCondIndex-1].Result)
				lastConditionNum = float64(conditions[currCondIndex-1].ConditionNum)
			} else {
				lastAvailableDay = float64(leave.Default)
			}

			if conditions[currCondIndex-1].ConditionType == enum.ConditionYear {
				lastConditionNum *= 12
			}

			if float64(months) >= lastConditionNum {
				if len(conditions) == 1 {
					// 滿一年，加上去年有滿足第一個條件，但是條件只有一個的情況下，直接帶入可用天數
					result = append(result, &types.LeaveCorrect{
						LeaveID:    leave.ID,
						EmployeeID: employee.ID,
						Effective:  driver.NewDate(fmt.Sprintf("%d-01-01", year)),
						Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
						Available:  availableDay,
					})
				} else {
					// 滿一年，加上去年有滿足第一個條件
					// 拆兩次給予
					// 1/1 ~ 12/31
					result = append(result, &types.LeaveCorrect{
						LeaveID:    leave.ID,
						EmployeeID: employee.ID,
						Effective:  driver.NewDate(fmt.Sprintf("%d-01-01", year)),
						Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
						Available:  (hireMonth + hireDay/daysByMonth) / lastConditionNum * lastAvailableDay,
					})
					// hireDate ~ 12/31
					result = append(result, &types.LeaveCorrect{
						LeaveID:    leave.ID,
						EmployeeID: employee.ID,
						Effective:  driver.NewDate(fmt.Sprintf("%d-%d-%d", year, employee.HireDate.Month(), employee.HireDate.Day())),
						Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
						Available:  availableDay - dayMonthPercent*availableDay,
					})
				}
			} else {
				// 滿一年，但去年沒有滿足第一個條件
				// 拆兩次給予
				// 滿足條件的當天 ~ 12/31
				result = append(result, &types.LeaveCorrect{
					LeaveID:    leave.ID,
					EmployeeID: employee.ID,
					Effective:  employee.HireDate.AddDate(0, int(lastConditionNum), 0),
					Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
					Available:  lastAvailableDay,
				})
				// hireDate ~ 12/31
				result = append(result, &types.LeaveCorrect{
					LeaveID:    leave.ID,
					EmployeeID: employee.ID,
					Effective:  driver.NewDate(fmt.Sprintf("%d-%d-%d", year, employee.HireDate.Month(), employee.HireDate.Day())),
					Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
					Available:  availableDay - dayMonthPercent*availableDay,
				})
			}
		} else {
			// 超過一年
			if years-1 >= int(conditions[currCondIndex].ConditionNum) {
				// 如果前一年也符合條件
				// 前一年的套用天數會等於今天套用天數
				lastAvailableDay = availableDay
			} else {
				lastAvailableDay = float64(conditions[currCondIndex-1].Result)
			}
			// 拆兩次給予
			// 1/1 ~ 12/31
			result = append(result, &types.LeaveCorrect{
				LeaveID:    leave.ID,
				EmployeeID: employee.ID,
				Effective:  driver.NewDate(fmt.Sprintf("%d-01-01", year)),
				Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
				Available:  dayMonthPercent * lastAvailableDay,
			})
			// hireDate ~ 12/31
			result = append(result, &types.LeaveCorrect{
				LeaveID:    leave.ID,
				EmployeeID: employee.ID,
				Effective:  driver.NewDate(fmt.Sprintf("%d-%d-%d", year, employee.HireDate.Month(), employee.HireDate.Day())),
				Expired:    driver.NewDate(fmt.Sprintf("%d-12-31", year)),
				Available:  availableDay - dayMonthPercent*availableDay,
			})
		}
	}

	// 四捨五入
	for _, r := range result {
		r.RoundA()
	}

	return result, nil
}
