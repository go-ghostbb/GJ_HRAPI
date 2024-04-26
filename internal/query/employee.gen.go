// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"hrapi/internal/types"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newEmployee(db *gorm.DB, opts ...gen.DOOption) employee {
	_employee := employee{}

	_employee.employeeDo.UseDB(db, opts...)
	_employee.employeeDo.UseModel(&types.Employee{})

	tableName := _employee.employeeDo.TableName()
	_employee.ALL = field.NewAsterisk(tableName)
	_employee.ID = field.NewUint(tableName, "id")
	_employee.CreatedAt = field.NewTime(tableName, "created_at")
	_employee.UpdatedAt = field.NewTime(tableName, "updated_at")
	_employee.DeletedAt = field.NewField(tableName, "deleted_at")
	_employee.Code = field.NewString(tableName, "code")
	_employee.CardNumber = field.NewString(tableName, "card_number")
	_employee.HireDate = field.NewField(tableName, "hire_date")
	_employee.TerminationDate = field.NewField(tableName, "termination_date")
	_employee.EmploymentStatus = field.NewField(tableName, "employment_status")
	_employee.Backend = field.NewBool(tableName, "backend")
	_employee.RealName = field.NewString(tableName, "real_name")
	_employee.NationalID = field.NewString(tableName, "national_id")
	_employee.Birth = field.NewTime(tableName, "birth")
	_employee.Email = field.NewString(tableName, "email")
	_employee.Mobile = field.NewString(tableName, "mobile")
	_employee.Avatar = field.NewString(tableName, "avatar")
	_employee.Salary = field.NewFloat32(tableName, "salary")
	_employee.SalaryCycle = field.NewField(tableName, "salary_cycle")
	_employee.DepartmentId = field.NewUint(tableName, "department_id")
	_employee.RankID = field.NewUint(tableName, "rank_id")
	_employee.GradeID = field.NewUint(tableName, "grade_id")
	_employee.LoginInformation = employeeHasOneLoginInformation{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("LoginInformation", "types.LoginInformation"),
		Employee: struct {
			field.RelationField
			Department struct {
				field.RelationField
				Manager struct {
					field.RelationField
				}
			}
			Rank struct {
				field.RelationField
				Grade struct {
					field.RelationField
					Rank struct {
						field.RelationField
					}
				}
			}
			Grade struct {
				field.RelationField
			}
			LoginInformation struct {
				field.RelationField
			}
			Roles struct {
				field.RelationField
				Employees struct {
					field.RelationField
				}
				Permissions struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}
				Menus struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}
			}
		}{
			RelationField: field.NewRelation("LoginInformation.Employee", "types.Employee"),
			Department: struct {
				field.RelationField
				Manager struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("LoginInformation.Employee.Department", "types.Department"),
				Manager: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("LoginInformation.Employee.Department.Manager", "types.Employee"),
				},
			},
			Rank: struct {
				field.RelationField
				Grade struct {
					field.RelationField
					Rank struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("LoginInformation.Employee.Rank", "types.PositionRank"),
				Grade: struct {
					field.RelationField
					Rank struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("LoginInformation.Employee.Rank.Grade", "types.PositionGrade"),
					Rank: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("LoginInformation.Employee.Rank.Grade.Rank", "types.PositionRank"),
					},
				},
			},
			Grade: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("LoginInformation.Employee.Grade", "types.PositionGrade"),
			},
			LoginInformation: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("LoginInformation.Employee.LoginInformation", "types.LoginInformation"),
			},
			Roles: struct {
				field.RelationField
				Employees struct {
					field.RelationField
				}
				Permissions struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}
				Menus struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("LoginInformation.Employee.Roles", "types.Role"),
				Employees: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("LoginInformation.Employee.Roles.Employees", "types.Employee"),
				},
				Permissions: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("LoginInformation.Employee.Roles.Permissions", "types.Permission"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("LoginInformation.Employee.Roles.Permissions.Roles", "types.Role"),
					},
				},
				Menus: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("LoginInformation.Employee.Roles.Menus", "types.Menu"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("LoginInformation.Employee.Roles.Menus.Roles", "types.Role"),
					},
				},
			},
		},
	}

	_employee.Department = employeeBelongsToDepartment{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Department", "types.Department"),
	}

	_employee.Rank = employeeBelongsToRank{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Rank", "types.PositionRank"),
	}

	_employee.Grade = employeeBelongsToGrade{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Grade", "types.PositionGrade"),
	}

	_employee.Roles = employeeManyToManyRoles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Roles", "types.Role"),
	}

	_employee.fillFieldMap()

	return _employee
}

type employee struct {
	employeeDo employeeDo

	ALL              field.Asterisk
	ID               field.Uint
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Code             field.String
	CardNumber       field.String
	HireDate         field.Field
	TerminationDate  field.Field
	EmploymentStatus field.Field
	Backend          field.Bool
	RealName         field.String
	NationalID       field.String
	Birth            field.Time
	Email            field.String
	Mobile           field.String
	Avatar           field.String
	Salary           field.Float32
	SalaryCycle      field.Field
	DepartmentId     field.Uint
	RankID           field.Uint
	GradeID          field.Uint
	LoginInformation employeeHasOneLoginInformation

	Department employeeBelongsToDepartment

	Rank employeeBelongsToRank

	Grade employeeBelongsToGrade

	Roles employeeManyToManyRoles

	fieldMap map[string]field.Expr
}

func (e employee) Table(newTableName string) *employee {
	e.employeeDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e employee) As(alias string) *employee {
	e.employeeDo.DO = *(e.employeeDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *employee) updateTableName(table string) *employee {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewUint(table, "id")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")
	e.Code = field.NewString(table, "code")
	e.CardNumber = field.NewString(table, "card_number")
	e.HireDate = field.NewField(table, "hire_date")
	e.TerminationDate = field.NewField(table, "termination_date")
	e.EmploymentStatus = field.NewField(table, "employment_status")
	e.Backend = field.NewBool(table, "backend")
	e.RealName = field.NewString(table, "real_name")
	e.NationalID = field.NewString(table, "national_id")
	e.Birth = field.NewTime(table, "birth")
	e.Email = field.NewString(table, "email")
	e.Mobile = field.NewString(table, "mobile")
	e.Avatar = field.NewString(table, "avatar")
	e.Salary = field.NewFloat32(table, "salary")
	e.SalaryCycle = field.NewField(table, "salary_cycle")
	e.DepartmentId = field.NewUint(table, "department_id")
	e.RankID = field.NewUint(table, "rank_id")
	e.GradeID = field.NewUint(table, "grade_id")

	e.fillFieldMap()

	return e
}

func (e *employee) WithContext(ctx context.Context) IEmployeeDo { return e.employeeDo.WithContext(ctx) }

func (e employee) TableName() string { return e.employeeDo.TableName() }

func (e employee) Alias() string { return e.employeeDo.Alias() }

func (e employee) Columns(cols ...field.Expr) gen.Columns { return e.employeeDo.Columns(cols...) }

func (e *employee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *employee) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 26)
	e.fieldMap["id"] = e.ID
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
	e.fieldMap["code"] = e.Code
	e.fieldMap["card_number"] = e.CardNumber
	e.fieldMap["hire_date"] = e.HireDate
	e.fieldMap["termination_date"] = e.TerminationDate
	e.fieldMap["employment_status"] = e.EmploymentStatus
	e.fieldMap["backend"] = e.Backend
	e.fieldMap["real_name"] = e.RealName
	e.fieldMap["national_id"] = e.NationalID
	e.fieldMap["birth"] = e.Birth
	e.fieldMap["email"] = e.Email
	e.fieldMap["mobile"] = e.Mobile
	e.fieldMap["avatar"] = e.Avatar
	e.fieldMap["salary"] = e.Salary
	e.fieldMap["salary_cycle"] = e.SalaryCycle
	e.fieldMap["department_id"] = e.DepartmentId
	e.fieldMap["rank_id"] = e.RankID
	e.fieldMap["grade_id"] = e.GradeID

}

func (e employee) clone(db *gorm.DB) employee {
	e.employeeDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e employee) replaceDB(db *gorm.DB) employee {
	e.employeeDo.ReplaceDB(db)
	return e
}

type employeeHasOneLoginInformation struct {
	db *gorm.DB

	field.RelationField

	Employee struct {
		field.RelationField
		Department struct {
			field.RelationField
			Manager struct {
				field.RelationField
			}
		}
		Rank struct {
			field.RelationField
			Grade struct {
				field.RelationField
				Rank struct {
					field.RelationField
				}
			}
		}
		Grade struct {
			field.RelationField
		}
		LoginInformation struct {
			field.RelationField
		}
		Roles struct {
			field.RelationField
			Employees struct {
				field.RelationField
			}
			Permissions struct {
				field.RelationField
				Roles struct {
					field.RelationField
				}
			}
			Menus struct {
				field.RelationField
				Roles struct {
					field.RelationField
				}
			}
		}
	}
}

func (a employeeHasOneLoginInformation) Where(conds ...field.Expr) *employeeHasOneLoginInformation {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a employeeHasOneLoginInformation) WithContext(ctx context.Context) *employeeHasOneLoginInformation {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a employeeHasOneLoginInformation) Session(session *gorm.Session) *employeeHasOneLoginInformation {
	a.db = a.db.Session(session)
	return &a
}

func (a employeeHasOneLoginInformation) Model(m *types.Employee) *employeeHasOneLoginInformationTx {
	return &employeeHasOneLoginInformationTx{a.db.Model(m).Association(a.Name())}
}

type employeeHasOneLoginInformationTx struct{ tx *gorm.Association }

func (a employeeHasOneLoginInformationTx) Find() (result *types.LoginInformation, err error) {
	return result, a.tx.Find(&result)
}

func (a employeeHasOneLoginInformationTx) Append(values ...*types.LoginInformation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a employeeHasOneLoginInformationTx) Replace(values ...*types.LoginInformation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a employeeHasOneLoginInformationTx) Delete(values ...*types.LoginInformation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a employeeHasOneLoginInformationTx) Clear() error {
	return a.tx.Clear()
}

func (a employeeHasOneLoginInformationTx) Count() int64 {
	return a.tx.Count()
}

type employeeBelongsToDepartment struct {
	db *gorm.DB

	field.RelationField
}

func (a employeeBelongsToDepartment) Where(conds ...field.Expr) *employeeBelongsToDepartment {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a employeeBelongsToDepartment) WithContext(ctx context.Context) *employeeBelongsToDepartment {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a employeeBelongsToDepartment) Session(session *gorm.Session) *employeeBelongsToDepartment {
	a.db = a.db.Session(session)
	return &a
}

func (a employeeBelongsToDepartment) Model(m *types.Employee) *employeeBelongsToDepartmentTx {
	return &employeeBelongsToDepartmentTx{a.db.Model(m).Association(a.Name())}
}

type employeeBelongsToDepartmentTx struct{ tx *gorm.Association }

func (a employeeBelongsToDepartmentTx) Find() (result *types.Department, err error) {
	return result, a.tx.Find(&result)
}

func (a employeeBelongsToDepartmentTx) Append(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a employeeBelongsToDepartmentTx) Replace(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a employeeBelongsToDepartmentTx) Delete(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a employeeBelongsToDepartmentTx) Clear() error {
	return a.tx.Clear()
}

func (a employeeBelongsToDepartmentTx) Count() int64 {
	return a.tx.Count()
}

type employeeBelongsToRank struct {
	db *gorm.DB

	field.RelationField
}

func (a employeeBelongsToRank) Where(conds ...field.Expr) *employeeBelongsToRank {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a employeeBelongsToRank) WithContext(ctx context.Context) *employeeBelongsToRank {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a employeeBelongsToRank) Session(session *gorm.Session) *employeeBelongsToRank {
	a.db = a.db.Session(session)
	return &a
}

func (a employeeBelongsToRank) Model(m *types.Employee) *employeeBelongsToRankTx {
	return &employeeBelongsToRankTx{a.db.Model(m).Association(a.Name())}
}

type employeeBelongsToRankTx struct{ tx *gorm.Association }

func (a employeeBelongsToRankTx) Find() (result *types.PositionRank, err error) {
	return result, a.tx.Find(&result)
}

func (a employeeBelongsToRankTx) Append(values ...*types.PositionRank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a employeeBelongsToRankTx) Replace(values ...*types.PositionRank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a employeeBelongsToRankTx) Delete(values ...*types.PositionRank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a employeeBelongsToRankTx) Clear() error {
	return a.tx.Clear()
}

func (a employeeBelongsToRankTx) Count() int64 {
	return a.tx.Count()
}

type employeeBelongsToGrade struct {
	db *gorm.DB

	field.RelationField
}

func (a employeeBelongsToGrade) Where(conds ...field.Expr) *employeeBelongsToGrade {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a employeeBelongsToGrade) WithContext(ctx context.Context) *employeeBelongsToGrade {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a employeeBelongsToGrade) Session(session *gorm.Session) *employeeBelongsToGrade {
	a.db = a.db.Session(session)
	return &a
}

func (a employeeBelongsToGrade) Model(m *types.Employee) *employeeBelongsToGradeTx {
	return &employeeBelongsToGradeTx{a.db.Model(m).Association(a.Name())}
}

type employeeBelongsToGradeTx struct{ tx *gorm.Association }

func (a employeeBelongsToGradeTx) Find() (result *types.PositionGrade, err error) {
	return result, a.tx.Find(&result)
}

func (a employeeBelongsToGradeTx) Append(values ...*types.PositionGrade) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a employeeBelongsToGradeTx) Replace(values ...*types.PositionGrade) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a employeeBelongsToGradeTx) Delete(values ...*types.PositionGrade) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a employeeBelongsToGradeTx) Clear() error {
	return a.tx.Clear()
}

func (a employeeBelongsToGradeTx) Count() int64 {
	return a.tx.Count()
}

type employeeManyToManyRoles struct {
	db *gorm.DB

	field.RelationField
}

func (a employeeManyToManyRoles) Where(conds ...field.Expr) *employeeManyToManyRoles {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a employeeManyToManyRoles) WithContext(ctx context.Context) *employeeManyToManyRoles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a employeeManyToManyRoles) Session(session *gorm.Session) *employeeManyToManyRoles {
	a.db = a.db.Session(session)
	return &a
}

func (a employeeManyToManyRoles) Model(m *types.Employee) *employeeManyToManyRolesTx {
	return &employeeManyToManyRolesTx{a.db.Model(m).Association(a.Name())}
}

type employeeManyToManyRolesTx struct{ tx *gorm.Association }

func (a employeeManyToManyRolesTx) Find() (result []*types.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a employeeManyToManyRolesTx) Append(values ...*types.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a employeeManyToManyRolesTx) Replace(values ...*types.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a employeeManyToManyRolesTx) Delete(values ...*types.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a employeeManyToManyRolesTx) Clear() error {
	return a.tx.Clear()
}

func (a employeeManyToManyRolesTx) Count() int64 {
	return a.tx.Count()
}

type employeeDo struct{ gen.DO }

type IEmployeeDo interface {
	gen.SubQuery
	Debug() IEmployeeDo
	WithContext(ctx context.Context) IEmployeeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEmployeeDo
	WriteDB() IEmployeeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEmployeeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEmployeeDo
	Not(conds ...gen.Condition) IEmployeeDo
	Or(conds ...gen.Condition) IEmployeeDo
	Select(conds ...field.Expr) IEmployeeDo
	Where(conds ...gen.Condition) IEmployeeDo
	Order(conds ...field.Expr) IEmployeeDo
	Distinct(cols ...field.Expr) IEmployeeDo
	Omit(cols ...field.Expr) IEmployeeDo
	Join(table schema.Tabler, on ...field.Expr) IEmployeeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo
	Group(cols ...field.Expr) IEmployeeDo
	Having(conds ...gen.Condition) IEmployeeDo
	Limit(limit int) IEmployeeDo
	Offset(offset int) IEmployeeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeDo
	Unscoped() IEmployeeDo
	Create(values ...*types.Employee) error
	CreateInBatches(values []*types.Employee, batchSize int) error
	Save(values ...*types.Employee) error
	First() (*types.Employee, error)
	Take() (*types.Employee, error)
	Last() (*types.Employee, error)
	Find() ([]*types.Employee, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.Employee, err error)
	FindInBatches(result *[]*types.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.Employee) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEmployeeDo
	Assign(attrs ...field.AssignExpr) IEmployeeDo
	Joins(fields ...field.RelationField) IEmployeeDo
	Preload(fields ...field.RelationField) IEmployeeDo
	FirstOrInit() (*types.Employee, error)
	FirstOrCreate() (*types.Employee, error)
	FindByPage(offset int, limit int) (result []*types.Employee, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEmployeeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	QueryIDWhereCardNum(cardNums []string) (result []map[string]interface{}, err error)
	CalculateEmployeeSeniority() (result []map[string]interface{}, err error)
	CalculateEmployeeSeniorityWithEndDate(dateOnly string) (result []map[string]interface{}, err error)
}

// select id, card_number from @@table where card_number in (@cardNums);
func (e employeeDo) QueryIDWhereCardNum(cardNums []string) (result []map[string]interface{}, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, cardNums)
	generateSQL.WriteString("select id, card_number from employee where card_number in (?); ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select * from FN_C_EmployeeSeniority()
func (e employeeDo) CalculateEmployeeSeniority() (result []map[string]interface{}, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select * from FN_C_EmployeeSeniority() ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select * from FN_C_EmployeeSeniorityWithEndDate(@dateOnly)
func (e employeeDo) CalculateEmployeeSeniorityWithEndDate(dateOnly string) (result []map[string]interface{}, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, dateOnly)
	generateSQL.WriteString("select * from FN_C_EmployeeSeniorityWithEndDate(?) ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (e employeeDo) Debug() IEmployeeDo {
	return e.withDO(e.DO.Debug())
}

func (e employeeDo) WithContext(ctx context.Context) IEmployeeDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e employeeDo) ReadDB() IEmployeeDo {
	return e.Clauses(dbresolver.Read)
}

func (e employeeDo) WriteDB() IEmployeeDo {
	return e.Clauses(dbresolver.Write)
}

func (e employeeDo) Session(config *gorm.Session) IEmployeeDo {
	return e.withDO(e.DO.Session(config))
}

func (e employeeDo) Clauses(conds ...clause.Expression) IEmployeeDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e employeeDo) Returning(value interface{}, columns ...string) IEmployeeDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e employeeDo) Not(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e employeeDo) Or(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e employeeDo) Select(conds ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e employeeDo) Where(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e employeeDo) Order(conds ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e employeeDo) Distinct(cols ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e employeeDo) Omit(cols ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e employeeDo) Join(table schema.Tabler, on ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e employeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e employeeDo) RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e employeeDo) Group(cols ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e employeeDo) Having(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e employeeDo) Limit(limit int) IEmployeeDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e employeeDo) Offset(offset int) IEmployeeDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e employeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e employeeDo) Unscoped() IEmployeeDo {
	return e.withDO(e.DO.Unscoped())
}

func (e employeeDo) Create(values ...*types.Employee) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e employeeDo) CreateInBatches(values []*types.Employee, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e employeeDo) Save(values ...*types.Employee) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e employeeDo) First() (*types.Employee, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.Employee), nil
	}
}

func (e employeeDo) Take() (*types.Employee, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.Employee), nil
	}
}

func (e employeeDo) Last() (*types.Employee, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.Employee), nil
	}
}

func (e employeeDo) Find() ([]*types.Employee, error) {
	result, err := e.DO.Find()
	return result.([]*types.Employee), err
}

func (e employeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.Employee, err error) {
	buf := make([]*types.Employee, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e employeeDo) FindInBatches(result *[]*types.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e employeeDo) Attrs(attrs ...field.AssignExpr) IEmployeeDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e employeeDo) Assign(attrs ...field.AssignExpr) IEmployeeDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e employeeDo) Joins(fields ...field.RelationField) IEmployeeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e employeeDo) Preload(fields ...field.RelationField) IEmployeeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e employeeDo) FirstOrInit() (*types.Employee, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.Employee), nil
	}
}

func (e employeeDo) FirstOrCreate() (*types.Employee, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.Employee), nil
	}
}

func (e employeeDo) FindByPage(offset int, limit int) (result []*types.Employee, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e employeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e employeeDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e employeeDo) Delete(models ...*types.Employee) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *employeeDo) withDO(do gen.Dao) *employeeDo {
	e.DO = *do.(*gen.DO)
	return e
}
