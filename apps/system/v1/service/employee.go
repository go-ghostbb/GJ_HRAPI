package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gbconv "ghostbb.io/gb/util/gb_conv"
	gbrand "ghostbb.io/gb/util/gb_rand"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"hrapi/apps/system/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/password"
)

func Employee(ctx context.Context) IEmployee {
	return &employee{ctx: ctx}
}

type (
	IEmployee interface {
		GetByKeyword(in model.GetByKeywordEmployeeReq) (out []*model.GetByKeywordEmployeeRes, err error)
		GetByID(in model.GetByIDEmployeeReq) (out model.GetByIDEmployeeRes, err error)
		Insert(in model.PostEmployeeReq) (out model.PostEmployeeRes, err error)
		Update(in model.PutEmployeeReq) (err error)
		Delete(in model.DeleteEmployeeReq) (err error)
	}

	employee struct {
		ctx context.Context
	}
)

func (e *employee) GetByKeyword(in model.GetByKeywordEmployeeReq) (out []*model.GetByKeywordEmployeeRes, err error) {
	var (
		qEmployee  = query.Employee
		qdEmployee = qEmployee.WithContext(dbcache.WithCtx(e.ctx))
		employees  []*types.Employee
		conds      = []gen.Condition{
			// keyword => where (real_name like ? or national_id like ?)
			qdEmployee.Where(qdEmployee.Where(qEmployee.RealName.Like("%" + in.Keyword + "%")).Or(qEmployee.NationalID.Like("%" + in.Keyword + "%"))),
			// status => and employment_status = ?
			qEmployee.EmploymentStatus.Eq(in.EmploymentStatus),
		}
	)
	if in.DepartmentId != "" {
		// department_id => and department_id = ?
		conds = append(conds, qEmployee.DepartmentId.Eq(gbconv.Uint(in.DepartmentId)))
	}

	employees, err = qdEmployee.Preload(qEmployee.LoginInformation).Where(conds...).Find()

	if err = copier.Copy(&out, employees); err != nil {
		return
	}
	return
}

func (e *employee) GetByID(in model.GetByIDEmployeeReq) (out model.GetByIDEmployeeRes, err error) {
	qEmployee := query.Employee
	out.Employee, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Preload(qEmployee.LoginInformation).Where(qEmployee.ID.Eq(in.ID)).First()
	return
}

func (e *employee) Insert(in model.PostEmployeeReq) (out model.PostEmployeeRes, err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		var (
			qEmployee  = tx.Employee
			qLoginInfo = tx.LoginInformation
			tmpPass    = gbrand.S(10)
		)

		err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Create(in.Employee)
		if err != nil {
			return err
		}
		// 建立預設帳號密碼
		loginInfo := new(types.LoginInformation)
		loginInfo.EmployeeID = in.ID
		loginInfo.Account = in.NationalID
		loginInfo.Password = password.Hash(tmpPass)
		err = qLoginInfo.WithContext(dbcache.WithCtx(e.ctx)).Create(loginInfo)
		if err != nil {
			return err
		}
		out.Account = loginInfo.Account
		out.Password = tmpPass

		// commit
		return nil
	})

	return
}

func (e *employee) Update(in model.PutEmployeeReq) (err error) {
	// 寫入ID
	in.Employee.ID = in.ID

	return query.Employee.WithContext(dbcache.WithCtx(e.ctx)).Save(in.Employee)
}

func (e *employee) Delete(in model.DeleteEmployeeReq) (err error) {
	var (
		qEmployee = query.Employee
	)

	_, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Unscoped().Preload(qEmployee.LoginInformation).Where(qEmployee.ID.Eq(in.ID)).Delete()
	return
}
