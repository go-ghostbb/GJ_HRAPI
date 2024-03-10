package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
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
		// SetVacationGroupEmployee 設置群組的員工
		SetVacationGroupEmployee(in model.SetVacationGroupEmployeeReq) (err error)
		// SetVacationGroupOvertimeRate 設置群組的條件
		SetVacationGroupOvertimeRate(in model.SetVacationGroupOvertimeRateReq) error
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
			qVacation = tx.Vacation
			qSchedule = tx.VacationSchedule
			err       error
		)

		_, err = qVacation.WithContext(dbcache.WithCtx(v.ctx)).Where(qVacation.ID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		_, err = qSchedule.WithContext(dbcache.WithCtx(v.ctx)).Where(qSchedule.VacationID.Eq(in.ID)).Unscoped().Delete()
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
