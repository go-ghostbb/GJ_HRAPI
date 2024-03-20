package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
)

func SignOffSetting(ctx context.Context) ISignOffSetting {
	return &signOffSetting{ctx: ctx}
}

type (
	ISignOffSetting interface {
		// GetLeaveSignOffSetting 查詢請假簽核設定
		GetLeaveSignOffSetting(in model.GetLeaveSignOffSettingReq) (out []*model.GetLeaveSignOffSettingRes, err error)
		// UpdateLeaveSignOffSettingBatch 批量更新leave sign off setting
		UpdateLeaveSignOffSettingBatch(in model.PutBatchLeaveSignOffSettingReq) error
		// GetOvertimeSignOffSetting 查詢加班簽核設定
		GetOvertimeSignOffSetting(in model.GetOvertimeSignOffSettingReq) (out []*model.GetOvertimeSignOffSettingRes, err error)
		// UpdateOvertimeSignOffSettingBatch 批量更新overtime sign off setting
		UpdateOvertimeSignOffSettingBatch(in model.PutBatchOvertimeSignOffSettingReq) error
	}

	signOffSetting struct {
		ctx context.Context
	}
)

// GetLeaveSignOffSetting 查詢請假簽核設定
func (s *signOffSetting) GetLeaveSignOffSetting(in model.GetLeaveSignOffSettingReq) (out []*model.GetLeaveSignOffSettingRes, err error) {
	var (
		qLeaveSetting = query.LeaveSignOffSetting
		leaveSetting  []*types.LeaveSignOffSetting
	)

	leaveSetting, err = qLeaveSetting.WithContext(dbcache.WithCtx(s.ctx)).
		Where(qLeaveSetting.LeaveID.Eq(in.LeaveID), qLeaveSetting.DepartmentID.Eq(in.DepartmentID)).
		Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(&out, leaveSetting); err != nil {
		return nil, err
	}

	return
}

// UpdateLeaveSignOffSettingBatch 批量更新leave sign off setting
func (s *signOffSetting) UpdateLeaveSignOffSettingBatch(in model.PutBatchLeaveSignOffSettingReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qLeaveSetting = tx.LeaveSignOffSetting
			err           error
		)
		// 刪除原本
		_, err = qLeaveSetting.WithContext(dbcache.WithCtx(s.ctx)).Unscoped().
			Where(qLeaveSetting.LeaveID.Eq(in.LeaveID), qLeaveSetting.DepartmentID.Eq(in.DepartmentID)).Delete()
		if err != nil {
			return err
		}

		// 建立新的
		for index, setting := range in.SignOffSetting {
			setting.Leave = nil
			setting.Department = nil
			setting.SpecificEmployee = nil
			setting.Level = uint(index) + 1
			setting.ID = 0
		}

		return qLeaveSetting.WithContext(dbcache.WithCtx(s.ctx)).Create(in.SignOffSetting...)
	})
}

// GetOvertimeSignOffSetting 查詢加班簽核設定
func (s *signOffSetting) GetOvertimeSignOffSetting(in model.GetOvertimeSignOffSettingReq) (out []*model.GetOvertimeSignOffSettingRes, err error) {
	var (
		qOvertimeSetting = query.OvertimeSignOffSetting
		overtimeSetting  []*types.OvertimeSignOffSetting
	)

	overtimeSetting, err = qOvertimeSetting.WithContext(dbcache.WithCtx(s.ctx)).
		Where(qOvertimeSetting.VacationID.Eq(in.VacationID), qOvertimeSetting.DepartmentID.Eq(in.DepartmentID)).
		Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(&out, overtimeSetting); err != nil {
		return nil, err
	}

	return
}

// UpdateOvertimeSignOffSettingBatch 批量更新overtime sign off setting
func (s *signOffSetting) UpdateOvertimeSignOffSettingBatch(in model.PutBatchOvertimeSignOffSettingReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qOvertimeSetting = tx.OvertimeSignOffSetting
			err              error
		)
		// 刪除原本
		_, err = qOvertimeSetting.WithContext(dbcache.WithCtx(s.ctx)).Unscoped().
			Where(qOvertimeSetting.VacationID.Eq(in.VacationID), qOvertimeSetting.DepartmentID.Eq(in.DepartmentID)).Delete()
		if err != nil {
			return err
		}

		// 建立新的
		for index, setting := range in.SignOffSetting {
			setting.Vacation = nil
			setting.Department = nil
			setting.SpecificEmployee = nil
			setting.Level = uint(index) + 1
			setting.ID = 0
		}

		return qOvertimeSetting.WithContext(dbcache.WithCtx(s.ctx)).Create(in.SignOffSetting...)
	})
}
