package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gbrand "ghostbb.io/gb/util/gb_rand"
	"hrapi/apps/system/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/password"
)

func Employee(ctx context.Context) IEmployee {
	return &employee{ctx: ctx}
}

type (
	IEmployee interface {
		GetByKeyword(in model.GetByKeywordEmployeeReq) (out model.GetByKeywordEmployeeRes, err error)
		GetByID(in model.GetByIDEmployeeReq) (out model.GetByIDEmployeeRes, err error)
		Insert(in model.PostEmployeeReq) (out model.PostEmployeeRes, err error)
		Update(in model.PutEmployeeReq) (out model.PutEmployeeRes, err error)
		Delete(in model.DeleteEmployeeReq) (out model.DeleteEmployeeRes, err error)
	}

	employee struct {
		ctx context.Context
	}
)

func (e *employee) GetByKeyword(in model.GetByKeywordEmployeeReq) (out model.GetByKeywordEmployeeRes, err error) {
	qEmployee := query.Employee
	out.Employees, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Preload(qEmployee.LoginInformation).
		Where(qEmployee.RealName.Like("%" + in.Keyword + "%")).Or(qEmployee.NationalID.Like("%" + in.Keyword + "%")).Find()
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

func (e *employee) Update(in model.PutEmployeeReq) (out model.PutEmployeeRes, err error) {
	var (
		qEmployee = query.Employee
	)

	_, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Where(qEmployee.ID.Eq(in.ID)).Updates(in.Employee)
	return
}

func (e *employee) Delete(in model.DeleteEmployeeReq) (out model.DeleteEmployeeRes, err error) {
	var (
		qEmployee = query.Employee
	)

	_, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Unscoped().Preload(qEmployee.LoginInformation).Where(qEmployee.ID.Eq(in.ID)).Delete()
	return
}
