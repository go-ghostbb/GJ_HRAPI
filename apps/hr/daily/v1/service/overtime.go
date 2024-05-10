package service

import (
	"context"
	"database/sql"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	gbi18n "ghostbb.io/gb/i18n/gb_i18n"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/daily/v1/model"
	"hrapi/apps/hr/sign-off/v1/service"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/identity"
	"hrapi/internal/utils/paginator"
	"slices"
)

func Overtime(ctx context.Context, page ...*paginator.Pagination) IOvertime {
	if len(page) > 0 {
		return &overtime{ctx: ctx, page: page[0]}
	}
	return &overtime{ctx: ctx}
}

type (
	IOvertime interface {
		// GetOvertimeFormByEmployee 查詢加班單
		GetOvertimeFormByEmployee(in model.GetOvertimeFormByEmployeeReq) (out model.GetOvertimeFormByEmployeeRes, err error)
		// SaveOvertimeForm 儲存加班單
		SaveOvertimeForm(in model.SaveOvertimeFormReq) (out model.SaveOvertimeFormRes, err error)
		// DeleteOvertimeForm 刪除加班單
		DeleteOvertimeForm(in model.DeleteOvertimeFormReq) error
	}

	overtime struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetOvertimeFormByEmployee 查詢加班單
func (o *overtime) GetOvertimeFormByEmployee(in model.GetOvertimeFormByEmployeeReq) (out model.GetOvertimeFormByEmployeeRes, err error) {
	var (
		qForm = query.OvertimeRequestForm
		start = gbconv.Time(in.StartDate)
		end   = gbconv.Time(in.EndDate)
		conds = []gen.Condition{qForm.EmployeeID.Eq(in.EmployeeID)}
	)

	if in.StartDate != "" && in.EndDate != "" {
		conds = append(conds, qForm.CreatedAt.Between(start, end.AddDate(0, 0, 1)))
	}

	out.Items, err = qForm.WithContext(dbcache.WithCtx(o.ctx)).
		Scopes(o.page.Where(conds...)).
		Preload(field.Associations).
		Order(qForm.CreatedAt.Desc()).Find()
	if err != nil {
		return
	}

	out.Total = o.page.Total

	return
}

// SaveOvertimeForm 儲存加班單
func (o *overtime) SaveOvertimeForm(in model.SaveOvertimeFormReq) (out model.SaveOvertimeFormRes, err error) {
	var (
		qForm = query.OvertimeRequestForm
		form  = new(types.OvertimeRequestForm)
	)

	if in.ID != 0 {
		// 如果id不為0代表進行更新
		// 查詢原有資料
		form, err = qForm.WithContext(o.ctx).Where(qForm.ID.Eq(in.ID)).First()
		if err != nil {
			return
		}
	}

	// 複製form(request body)到database model
	if err = copier.Copy(form, &in); err != nil {
		return
	}

	// 判斷時間有沒有重疊
	var overlapCheck bool
	if overlapCheck, err = o.isOverlap(form); err != nil {
		return
	} else if overlapCheck {
		// 回傳時間重疊錯誤
		err = gberror.New("time overlap")
		return
	}

	// 寫入基本訊息, 部門id, 代理人部門id等...
	if err = o.writeBasicInfo(form); err != nil {
		return
	}

	// 開啟事務
	err = query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm = tx.OvertimeRequestForm
			qFlow = tx.OvertimeSignOffFlow
			qRate = tx.OvertimeRequestFormRate
			err   error
		)

		if form.ID == 0 {
			form.Order, err = identity.OvertimeOrder(o.ctx)
			if err != nil {
				return err
			}
		}

		// 刪除原有的倍率
		_, err = qRate.WithContext(dbcache.WithCtx(o.ctx)).Unscoped().Where(qRate.OvertimeRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除原本的簽核流程
		_, err = qFlow.WithContext(dbcache.WithCtx(o.ctx)).Unscoped().Where(qFlow.OvertimeRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 創建流程
		if err = o.createFlow(form); err != nil {
			return err
		}

		// 創建倍率
		if err = o.createRate(form); err != nil {
			return err
		}

		// 儲存至資料庫
		// create or update
		if err = qForm.WithContext(dbcache.WithCtx(o.ctx)).Save(form); err != nil {
			return err
		}

		// commit
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
	if err != nil {
		return
	}

	// 開啟一個新的service
	// 因為如果response的話，context會被cancel
	var locale = gbi18n.LanguageFromCtx(o.ctx)
	go service.Overtime(gbi18n.WithLanguage(gbctx.New(), locale)).NextFlow(form.SignOffFlow[0].UUID)

	out.ID = form.ID
	out.Order = form.Order
	return out, nil
}

// DeleteOvertimeForm 刪除加班單
func (o *overtime) DeleteOvertimeForm(in model.DeleteOvertimeFormReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm = tx.OvertimeRequestForm
			qFlow = tx.OvertimeSignOffFlow
			qRate = tx.OvertimeRequestFormRate
			err   error
		)

		// 刪除流程
		_, err = qFlow.WithContext(dbcache.WithCtx(o.ctx)).Where(qFlow.OvertimeRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除倍率
		_, err = qRate.WithContext(dbcache.WithCtx(o.ctx)).Where(qRate.OvertimeRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除表單
		_, err = qForm.WithContext(dbcache.WithCtx(o.ctx)).Where(qForm.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
}

// 寫入基本訊息
func (o *overtime) writeBasicInfo(form *types.OvertimeRequestForm) (err error) {
	var (
		qEmployee         = query.Employee
		qVacation         = query.Vacation
		qVacationSchedule = query.VacationSchedule
	)

	// 查詢員工部門id
	err = qEmployee.WithContext(dbcache.WithCtx(o.ctx)).Select(qEmployee.DepartmentId).Where(qEmployee.ID.Eq(form.EmployeeID)).Scan(&form.DepartmentID)
	if err != nil {
		return
	}

	// 尋找vacation id
	err = qVacationSchedule.WithContext(o.ctx).Select(qVacationSchedule.VacationID).
		Join(qVacation, qVacationSchedule.VacationID.EqCol(qVacation.ID)).
		Where(qVacationSchedule.ScheduleDate.Eq(form.Date)).
		Order(qVacation.Weight.Desc()).
		Offset(0).Limit(1).Scan(&form.VacationID)
	if err != nil {
		return err
	}

	// 計算預計時數
	var (
		start = form.StartTime.Time()
		end   = form.EndTime.Time()
	)

	if form.StartTime > form.EndTime {
		end = end.AddDate(0, 0, 1)
	}

	form.EstimatedHours = float32(end.Unix()-start.Unix()) / 60 / 60

	return nil
}

// 判斷有沒有重疊
func (o *overtime) isOverlap(form *types.OvertimeRequestForm) (bool, error) {
	var (
		qForm = query.OvertimeRequestForm
	)

	queryRes, err := qForm.WithContext(dbcache.WithCtx(o.ctx)).Where(qForm.Date.Eq(form.Date)).Find()
	if err != nil {
		return false, err
	}

	// 判斷重疊的時間裡時間是不是也重疊
	// 遍歷尋出來的表單
	for _, f := range queryRes {
		// A1 < B2 && A2 > B1
		var (
			A1 = gbconv.Time(f.Date.Format() + " " + f.StartTime.Format())
			A2 = gbconv.Time(f.Date.Format() + " " + f.EndTime.Format())
		)

		// 如果開始時間>結束時間
		// 代表隔夜，需要+1天
		if f.StartTime.Unix() > f.EndTime.Unix() {
			A2 = A2.AddDate(0, 0, 1)
		}

		var (
			B1 = gbconv.Time(form.Date.Format() + " " + form.StartTime.Format())
			B2 = gbconv.Time(form.Date.Format() + " " + form.EndTime.Format())
		)

		// 如果開始時間>結束時間
		// 代表隔夜，需要+1天
		if form.StartTime.Unix() > form.EndTime.Unix() {
			B2 = B2.AddDate(0, 0, 1)
		}

		// 判斷是否重疊
		// 公式：A1 < B2 && A2 > B1
		if A1.Unix() < B2.Unix() && A2.Unix() > B1.Unix() {
			return true, nil
		}
	}

	return false, nil
}

// 創建流程
func (o *overtime) createFlow(form *types.OvertimeRequestForm) error {
	var (
		qFlowSetting = query.OvertimeSignOffSetting
		qDepartment  = query.Department
		flowSetting  []*types.OvertimeSignOffSetting
		err          error
	)

	// 查詢簽核設定
	flowSetting, err = qFlowSetting.WithContext(dbcache.WithCtx(o.ctx)).
		Where(qFlowSetting.DepartmentID.Eq(form.DepartmentID), qFlowSetting.VacationID.Eq(form.VacationID)).
		Find()
	if err != nil {
		return err
	}

	// 查詢部門所有父子關係
	// with tree as ((?) union all select b.* from tree inner join department b on b.id = tree.parent_id
	var departments []*types.Department
	departments, err = qDepartment.WithContext(o.ctx).QueryUp(form.DepartmentID)
	if err != nil {
		return err
	}

	// 流程判斷
	var deptNowID = form.DepartmentID // 記錄當前部門id, 方便向上查詢

	for _, fs := range flowSetting {
		var signEmpID = uint(0) // 紀錄當前關卡需簽核的員工

		// 如果條件(>=時數)為0或是達成條件
		// 0代表無論如何都必須使用此關卡
		if fs.GteHour == 0 || form.EstimatedHours >= fs.GteHour {
			switch fs.SignType {
			case enum.DepartmentManager:
				// 簽核對象為部門主管
				for {
					index := slices.IndexFunc(departments, func(d *types.Department) bool {
						// 找到當前部門
						return d.ID == deptNowID
					})

					// 如果不是-1代表有找到部門
					if index != -1 {
						if departments[index].ManagerId == form.EmployeeID {
							// 如果部門主管為起單人，繼續向上搜尋
							deptNowID = departments[index].ParentID
							continue
						} else {
							// 如果不是起單人，代表此關卡簽核員工為此部門的主管
							signEmpID = departments[index].ManagerId
						}
					}

					// 如果沒搜尋到或是搜尋到且部門主管不是起單人則離開
					break
				}
			case enum.SpecificEmployee:
				// 簽核對象為特定人員
				// 如果特定人員等於起單人，則不加入流程
				if fs.SpecificEmployeeID != form.EmployeeID {
					signEmpID = fs.SpecificEmployeeID
				}
			}
		}

		// 不等於0代表條件達成，加入流程
		// 等於0代表向上沒有部門或其他邏輯問題
		if signEmpID != 0 {
			form.SignOffFlow = append(form.SignOffFlow, &types.OvertimeSignOffFlow{
				OvertimeRequestFormID: form.ID,
				SignOffEmployeeID:     signEmpID,
				SignType:              fs.SignType,
				Notify:                fs.Notify,
				Remark:                fs.Remark,
			})
		}
	}

	// 最後一關固定為通知起單人
	// 簽核完成後通知起單人
	form.SignOffFlow = append(form.SignOffFlow, &types.OvertimeSignOffFlow{
		OvertimeRequestFormID: form.ID,
		SignOffEmployeeID:     form.EmployeeID,
		SignType:              enum.SpecificEmployee, // 特定人員:起單人
		Notify:                enum.NotifyOnly,       // 僅通知
	})

	// 遍歷關卡
	for index, flow := range form.SignOffFlow {
		flow.Level = uint(index + 1)                             // 重新定義關卡
		flow.UUID = uuid.Must(uuid.NewRandom()).String()         // 為每一關生成uuid
		flow.Locale = enum.Locale(gbi18n.LanguageFromCtx(o.ctx)) // 語言為請求時使用的語言
	}

	return nil
}

// 創建倍率
func (o *overtime) createRate(form *types.OvertimeRequestForm) error {
	var (
		qM2M   = query.VacationGroupEmployee
		qGroup = query.VacationGroup
		qRate  = query.VacationGroupOvertimeRate
		rate   []*types.VacationGroupOvertimeRate
		err    error
	)

	// 搜尋套用倍率
	// select rate.*
	// from vacation_group_overtime_rate as rate
	// join vacation_group as [group] on (rate.vacation_group_id = [group].id)
	// join vacation_group_employee as m2m on (m2m.vacation_group_id = [group].id)
	// where m2m.employee_id = @empID and [group].vacation_id = @vacationID;
	rate, err = qRate.WithContext(dbcache.WithCtx(o.ctx)).Select(qRate.ALL).
		Join(qGroup, qRate.VacationGroupID.EqCol(qGroup.ID)).
		Join(qM2M, qM2M.VacationGroupID.EqCol(qGroup.ID)).
		Where(qM2M.EmployeeID.Eq(form.EmployeeID), qGroup.VacationID.Eq(form.VacationID)).
		Find()
	if err != nil {
		return err
	}

	// copier
	if err = copier.Copy(&form.Rate, rate); err != nil {
		return err
	}

	return nil
}
