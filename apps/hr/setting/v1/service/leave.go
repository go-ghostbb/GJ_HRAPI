package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/paginator"
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
