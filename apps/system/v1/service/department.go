package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"hrapi/apps/system/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
)

func Department(ctx context.Context) IDepartment {
	return &department{ctx: ctx}
}

type (
	IDepartment interface {
		// GetByKeyword 根據keyword, status獲取部門
		GetByKeyword(in model.GetByKeywordDepartmentReq) (out []*model.GetByKeywordDepartmentRes, err error)
		// GetByID 根據ID獲取部門
		GetByID(in model.GetByIDDepartmentReq) (out model.GetByIDDepartmentRes, err error)
		// Insert 新增部門
		Insert(in model.PostDepartmentReq) (err error)
		// Update 更新部門
		Update(in model.PutDepartmentReq) (err error)
		// Delete 刪除部門
		Delete(in model.DeleteDepartmentReq) (err error)
		// SetStatus 設置狀態
		SetStatus(in model.SetStatusDepartmentReq) (err error)
	}

	department struct {
		ctx context.Context
	}
)

// GetByKeyword 根據keyword, status獲取部門
func (d *department) GetByKeyword(in model.GetByKeywordDepartmentReq) (out []*model.GetByKeywordDepartmentRes, err error) {
	var (
		qDept = query.Department
		depts []*types.Department
		conds = []gen.Condition{qDept.Name.Like("%" + in.Keyword + "%")}
	)
	if in.Status != "" {
		conds = append(conds, qDept.Status.Is(gbconv.Bool(in.Status)))
	}

	depts, err = qDept.WithContext(dbcache.WithCtx(d.ctx)).Preload(qDept.Manager).Where(conds...).Find()

	// 組裝
	newDepts := d.assemble(depts)

	if err = copier.Copy(&out, newDepts); err != nil {
		return
	}

	return
}

// GetByID 根據ID獲取部門
func (d *department) GetByID(in model.GetByIDDepartmentReq) (out model.GetByIDDepartmentRes, err error) {
	var (
		qDept = query.Department
	)
	out.Department, err = qDept.WithContext(dbcache.WithCtx(d.ctx)).Where(qDept.ID.Eq(in.ID)).First()
	return
}

// Insert 新增部門
func (d *department) Insert(in model.PostDepartmentReq) (err error) {
	return query.Department.WithContext(dbcache.WithCtx(d.ctx)).Create(in.Department)
}

// Update 更新部門
func (d *department) Update(in model.PutDepartmentReq) (err error) {
	// 寫入ID
	in.Department.ID = in.ID

	return query.Department.WithContext(dbcache.WithCtx(d.ctx)).Save(in.Department)
}

// Delete 刪除部門
func (d *department) Delete(in model.DeleteDepartmentReq) (err error) {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qDept     = tx.Department
			qEmployee = tx.Employee
		)

		// 刪除部門, 使用硬刪除
		_, err = qDept.WithContext(dbcache.WithCtx(d.ctx)).Unscoped().Where(qDept.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 將這個部門的員工改為無部門
		_, err = qEmployee.WithContext(dbcache.WithCtx(d.ctx)).Where(qEmployee.DepartmentId.Eq(in.ID)).Update(qEmployee.DepartmentId, 0)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetStatus 設置狀態
func (d *department) SetStatus(in model.SetStatusDepartmentReq) (err error) {
	_, err = query.Department.WithContext(dbcache.WithCtx(d.ctx)).
		Where(query.Department.ID.Eq(in.ID)).Update(query.Department.Status, in.Status)
	return err
}

func (d *department) assemble(depts []*types.Department) (result []*types.Department) {
	var (
		check   = make(map[uint]struct{})
		item    = make([]*types.Department, 0)
		deptMap = make(map[uint]*types.Department)
	)

	result = make([]*types.Department, 0)

	// 利用指針組裝
	for _, dept := range depts {
		if _, ok := check[dept.ID]; !ok {
			if dept.ParentID == 0 {
				// 如果沒父節點
				// 直接append回傳陣列
				result = append(result, dept)
				// 加入map
				// 等待item加入
				deptMap[dept.ID] = dept
			} else {
				// 否則append組裝陣列
				item = append(item, dept)
			}

			// 標記
			check[dept.ID] = struct{}{}
		}
	}

	for _, dept := range item {
		if _, ok := deptMap[dept.ParentID]; ok {
			deptMap[dept.ParentID].Children = append(deptMap[dept.ParentID].Children, dept)
		} else {
			result = append(result, dept)
		}
	}

	return
}
