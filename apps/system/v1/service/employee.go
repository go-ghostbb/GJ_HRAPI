package service

import (
	"context"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	gbrand "ghostbb.io/gb/util/gb_rand"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/system/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/paginator"
	"hrapi/internal/utils/password"
	"os"
)

func Employee(ctx context.Context, page ...*paginator.Pagination) IEmployee {
	if len(page) != 0 {
		return &employee{page: page[0], ctx: ctx}
	}
	return &employee{ctx: ctx}
}

type (
	IEmployee interface {
		GetByKeyword(in model.GetByKeywordEmployeeReq) (out model.GetByKeywordEmployeeRes, err error)
		GetByID(in model.GetByIDEmployeeReq) (out model.GetByIDEmployeeRes, err error)
		Insert(in model.PostEmployeeReq) (out model.PostEmployeeRes, err error)
		Update(in model.PutEmployeeReq) (err error)
		Delete(in model.DeleteEmployeeReq) (err error)
		ResetPassword(in model.ResetPasswordReq) (err error)
		SetEmploymentStatus(in model.SetEmploymentStatusReq) (err error)
		SetLoginStatus(in model.SetLoginStatusReq) (err error)
		GetByDateRangeCheckInStatus(in model.GetByDateRangeCheckInStatusReq) (out []model.GetByDateRangeCheckInStatusRes, err error)
		GetByKeywordSalaryAddItem(in model.GetByKeywordSalaryAddItemReq) (out []*model.GetByKeywordSalaryAddItemRes, err error)
		SetAmountSalaryAddItem(in model.SetAmountSalaryAddItemReq) (err error)
	}

	employee struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

func (e *employee) GetByKeyword(in model.GetByKeywordEmployeeReq) (out model.GetByKeywordEmployeeRes, err error) {
	var (
		qEmployee  = query.Employee
		qdEmployee = qEmployee.WithContext(dbcache.WithCtx(e.ctx))
		employees  []*types.Employee
		conds      = []gen.Condition{
			// keyword => where (real_name like ? or national_id like ?)
			qdEmployee.Where(qdEmployee.Where(qEmployee.RealName.Like("%" + in.Keyword + "%")).Or(qEmployee.NationalID.Like("%" + in.Keyword + "%"))),
		}
	)

	if in.EmploymentStatus != "" {
		// status => and employment_status = ?
		conds = append(conds, qEmployee.EmploymentStatus.Eq(in.EmploymentStatus))
	}

	if in.DepartmentId != "" {
		// department_id => and department_id = ?
		conds = append(conds, qEmployee.DepartmentId.Eq(gbconv.Uint(in.DepartmentId)))
	}

	employees, err = qdEmployee.Scopes(e.page.Where(conds...)).Preload(
		qEmployee.LoginInformation,
		qEmployee.Department,
	).Find()

	if err = copier.Copy(&out.Items, employees); err != nil {
		return
	}

	out.Total = e.page.Total

	return
}

func (e *employee) GetByID(in model.GetByIDEmployeeReq) (out model.GetByIDEmployeeRes, err error) {
	qEmployee := query.Employee
	out.Employee, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Preload(
		qEmployee.LoginInformation,
		qEmployee.Department,
	).Where(qEmployee.ID.Eq(in.ID)).First()
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

		// 儲存登入資訊
		err = qLoginInfo.WithContext(dbcache.WithCtx(e.ctx)).Create(loginInfo)
		if err != nil {
			return err
		}

		// 回傳
		out.Account = loginInfo.Account
		out.Password = tmpPass

		// commit
		return nil
	})

	return
}

func (e *employee) Update(in model.PutEmployeeReq) error {
	// 寫入ID
	in.Employee.ID = in.ID

	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qEmployee    = query.Employee
			qDepartment  = query.Department
			originDeptID uint
			err          error
		)

		// 查詢該員工原本的部門ID
		err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Select(qEmployee.DepartmentId).Where(qEmployee.ID.Eq(in.ID)).Scan(&originDeptID)
		if err != nil {
			return err
		}

		// 判斷部門是否有更動
		if originDeptID != in.Employee.DepartmentId {
			// 有更動的話
			// 如果該員工是某部門主管, 將該部門主管id改為0
			_, err = qDepartment.WithContext(dbcache.WithCtx(e.ctx)).Where(qDepartment.ManagerId.Eq(in.ID)).Update(qDepartment.ManagerId, 0)
			if err != nil {
				return err
			}
		}

		// 儲存員工資訊
		err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Save(in.Employee)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

func (e *employee) Delete(in model.DeleteEmployeeReq) (err error) {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qEmployee   = tx.Employee
			qLoginInfo  = tx.LoginInformation
			qDepartment = tx.Department
		)
		// 刪除員工靜態文件資料夾
		if err = os.RemoveAll(fmt.Sprintf("assets/employee/%d", in.ID)); err != nil {
			return err
		}

		// 刪除員工
		_, err = qEmployee.WithContext(dbcache.WithCtx(e.ctx)).Unscoped().Where(qEmployee.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除員工登入資訊
		_, err = qLoginInfo.WithContext(dbcache.WithCtx(e.ctx)).Unscoped().Where(qLoginInfo.EmployeeID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 如果該員工是某部門主管, 將該部門主管id改為0
		_, err = qDepartment.WithContext(dbcache.WithCtx(e.ctx)).Where(qDepartment.ManagerId.Eq(in.ID)).Update(qDepartment.ManagerId, 0)
		if err != nil {
			return err
		}

		return nil
	})
}

func (e *employee) ResetPassword(in model.ResetPasswordReq) (err error) {
	// 輸入驗證
	if err = g.Validator().Data(in).Run(e.ctx); err != nil {
		return err
	}

	var (
		qLoginInfo = query.LoginInformation
	)

	// 進行密碼修改
	_, err = qLoginInfo.WithContext(dbcache.WithCtx(e.ctx)).
		Where(qLoginInfo.EmployeeID.Eq(in.ID)).
		Update(qLoginInfo.Password, password.Hash(in.Password))
	if err != nil {
		return err
	}

	return nil
}

func (e *employee) SetEmploymentStatus(in model.SetEmploymentStatusReq) (err error) {
	_, err = query.Employee.WithContext(dbcache.WithCtx(e.ctx)).Where(query.Employee.ID.Eq(in.ID)).Update(query.Employee.EmploymentStatus, in.EmploymentStatus)
	return err
}

func (e *employee) SetLoginStatus(in model.SetLoginStatusReq) (err error) {
	_, err = query.LoginInformation.WithContext(dbcache.WithCtx(e.ctx)).Where(query.LoginInformation.EmployeeID.Eq(in.ID)).Update(query.LoginInformation.Status, in.Status)
	return err
}

func (e *employee) GetByDateRangeCheckInStatus(in model.GetByDateRangeCheckInStatusReq) (out []model.GetByDateRangeCheckInStatusRes, err error) {
	var (
		qCheckInStatus = query.CheckInStatus
		result         []*types.CheckInStatus
	)

	result, err = qCheckInStatus.WithContext(dbcache.WithCtx(e.ctx)).Preload(field.Associations, qCheckInStatus.Employee.Department).QueryByDateRangeAndEmpID(in.EmployeeID, in.DateRangeStart, in.DateRangeEnd, in.Abnormal)
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(&out, result); err != nil {
		return nil, err
	}

	return out, nil
}

// GetByKeywordSalaryAddItem 查詢員工的加項
func (e *employee) GetByKeywordSalaryAddItem(in model.GetByKeywordSalaryAddItemReq) (out []*model.GetByKeywordSalaryAddItemRes, err error) {
	var (
		qSalaryAddItem = query.SalaryAddItem
		queryRes       []g.Map
	)

	queryRes, err = qSalaryAddItem.WithContext(e.ctx).QueryByEmployeeID(in.EmployeeID, "%"+in.Keyword+"%")
	if err != nil {
		return nil, err
	}

	if err = gbconv.Structs(queryRes, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// SetAmountSalaryAddItem 設定加項金額
func (e *employee) SetAmountSalaryAddItem(in model.SetAmountSalaryAddItemReq) (err error) {
	var (
		qSalaryAddItemEmployee = query.SalaryAddItemEmployee
		saie                   *types.SalaryAddItemEmployee
	)

	saie, err = qSalaryAddItemEmployee.WithContext(dbcache.WithCtx(e.ctx)).
		Where(qSalaryAddItemEmployee.SalaryAddItemID.Eq(in.SalaryAddItemID), qSalaryAddItemEmployee.EmployeeID.Eq(in.EmployeeID)).
		First()
	if err != nil {
		return err
	}

	saie.UseCustom = in.UseCustom
	saie.CustomAmount = in.CustomAmount

	return qSalaryAddItemEmployee.WithContext(dbcache.WithCtx(e.ctx)).Save(saie)
}
