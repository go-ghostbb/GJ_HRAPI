package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/paginator"
)

func Vacation(ctx context.Context, page ...*paginator.Pagination) IVacation {
	if len(page) != 0 {
		return &vacation{ctx: ctx, page: page[0]}
	}
	return &vacation{ctx: ctx}
}

type (
	IVacation interface {
		// GetByKeyword 根據keyword, status獲取休假日類別
		GetByKeyword(in model.GetByKeywordVacationReq) (out model.GetByKeywordVacationRes, err error)
		// GetByID 根據ID獲取休假日類別
		GetByID(in model.GetByIDVacationReq) (out model.GetByIDVacationRes, err error)
		// Insert 新增休假日
		Insert(in model.PostVacationReq) (err error)
		// Update 更新休假日
		Update(in model.PutVacationReq) (err error)
		// Delete 刪除休假日
		Delete(in model.DeleteVacationReq) error
		// SetStatus 設定狀態
		SetStatus(in model.SetStatusVacationReq) (err error)
		// GetVacationGroup 獲取休假日群組列表
		GetVacationGroup(in model.GetVacationGroupReq) (out []*model.GetVacationGroupRes, err error)
		// InsertGroup 新增群組
		InsertGroup(in model.PostVacationGroupReq) error
		// DeleteGroup 刪除群組
		DeleteGroup(in model.DeleteVacationGroupReq) error
		// SetVacationGroupName 設置群組名稱
		SetVacationGroupName(in model.SetVacationGroupNameReq) (err error)
		// SetVacationGroupEmployee 設置群組的員工
		SetVacationGroupEmployee(in model.SetVacationGroupEmployeeReq) (err error)
		// SetVacationGroupOvertimeRate 設置群組的條件
		SetVacationGroupOvertimeRate(in model.SetVacationGroupOvertimeRateReq) error
		// GetScheduleByDate 根據日期區間查詢schedule
		GetScheduleByDate(in model.GetByDateVacationScheduleReq) (out []*model.GetByDateVacationScheduleRes, err error)
		// CreateSchedule 創建schedule
		CreateSchedule(in model.PostVacationScheduleReq) (err error)
		// UpdateSchedule 更新schedule
		UpdateSchedule(in model.PutVacationScheduleReq) (err error)
		// DeleteSchedule 刪除schedule
		DeleteSchedule(in model.DeleteVacationScheduleReq) (err error)
	}

	vacation struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword, status獲取休假日類別
func (v *vacation) GetByKeyword(in model.GetByKeywordVacationReq) (out model.GetByKeywordVacationRes, err error) {
	var (
		qVacation  = query.Vacation
		qdVacation = qVacation.WithContext(dbcache.WithCtx(v.ctx))
		conds      = []gen.Condition{qdVacation.Where(qVacation.Name.Like("%" + in.Keyword + "%")).Or(qVacation.Code.Like("%" + in.Keyword + "%"))}
		vacations  []*types.Vacation
	)
	if in.Status != "" {
		conds = append(conds, qVacation.Status.Is(gbconv.Bool(in.Status)))
	}

	vacations, err = qdVacation.Scopes(v.page.Where(conds...)).Preload(field.Associations).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, vacations); err != nil {
		return
	}

	out.Total = v.page.Total

	return out, nil
}

// GetByID 根據ID獲取休假日類別
func (v *vacation) GetByID(in model.GetByIDVacationReq) (out model.GetByIDVacationRes, err error) {
	var (
		qVacation = query.Vacation
	)

	out.Vacation, err = qVacation.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacation.ID.Eq(in.ID)).Preload(field.Associations).First()
	if err != nil {
		return
	}

	return out, nil
}

// Insert 新增休假日
func (v *vacation) Insert(in model.PostVacationReq) (err error) {
	var (
		qVacation = query.Vacation
	)
	in.Schedule = nil
	return qVacation.WithContext(dbcache.WithCtx(v.ctx)).Create(in.Vacation)
}

// Update 更新休假日
func (v *vacation) Update(in model.PutVacationReq) (err error) {
	var (
		qVacation = query.Vacation
	)
	in.Vacation.ID = in.ID
	in.Vacation.Schedule = nil

	return qVacation.WithContext(dbcache.WithCtx(v.ctx)).Save(in.Vacation)
}

// Delete 刪除休假日
func (v *vacation) Delete(in model.DeleteVacationReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qVacation          = tx.Vacation
			qSchedule          = tx.VacationSchedule
			qGroup             = tx.VacationGroup
			qGroupOvertimeRate = tx.VacationGroupOvertimeRate
			qGroupEmployee     = tx.VacationGroupEmployee
			groupIDs           = make([]uint, 0)
			err                error
		)

		// 刪除休假日
		_, err = qVacation.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacation.ID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// 刪除schedule
		_, err = qSchedule.WithContext(dbcache.WithCtx(v.ctx)).Where(qSchedule.VacationID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// 查詢有關連的群組ID
		err = qGroup.WithContext(v.ctx).Select(qGroup.ID).Where(qGroup.VacationID.Eq(in.ID)).Scan(&groupIDs)
		if err != nil {
			return err
		}

		// 刪除群組
		_, err = qGroup.WithContext(dbcache.WithCtx(v.ctx)).Where(qGroup.VacationID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// 刪除群組加班倍率
		_, err = qGroupOvertimeRate.WithContext(dbcache.WithCtx(v.ctx)).Where(qGroupOvertimeRate.VacationGroupID.In(groupIDs...)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// 刪除群組員工
		_, err = qGroupEmployee.WithContext(dbcache.WithCtx(v.ctx)).Where(qGroupEmployee.VacationGroupID.In(groupIDs...)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetStatus 設定狀態
func (v *vacation) SetStatus(in model.SetStatusVacationReq) (err error) {
	var (
		qVacation = query.Vacation
	)
	_, err = qVacation.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacation.ID.Eq(in.ID)).Update(qVacation.Status, in.Status)
	return err
}

// GetVacationGroup 獲取休假日群組列表
func (v *vacation) GetVacationGroup(in model.GetVacationGroupReq) (out []*model.GetVacationGroupRes, err error) {
	var (
		qVacationGroup = query.VacationGroup
		leaveGroups    []*types.VacationGroup
	)
	leaveGroups, err = qVacationGroup.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacationGroup.VacationID.Eq(in.VacationID)).
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
func (v *vacation) InsertGroup(in model.PostVacationGroupReq) error {
	var (
		qVacationGroup = query.VacationGroup
		err            error
	)

	// 避免從postman或swagger近來的資料污染
	in.Vacation = nil
	in.Employee = nil
	in.VacationGroupOvertimeRate = nil

	// 創建Vacation group
	err = qVacationGroup.WithContext(dbcache.WithCtx(v.ctx)).Create(in.VacationGroup)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup 刪除群組
func (v *vacation) DeleteGroup(in model.DeleteVacationGroupReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qVacationGroup             = tx.VacationGroup
			qVacationGroupOvertimeRate = tx.VacationGroupOvertimeRate
			qVacationGroupEmployee     = tx.VacationGroupEmployee
			err                        error
		)
		// 刪除群組
		_, err = qVacationGroup.WithContext(dbcache.WithCtx(v.ctx)).Unscoped().Where(qVacationGroup.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除倍率表
		_, err = qVacationGroupOvertimeRate.WithContext(dbcache.WithCtx(v.ctx)).Unscoped().Where(qVacationGroupOvertimeRate.VacationGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除員工表
		_, err = qVacationGroupEmployee.WithContext(dbcache.WithCtx(v.ctx)).Unscoped().Where(qVacationGroupEmployee.VacationGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetVacationGroupName 設置群組名稱
func (v *vacation) SetVacationGroupName(in model.SetVacationGroupNameReq) (err error) {
	_, err = query.VacationGroup.WithContext(dbcache.WithCtx(v.ctx)).Where(query.VacationGroup.ID.Eq(in.ID)).Update(query.VacationGroup.Name, in.Name)
	return err
}

// SetVacationGroupEmployee 設置群組的員工
func (v *vacation) SetVacationGroupEmployee(in model.SetVacationGroupEmployeeReq) (err error) {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			groupEmployee = make([]*types.VacationGroupEmployee, 0)
			qM2M          = tx.VacationGroupEmployee
		)

		for _, employeeID := range in.EmployeeID {
			groupEmployee = append(groupEmployee, &types.VacationGroupEmployee{VacationGroupID: in.ID, EmployeeID: employeeID})
		}

		// 先刪除原有員工
		_, err = qM2M.WithContext(dbcache.WithCtx(v.ctx)).Where(qM2M.VacationGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 建立新的
		err = qM2M.WithContext(dbcache.WithCtx(v.ctx)).Create(groupEmployee...)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetVacationGroupOvertimeRate 設置群組的條件
func (v *vacation) SetVacationGroupOvertimeRate(in model.SetVacationGroupOvertimeRateReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qVacationGroupOvertimeRate = tx.VacationGroupOvertimeRate
			err                        error
		)

		// 刪除原有的條件
		_, err = qVacationGroupOvertimeRate.WithContext(dbcache.WithCtx(v.ctx)).Unscoped().Where(qVacationGroupOvertimeRate.VacationGroupID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 將ID都變為0, ID將由資料庫自動遞增
		for index, rate := range in.OvertimeRate {
			rate.ID = 0
			rate.VacationGroupID = in.ID
			rate.Level = uint(index) + 1

			// 驗證輸入
			if rate.IsFill {
				if rate.Fill == 0 {
					return gberror.New("fill cannot be 0")
				}
				if len(in.OvertimeRate) > index+1 && rate.Fill >= in.OvertimeRate[index+1].Hours {
					return gberror.New("fill cannot be larger than the hours of the next interval")
				}
			}
		}

		// 建立新的倍率區間
		err = qVacationGroupOvertimeRate.WithContext(dbcache.WithCtx(v.ctx)).Create(in.OvertimeRate...)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// GetScheduleByDate 根據日期區間查詢schedule
func (v *vacation) GetScheduleByDate(in model.GetByDateVacationScheduleReq) (out []*model.GetByDateVacationScheduleRes, err error) {
	// 驗證
	err = g.Validator().Data(in).Run(v.ctx)
	if err != nil {
		return nil, err
	}

	// 查詢
	var (
		qVacationSchedule = query.VacationSchedule
		startDate         = gbconv.Time(in.Start)
		endDate           = gbconv.Time(in.End)
		schedule          []*types.VacationSchedule
	)
	schedule, err = qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).Preload(field.Associations).Where(qVacationSchedule.ScheduleDate.Between(startDate, endDate)).Find()
	if err != nil {
		return nil, err
	}

	// copier
	if err = copier.Copy(&out, schedule); err != nil {
		return nil, err
	}

	return out, nil
}

// CreateSchedule 創建schedule
func (v *vacation) CreateSchedule(in model.PostVacationScheduleReq) (err error) {
	var (
		qVacationSchedule = query.VacationSchedule
		schedule          []*types.VacationSchedule
	)
	schedule = v.genSchedule(in.VacationScheduleConfig, in.Remark, in.VacationID)
	return qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).Create(schedule...)
}

// UpdateSchedule 更新schedule
func (v *vacation) UpdateSchedule(in model.PutVacationScheduleReq) (err error) {
	// 暴力解決
	// 直接刪除原本資料
	// 重新建立新的
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qVacationSchedule = tx.VacationSchedule
			schedule          []*types.VacationSchedule
		)

		_, err = qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacationSchedule.GeneralKey.Eq(in.GeneralKey)).Unscoped().Delete()
		if err != nil {
			return err
		}

		schedule = v.genSchedule(in.VacationScheduleConfig, in.Remark, in.VacationID)
		return qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).Create(schedule...)
	})
}

// DeleteSchedule 刪除schedule
func (v *vacation) DeleteSchedule(in model.DeleteVacationScheduleReq) (err error) {
	var (
		qVacationSchedule  = query.VacationSchedule
		subQueryGeneralKey = qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).Select(qVacationSchedule.GeneralKey).Where(qVacationSchedule.ID.Eq(in.ID))
	)
	if in.Repeat {
		// 刪除重複
		_, err = qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).
			Where(qVacationSchedule.Columns(qVacationSchedule.GeneralKey).Eq(subQueryGeneralKey)).
			Unscoped().Delete()
	} else {
		// 只刪除自己
		_, err = qVacationSchedule.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacationSchedule.ID.Eq(in.ID)).Unscoped().Delete()
	}

	return err
}

// 產生schedule
func (v *vacation) genSchedule(config *types.VacationScheduleConfig, remark string, vacationID uint) []*types.VacationSchedule {
	var (
		result = make([]*types.VacationSchedule, 0)
		// create general_key
		generalKey = uuid.NewString()

		start     = config.StartDate
		end       = config.EndDate
		endRepeat = config.EndRepeat

		stepDay   = 0
		stepMonth = 0
		stepYear  = 0
	)

	switch config.Repeat {
	case enum.RepeatNone:
		// 如果不重複, 直接按照start, end遍歷
		for start.Before(end.AddDate(0, 0, 1)) {
			result = append(result, &types.VacationSchedule{
				ScheduleDate:           start,
				GeneralKey:             generalKey,
				VacationID:             vacationID,
				Remark:                 remark,
				VacationScheduleConfig: config,
			})
			start = start.AddDate(0, 0, 1)
		}
	default:
		// 判斷step
		switch config.Repeat {
		case enum.RepeatWeek:
			stepDay = 7
		case enum.RepeatMonth:
			stepMonth = 1
		case enum.RepeatYear:
			stepYear = 1
		}

		// 遍歷start, end每一天, 內部在遍歷一次到endRepeat
		for start.Before(end.AddDate(0, 0, 1)) {
			tmpStart := start
			// 遍歷日期 step 7 days
			for tmpStart.Before(endRepeat.AddDate(0, 0, 1)) {
				result = append(result, &types.VacationSchedule{
					ScheduleDate:           tmpStart,
					GeneralKey:             generalKey,
					VacationID:             vacationID,
					Remark:                 remark,
					VacationScheduleConfig: config,
				})
				// step
				tmpStart = tmpStart.AddDate(stepYear, stepMonth, stepDay)
			}
			// 下一天
			start = start.AddDate(0, 0, 1)
		}
	}
	return result
}
