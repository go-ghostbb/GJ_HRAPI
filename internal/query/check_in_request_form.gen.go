// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"hrapi/internal/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCheckInRequestForm(db *gorm.DB, opts ...gen.DOOption) checkInRequestForm {
	_checkInRequestForm := checkInRequestForm{}

	_checkInRequestForm.checkInRequestFormDo.UseDB(db, opts...)
	_checkInRequestForm.checkInRequestFormDo.UseModel(&types.CheckInRequestForm{})

	tableName := _checkInRequestForm.checkInRequestFormDo.TableName()
	_checkInRequestForm.ALL = field.NewAsterisk(tableName)
	_checkInRequestForm.ID = field.NewUint(tableName, "id")
	_checkInRequestForm.CreatedAt = field.NewTime(tableName, "created_at")
	_checkInRequestForm.UpdatedAt = field.NewTime(tableName, "updated_at")
	_checkInRequestForm.DeletedAt = field.NewField(tableName, "deleted_at")
	_checkInRequestForm.Order = field.NewString(tableName, "order")
	_checkInRequestForm.SignStatus = field.NewField(tableName, "sign_status")
	_checkInRequestForm.Attach = field.NewString(tableName, "attach")
	_checkInRequestForm.Remark = field.NewString(tableName, "remark")
	_checkInRequestForm.EmployeeID = field.NewUint(tableName, "employee_id")
	_checkInRequestForm.DepartmentID = field.NewUint(tableName, "department_id")
	_checkInRequestForm.SignOffFlow = checkInRequestFormHasManySignOffFlow{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SignOffFlow", "types.CheckInSignOffFlow"),
		CheckInRequestForm: struct {
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
					Employee struct {
						field.RelationField
					}
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
			Department struct {
				field.RelationField
			}
			SignOffFlow struct {
				field.RelationField
			}
			Detail struct {
				field.RelationField
				CheckInRequestForm struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm", "types.CheckInRequestForm"),
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
					Employee struct {
						field.RelationField
					}
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
				RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee", "types.Employee"),
				Department: struct {
					field.RelationField
					Manager struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Department", "types.Department"),
					Manager: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Department.Manager", "types.Employee"),
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
					RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Rank", "types.PositionRank"),
					Grade: struct {
						field.RelationField
						Rank struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Rank.Grade", "types.PositionGrade"),
						Rank: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Rank.Grade.Rank", "types.PositionRank"),
						},
					},
				},
				Grade: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Grade", "types.PositionGrade"),
				},
				LoginInformation: struct {
					field.RelationField
					Employee struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.LoginInformation", "types.LoginInformation"),
					Employee: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.LoginInformation.Employee", "types.Employee"),
					},
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
					RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Roles", "types.Role"),
					Employees: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Roles.Employees", "types.Employee"),
					},
					Permissions: struct {
						field.RelationField
						Roles struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Roles.Permissions", "types.Permission"),
						Roles: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Roles.Permissions.Roles", "types.Role"),
						},
					},
					Menus: struct {
						field.RelationField
						Roles struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Roles.Menus", "types.Menu"),
						Roles: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Employee.Roles.Menus.Roles", "types.Role"),
						},
					},
				},
			},
			Department: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Department", "types.Department"),
			},
			SignOffFlow: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.SignOffFlow", "types.CheckInSignOffFlow"),
			},
			Detail: struct {
				field.RelationField
				CheckInRequestForm struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Detail", "types.CheckInRequestFormDetail"),
				CheckInRequestForm: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SignOffFlow.CheckInRequestForm.Detail.CheckInRequestForm", "types.CheckInRequestForm"),
				},
			},
		},
		SignOffEmployee: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("SignOffFlow.SignOffEmployee", "types.Employee"),
		},
	}

	_checkInRequestForm.Detail = checkInRequestFormHasManyDetail{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Detail", "types.CheckInRequestFormDetail"),
	}

	_checkInRequestForm.Employee = checkInRequestFormBelongsToEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Employee", "types.Employee"),
	}

	_checkInRequestForm.Department = checkInRequestFormBelongsToDepartment{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Department", "types.Department"),
	}

	_checkInRequestForm.fillFieldMap()

	return _checkInRequestForm
}

type checkInRequestForm struct {
	checkInRequestFormDo checkInRequestFormDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	Order        field.String
	SignStatus   field.Field
	Attach       field.String
	Remark       field.String
	EmployeeID   field.Uint
	DepartmentID field.Uint
	SignOffFlow  checkInRequestFormHasManySignOffFlow

	Detail checkInRequestFormHasManyDetail

	Employee checkInRequestFormBelongsToEmployee

	Department checkInRequestFormBelongsToDepartment

	fieldMap map[string]field.Expr
}

func (c checkInRequestForm) Table(newTableName string) *checkInRequestForm {
	c.checkInRequestFormDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c checkInRequestForm) As(alias string) *checkInRequestForm {
	c.checkInRequestFormDo.DO = *(c.checkInRequestFormDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *checkInRequestForm) updateTableName(table string) *checkInRequestForm {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Order = field.NewString(table, "order")
	c.SignStatus = field.NewField(table, "sign_status")
	c.Attach = field.NewString(table, "attach")
	c.Remark = field.NewString(table, "remark")
	c.EmployeeID = field.NewUint(table, "employee_id")
	c.DepartmentID = field.NewUint(table, "department_id")

	c.fillFieldMap()

	return c
}

func (c *checkInRequestForm) WithContext(ctx context.Context) ICheckInRequestFormDo {
	return c.checkInRequestFormDo.WithContext(ctx)
}

func (c checkInRequestForm) TableName() string { return c.checkInRequestFormDo.TableName() }

func (c checkInRequestForm) Alias() string { return c.checkInRequestFormDo.Alias() }

func (c checkInRequestForm) Columns(cols ...field.Expr) gen.Columns {
	return c.checkInRequestFormDo.Columns(cols...)
}

func (c *checkInRequestForm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *checkInRequestForm) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["order"] = c.Order
	c.fieldMap["sign_status"] = c.SignStatus
	c.fieldMap["attach"] = c.Attach
	c.fieldMap["remark"] = c.Remark
	c.fieldMap["employee_id"] = c.EmployeeID
	c.fieldMap["department_id"] = c.DepartmentID

}

func (c checkInRequestForm) clone(db *gorm.DB) checkInRequestForm {
	c.checkInRequestFormDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c checkInRequestForm) replaceDB(db *gorm.DB) checkInRequestForm {
	c.checkInRequestFormDo.ReplaceDB(db)
	return c
}

type checkInRequestFormHasManySignOffFlow struct {
	db *gorm.DB

	field.RelationField

	CheckInRequestForm struct {
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
				Employee struct {
					field.RelationField
				}
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
		Department struct {
			field.RelationField
		}
		SignOffFlow struct {
			field.RelationField
		}
		Detail struct {
			field.RelationField
			CheckInRequestForm struct {
				field.RelationField
			}
		}
	}
	SignOffEmployee struct {
		field.RelationField
	}
}

func (a checkInRequestFormHasManySignOffFlow) Where(conds ...field.Expr) *checkInRequestFormHasManySignOffFlow {
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

func (a checkInRequestFormHasManySignOffFlow) WithContext(ctx context.Context) *checkInRequestFormHasManySignOffFlow {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a checkInRequestFormHasManySignOffFlow) Session(session *gorm.Session) *checkInRequestFormHasManySignOffFlow {
	a.db = a.db.Session(session)
	return &a
}

func (a checkInRequestFormHasManySignOffFlow) Model(m *types.CheckInRequestForm) *checkInRequestFormHasManySignOffFlowTx {
	return &checkInRequestFormHasManySignOffFlowTx{a.db.Model(m).Association(a.Name())}
}

type checkInRequestFormHasManySignOffFlowTx struct{ tx *gorm.Association }

func (a checkInRequestFormHasManySignOffFlowTx) Find() (result []*types.CheckInSignOffFlow, err error) {
	return result, a.tx.Find(&result)
}

func (a checkInRequestFormHasManySignOffFlowTx) Append(values ...*types.CheckInSignOffFlow) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a checkInRequestFormHasManySignOffFlowTx) Replace(values ...*types.CheckInSignOffFlow) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a checkInRequestFormHasManySignOffFlowTx) Delete(values ...*types.CheckInSignOffFlow) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a checkInRequestFormHasManySignOffFlowTx) Clear() error {
	return a.tx.Clear()
}

func (a checkInRequestFormHasManySignOffFlowTx) Count() int64 {
	return a.tx.Count()
}

type checkInRequestFormHasManyDetail struct {
	db *gorm.DB

	field.RelationField
}

func (a checkInRequestFormHasManyDetail) Where(conds ...field.Expr) *checkInRequestFormHasManyDetail {
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

func (a checkInRequestFormHasManyDetail) WithContext(ctx context.Context) *checkInRequestFormHasManyDetail {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a checkInRequestFormHasManyDetail) Session(session *gorm.Session) *checkInRequestFormHasManyDetail {
	a.db = a.db.Session(session)
	return &a
}

func (a checkInRequestFormHasManyDetail) Model(m *types.CheckInRequestForm) *checkInRequestFormHasManyDetailTx {
	return &checkInRequestFormHasManyDetailTx{a.db.Model(m).Association(a.Name())}
}

type checkInRequestFormHasManyDetailTx struct{ tx *gorm.Association }

func (a checkInRequestFormHasManyDetailTx) Find() (result []*types.CheckInRequestFormDetail, err error) {
	return result, a.tx.Find(&result)
}

func (a checkInRequestFormHasManyDetailTx) Append(values ...*types.CheckInRequestFormDetail) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a checkInRequestFormHasManyDetailTx) Replace(values ...*types.CheckInRequestFormDetail) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a checkInRequestFormHasManyDetailTx) Delete(values ...*types.CheckInRequestFormDetail) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a checkInRequestFormHasManyDetailTx) Clear() error {
	return a.tx.Clear()
}

func (a checkInRequestFormHasManyDetailTx) Count() int64 {
	return a.tx.Count()
}

type checkInRequestFormBelongsToEmployee struct {
	db *gorm.DB

	field.RelationField
}

func (a checkInRequestFormBelongsToEmployee) Where(conds ...field.Expr) *checkInRequestFormBelongsToEmployee {
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

func (a checkInRequestFormBelongsToEmployee) WithContext(ctx context.Context) *checkInRequestFormBelongsToEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a checkInRequestFormBelongsToEmployee) Session(session *gorm.Session) *checkInRequestFormBelongsToEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a checkInRequestFormBelongsToEmployee) Model(m *types.CheckInRequestForm) *checkInRequestFormBelongsToEmployeeTx {
	return &checkInRequestFormBelongsToEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type checkInRequestFormBelongsToEmployeeTx struct{ tx *gorm.Association }

func (a checkInRequestFormBelongsToEmployeeTx) Find() (result *types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a checkInRequestFormBelongsToEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a checkInRequestFormBelongsToEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a checkInRequestFormBelongsToEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a checkInRequestFormBelongsToEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a checkInRequestFormBelongsToEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type checkInRequestFormBelongsToDepartment struct {
	db *gorm.DB

	field.RelationField
}

func (a checkInRequestFormBelongsToDepartment) Where(conds ...field.Expr) *checkInRequestFormBelongsToDepartment {
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

func (a checkInRequestFormBelongsToDepartment) WithContext(ctx context.Context) *checkInRequestFormBelongsToDepartment {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a checkInRequestFormBelongsToDepartment) Session(session *gorm.Session) *checkInRequestFormBelongsToDepartment {
	a.db = a.db.Session(session)
	return &a
}

func (a checkInRequestFormBelongsToDepartment) Model(m *types.CheckInRequestForm) *checkInRequestFormBelongsToDepartmentTx {
	return &checkInRequestFormBelongsToDepartmentTx{a.db.Model(m).Association(a.Name())}
}

type checkInRequestFormBelongsToDepartmentTx struct{ tx *gorm.Association }

func (a checkInRequestFormBelongsToDepartmentTx) Find() (result *types.Department, err error) {
	return result, a.tx.Find(&result)
}

func (a checkInRequestFormBelongsToDepartmentTx) Append(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a checkInRequestFormBelongsToDepartmentTx) Replace(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a checkInRequestFormBelongsToDepartmentTx) Delete(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a checkInRequestFormBelongsToDepartmentTx) Clear() error {
	return a.tx.Clear()
}

func (a checkInRequestFormBelongsToDepartmentTx) Count() int64 {
	return a.tx.Count()
}

type checkInRequestFormDo struct{ gen.DO }

type ICheckInRequestFormDo interface {
	gen.SubQuery
	Debug() ICheckInRequestFormDo
	WithContext(ctx context.Context) ICheckInRequestFormDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICheckInRequestFormDo
	WriteDB() ICheckInRequestFormDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICheckInRequestFormDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICheckInRequestFormDo
	Not(conds ...gen.Condition) ICheckInRequestFormDo
	Or(conds ...gen.Condition) ICheckInRequestFormDo
	Select(conds ...field.Expr) ICheckInRequestFormDo
	Where(conds ...gen.Condition) ICheckInRequestFormDo
	Order(conds ...field.Expr) ICheckInRequestFormDo
	Distinct(cols ...field.Expr) ICheckInRequestFormDo
	Omit(cols ...field.Expr) ICheckInRequestFormDo
	Join(table schema.Tabler, on ...field.Expr) ICheckInRequestFormDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICheckInRequestFormDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICheckInRequestFormDo
	Group(cols ...field.Expr) ICheckInRequestFormDo
	Having(conds ...gen.Condition) ICheckInRequestFormDo
	Limit(limit int) ICheckInRequestFormDo
	Offset(offset int) ICheckInRequestFormDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICheckInRequestFormDo
	Unscoped() ICheckInRequestFormDo
	Create(values ...*types.CheckInRequestForm) error
	CreateInBatches(values []*types.CheckInRequestForm, batchSize int) error
	Save(values ...*types.CheckInRequestForm) error
	First() (*types.CheckInRequestForm, error)
	Take() (*types.CheckInRequestForm, error)
	Last() (*types.CheckInRequestForm, error)
	Find() ([]*types.CheckInRequestForm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.CheckInRequestForm, err error)
	FindInBatches(result *[]*types.CheckInRequestForm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.CheckInRequestForm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICheckInRequestFormDo
	Assign(attrs ...field.AssignExpr) ICheckInRequestFormDo
	Joins(fields ...field.RelationField) ICheckInRequestFormDo
	Preload(fields ...field.RelationField) ICheckInRequestFormDo
	FirstOrInit() (*types.CheckInRequestForm, error)
	FirstOrCreate() (*types.CheckInRequestForm, error)
	FindByPage(offset int, limit int) (result []*types.CheckInRequestForm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICheckInRequestFormDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c checkInRequestFormDo) Debug() ICheckInRequestFormDo {
	return c.withDO(c.DO.Debug())
}

func (c checkInRequestFormDo) WithContext(ctx context.Context) ICheckInRequestFormDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c checkInRequestFormDo) ReadDB() ICheckInRequestFormDo {
	return c.Clauses(dbresolver.Read)
}

func (c checkInRequestFormDo) WriteDB() ICheckInRequestFormDo {
	return c.Clauses(dbresolver.Write)
}

func (c checkInRequestFormDo) Session(config *gorm.Session) ICheckInRequestFormDo {
	return c.withDO(c.DO.Session(config))
}

func (c checkInRequestFormDo) Clauses(conds ...clause.Expression) ICheckInRequestFormDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c checkInRequestFormDo) Returning(value interface{}, columns ...string) ICheckInRequestFormDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c checkInRequestFormDo) Not(conds ...gen.Condition) ICheckInRequestFormDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c checkInRequestFormDo) Or(conds ...gen.Condition) ICheckInRequestFormDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c checkInRequestFormDo) Select(conds ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c checkInRequestFormDo) Where(conds ...gen.Condition) ICheckInRequestFormDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c checkInRequestFormDo) Order(conds ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c checkInRequestFormDo) Distinct(cols ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c checkInRequestFormDo) Omit(cols ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c checkInRequestFormDo) Join(table schema.Tabler, on ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c checkInRequestFormDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c checkInRequestFormDo) RightJoin(table schema.Tabler, on ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c checkInRequestFormDo) Group(cols ...field.Expr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c checkInRequestFormDo) Having(conds ...gen.Condition) ICheckInRequestFormDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c checkInRequestFormDo) Limit(limit int) ICheckInRequestFormDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c checkInRequestFormDo) Offset(offset int) ICheckInRequestFormDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c checkInRequestFormDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICheckInRequestFormDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c checkInRequestFormDo) Unscoped() ICheckInRequestFormDo {
	return c.withDO(c.DO.Unscoped())
}

func (c checkInRequestFormDo) Create(values ...*types.CheckInRequestForm) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c checkInRequestFormDo) CreateInBatches(values []*types.CheckInRequestForm, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c checkInRequestFormDo) Save(values ...*types.CheckInRequestForm) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c checkInRequestFormDo) First() (*types.CheckInRequestForm, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.CheckInRequestForm), nil
	}
}

func (c checkInRequestFormDo) Take() (*types.CheckInRequestForm, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.CheckInRequestForm), nil
	}
}

func (c checkInRequestFormDo) Last() (*types.CheckInRequestForm, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.CheckInRequestForm), nil
	}
}

func (c checkInRequestFormDo) Find() ([]*types.CheckInRequestForm, error) {
	result, err := c.DO.Find()
	return result.([]*types.CheckInRequestForm), err
}

func (c checkInRequestFormDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.CheckInRequestForm, err error) {
	buf := make([]*types.CheckInRequestForm, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c checkInRequestFormDo) FindInBatches(result *[]*types.CheckInRequestForm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c checkInRequestFormDo) Attrs(attrs ...field.AssignExpr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c checkInRequestFormDo) Assign(attrs ...field.AssignExpr) ICheckInRequestFormDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c checkInRequestFormDo) Joins(fields ...field.RelationField) ICheckInRequestFormDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c checkInRequestFormDo) Preload(fields ...field.RelationField) ICheckInRequestFormDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c checkInRequestFormDo) FirstOrInit() (*types.CheckInRequestForm, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.CheckInRequestForm), nil
	}
}

func (c checkInRequestFormDo) FirstOrCreate() (*types.CheckInRequestForm, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.CheckInRequestForm), nil
	}
}

func (c checkInRequestFormDo) FindByPage(offset int, limit int) (result []*types.CheckInRequestForm, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c checkInRequestFormDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c checkInRequestFormDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c checkInRequestFormDo) Delete(models ...*types.CheckInRequestForm) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *checkInRequestFormDo) withDO(do gen.Dao) *checkInRequestFormDo {
	c.DO = *do.(*gen.DO)
	return c
}
