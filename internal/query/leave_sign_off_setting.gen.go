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

func newLeaveSignOffSetting(db *gorm.DB, opts ...gen.DOOption) leaveSignOffSetting {
	_leaveSignOffSetting := leaveSignOffSetting{}

	_leaveSignOffSetting.leaveSignOffSettingDo.UseDB(db, opts...)
	_leaveSignOffSetting.leaveSignOffSettingDo.UseModel(&types.LeaveSignOffSetting{})

	tableName := _leaveSignOffSetting.leaveSignOffSettingDo.TableName()
	_leaveSignOffSetting.ALL = field.NewAsterisk(tableName)
	_leaveSignOffSetting.ID = field.NewUint(tableName, "id")
	_leaveSignOffSetting.CreatedAt = field.NewTime(tableName, "created_at")
	_leaveSignOffSetting.UpdatedAt = field.NewTime(tableName, "updated_at")
	_leaveSignOffSetting.DeletedAt = field.NewField(tableName, "deleted_at")
	_leaveSignOffSetting.DepartmentID = field.NewUint(tableName, "department_id")
	_leaveSignOffSetting.LeaveID = field.NewUint(tableName, "leave_id")
	_leaveSignOffSetting.Level = field.NewUint(tableName, "level")
	_leaveSignOffSetting.GteDay = field.NewFloat32(tableName, "gte_day")
	_leaveSignOffSetting.SignType = field.NewField(tableName, "sign_type")
	_leaveSignOffSetting.SpecificEmployeeID = field.NewUint(tableName, "specific_employee_id")
	_leaveSignOffSetting.Notify = field.NewField(tableName, "notify")
	_leaveSignOffSetting.Remark = field.NewString(tableName, "remark")
	_leaveSignOffSetting.Department = leaveSignOffSettingBelongsToDepartment{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Department", "types.Department"),
		Manager: struct {
			field.RelationField
			Department struct {
				field.RelationField
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
			RelationField: field.NewRelation("Department.Manager", "types.Employee"),
			Department: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Department.Manager.Department", "types.Department"),
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
				RelationField: field.NewRelation("Department.Manager.Rank", "types.PositionRank"),
				Grade: struct {
					field.RelationField
					Rank struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Department.Manager.Rank.Grade", "types.PositionGrade"),
					Rank: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Department.Manager.Rank.Grade.Rank", "types.PositionRank"),
					},
				},
			},
			Grade: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Department.Manager.Grade", "types.PositionGrade"),
			},
			LoginInformation: struct {
				field.RelationField
				Employee struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Department.Manager.LoginInformation", "types.LoginInformation"),
				Employee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Department.Manager.LoginInformation.Employee", "types.Employee"),
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
				RelationField: field.NewRelation("Department.Manager.Roles", "types.Role"),
				Employees: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Department.Manager.Roles.Employees", "types.Employee"),
				},
				Permissions: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Department.Manager.Roles.Permissions", "types.Permission"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Department.Manager.Roles.Permissions.Roles", "types.Role"),
					},
				},
				Menus: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Department.Manager.Roles.Menus", "types.Menu"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Department.Manager.Roles.Menus.Roles", "types.Role"),
					},
				},
			},
		},
	}

	_leaveSignOffSetting.Leave = leaveSignOffSettingBelongsToLeave{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Leave", "types.Leave"),
		LeaveGroup: struct {
			field.RelationField
			Leave struct {
				field.RelationField
			}
			LeaveGroupCondition struct {
				field.RelationField
				LeaveGroup struct {
					field.RelationField
				}
			}
			Employee struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Leave.LeaveGroup", "types.LeaveGroup"),
			Leave: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Leave.LeaveGroup.Leave", "types.Leave"),
			},
			LeaveGroupCondition: struct {
				field.RelationField
				LeaveGroup struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Leave.LeaveGroup.LeaveGroupCondition", "types.LeaveGroupCondition"),
				LeaveGroup: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Leave.LeaveGroup.LeaveGroupCondition.LeaveGroup", "types.LeaveGroup"),
				},
			},
			Employee: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Leave.LeaveGroup.Employee", "types.Employee"),
			},
		},
	}

	_leaveSignOffSetting.SpecificEmployee = leaveSignOffSettingBelongsToSpecificEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SpecificEmployee", "types.Employee"),
	}

	_leaveSignOffSetting.fillFieldMap()

	return _leaveSignOffSetting
}

type leaveSignOffSetting struct {
	leaveSignOffSettingDo leaveSignOffSettingDo

	ALL                field.Asterisk
	ID                 field.Uint
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field
	DepartmentID       field.Uint
	LeaveID            field.Uint
	Level              field.Uint
	GteDay             field.Float32
	SignType           field.Field
	SpecificEmployeeID field.Uint
	Notify             field.Field
	Remark             field.String
	Department         leaveSignOffSettingBelongsToDepartment

	Leave leaveSignOffSettingBelongsToLeave

	SpecificEmployee leaveSignOffSettingBelongsToSpecificEmployee

	fieldMap map[string]field.Expr
}

func (l leaveSignOffSetting) Table(newTableName string) *leaveSignOffSetting {
	l.leaveSignOffSettingDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l leaveSignOffSetting) As(alias string) *leaveSignOffSetting {
	l.leaveSignOffSettingDo.DO = *(l.leaveSignOffSettingDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *leaveSignOffSetting) updateTableName(table string) *leaveSignOffSetting {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewUint(table, "id")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.DeletedAt = field.NewField(table, "deleted_at")
	l.DepartmentID = field.NewUint(table, "department_id")
	l.LeaveID = field.NewUint(table, "leave_id")
	l.Level = field.NewUint(table, "level")
	l.GteDay = field.NewFloat32(table, "gte_day")
	l.SignType = field.NewField(table, "sign_type")
	l.SpecificEmployeeID = field.NewUint(table, "specific_employee_id")
	l.Notify = field.NewField(table, "notify")
	l.Remark = field.NewString(table, "remark")

	l.fillFieldMap()

	return l
}

func (l *leaveSignOffSetting) WithContext(ctx context.Context) ILeaveSignOffSettingDo {
	return l.leaveSignOffSettingDo.WithContext(ctx)
}

func (l leaveSignOffSetting) TableName() string { return l.leaveSignOffSettingDo.TableName() }

func (l leaveSignOffSetting) Alias() string { return l.leaveSignOffSettingDo.Alias() }

func (l leaveSignOffSetting) Columns(cols ...field.Expr) gen.Columns {
	return l.leaveSignOffSettingDo.Columns(cols...)
}

func (l *leaveSignOffSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *leaveSignOffSetting) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 15)
	l.fieldMap["id"] = l.ID
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["deleted_at"] = l.DeletedAt
	l.fieldMap["department_id"] = l.DepartmentID
	l.fieldMap["leave_id"] = l.LeaveID
	l.fieldMap["level"] = l.Level
	l.fieldMap["gte_day"] = l.GteDay
	l.fieldMap["sign_type"] = l.SignType
	l.fieldMap["specific_employee_id"] = l.SpecificEmployeeID
	l.fieldMap["notify"] = l.Notify
	l.fieldMap["remark"] = l.Remark

}

func (l leaveSignOffSetting) clone(db *gorm.DB) leaveSignOffSetting {
	l.leaveSignOffSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l leaveSignOffSetting) replaceDB(db *gorm.DB) leaveSignOffSetting {
	l.leaveSignOffSettingDo.ReplaceDB(db)
	return l
}

type leaveSignOffSettingBelongsToDepartment struct {
	db *gorm.DB

	field.RelationField

	Manager struct {
		field.RelationField
		Department struct {
			field.RelationField
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
}

func (a leaveSignOffSettingBelongsToDepartment) Where(conds ...field.Expr) *leaveSignOffSettingBelongsToDepartment {
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

func (a leaveSignOffSettingBelongsToDepartment) WithContext(ctx context.Context) *leaveSignOffSettingBelongsToDepartment {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a leaveSignOffSettingBelongsToDepartment) Session(session *gorm.Session) *leaveSignOffSettingBelongsToDepartment {
	a.db = a.db.Session(session)
	return &a
}

func (a leaveSignOffSettingBelongsToDepartment) Model(m *types.LeaveSignOffSetting) *leaveSignOffSettingBelongsToDepartmentTx {
	return &leaveSignOffSettingBelongsToDepartmentTx{a.db.Model(m).Association(a.Name())}
}

type leaveSignOffSettingBelongsToDepartmentTx struct{ tx *gorm.Association }

func (a leaveSignOffSettingBelongsToDepartmentTx) Find() (result *types.Department, err error) {
	return result, a.tx.Find(&result)
}

func (a leaveSignOffSettingBelongsToDepartmentTx) Append(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a leaveSignOffSettingBelongsToDepartmentTx) Replace(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a leaveSignOffSettingBelongsToDepartmentTx) Delete(values ...*types.Department) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a leaveSignOffSettingBelongsToDepartmentTx) Clear() error {
	return a.tx.Clear()
}

func (a leaveSignOffSettingBelongsToDepartmentTx) Count() int64 {
	return a.tx.Count()
}

type leaveSignOffSettingBelongsToLeave struct {
	db *gorm.DB

	field.RelationField

	LeaveGroup struct {
		field.RelationField
		Leave struct {
			field.RelationField
		}
		LeaveGroupCondition struct {
			field.RelationField
			LeaveGroup struct {
				field.RelationField
			}
		}
		Employee struct {
			field.RelationField
		}
	}
}

func (a leaveSignOffSettingBelongsToLeave) Where(conds ...field.Expr) *leaveSignOffSettingBelongsToLeave {
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

func (a leaveSignOffSettingBelongsToLeave) WithContext(ctx context.Context) *leaveSignOffSettingBelongsToLeave {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a leaveSignOffSettingBelongsToLeave) Session(session *gorm.Session) *leaveSignOffSettingBelongsToLeave {
	a.db = a.db.Session(session)
	return &a
}

func (a leaveSignOffSettingBelongsToLeave) Model(m *types.LeaveSignOffSetting) *leaveSignOffSettingBelongsToLeaveTx {
	return &leaveSignOffSettingBelongsToLeaveTx{a.db.Model(m).Association(a.Name())}
}

type leaveSignOffSettingBelongsToLeaveTx struct{ tx *gorm.Association }

func (a leaveSignOffSettingBelongsToLeaveTx) Find() (result *types.Leave, err error) {
	return result, a.tx.Find(&result)
}

func (a leaveSignOffSettingBelongsToLeaveTx) Append(values ...*types.Leave) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a leaveSignOffSettingBelongsToLeaveTx) Replace(values ...*types.Leave) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a leaveSignOffSettingBelongsToLeaveTx) Delete(values ...*types.Leave) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a leaveSignOffSettingBelongsToLeaveTx) Clear() error {
	return a.tx.Clear()
}

func (a leaveSignOffSettingBelongsToLeaveTx) Count() int64 {
	return a.tx.Count()
}

type leaveSignOffSettingBelongsToSpecificEmployee struct {
	db *gorm.DB

	field.RelationField
}

func (a leaveSignOffSettingBelongsToSpecificEmployee) Where(conds ...field.Expr) *leaveSignOffSettingBelongsToSpecificEmployee {
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

func (a leaveSignOffSettingBelongsToSpecificEmployee) WithContext(ctx context.Context) *leaveSignOffSettingBelongsToSpecificEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a leaveSignOffSettingBelongsToSpecificEmployee) Session(session *gorm.Session) *leaveSignOffSettingBelongsToSpecificEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a leaveSignOffSettingBelongsToSpecificEmployee) Model(m *types.LeaveSignOffSetting) *leaveSignOffSettingBelongsToSpecificEmployeeTx {
	return &leaveSignOffSettingBelongsToSpecificEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type leaveSignOffSettingBelongsToSpecificEmployeeTx struct{ tx *gorm.Association }

func (a leaveSignOffSettingBelongsToSpecificEmployeeTx) Find() (result *types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a leaveSignOffSettingBelongsToSpecificEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a leaveSignOffSettingBelongsToSpecificEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a leaveSignOffSettingBelongsToSpecificEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a leaveSignOffSettingBelongsToSpecificEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a leaveSignOffSettingBelongsToSpecificEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type leaveSignOffSettingDo struct{ gen.DO }

type ILeaveSignOffSettingDo interface {
	gen.SubQuery
	Debug() ILeaveSignOffSettingDo
	WithContext(ctx context.Context) ILeaveSignOffSettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILeaveSignOffSettingDo
	WriteDB() ILeaveSignOffSettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILeaveSignOffSettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILeaveSignOffSettingDo
	Not(conds ...gen.Condition) ILeaveSignOffSettingDo
	Or(conds ...gen.Condition) ILeaveSignOffSettingDo
	Select(conds ...field.Expr) ILeaveSignOffSettingDo
	Where(conds ...gen.Condition) ILeaveSignOffSettingDo
	Order(conds ...field.Expr) ILeaveSignOffSettingDo
	Distinct(cols ...field.Expr) ILeaveSignOffSettingDo
	Omit(cols ...field.Expr) ILeaveSignOffSettingDo
	Join(table schema.Tabler, on ...field.Expr) ILeaveSignOffSettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILeaveSignOffSettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILeaveSignOffSettingDo
	Group(cols ...field.Expr) ILeaveSignOffSettingDo
	Having(conds ...gen.Condition) ILeaveSignOffSettingDo
	Limit(limit int) ILeaveSignOffSettingDo
	Offset(offset int) ILeaveSignOffSettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILeaveSignOffSettingDo
	Unscoped() ILeaveSignOffSettingDo
	Create(values ...*types.LeaveSignOffSetting) error
	CreateInBatches(values []*types.LeaveSignOffSetting, batchSize int) error
	Save(values ...*types.LeaveSignOffSetting) error
	First() (*types.LeaveSignOffSetting, error)
	Take() (*types.LeaveSignOffSetting, error)
	Last() (*types.LeaveSignOffSetting, error)
	Find() ([]*types.LeaveSignOffSetting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.LeaveSignOffSetting, err error)
	FindInBatches(result *[]*types.LeaveSignOffSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.LeaveSignOffSetting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILeaveSignOffSettingDo
	Assign(attrs ...field.AssignExpr) ILeaveSignOffSettingDo
	Joins(fields ...field.RelationField) ILeaveSignOffSettingDo
	Preload(fields ...field.RelationField) ILeaveSignOffSettingDo
	FirstOrInit() (*types.LeaveSignOffSetting, error)
	FirstOrCreate() (*types.LeaveSignOffSetting, error)
	FindByPage(offset int, limit int) (result []*types.LeaveSignOffSetting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILeaveSignOffSettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l leaveSignOffSettingDo) Debug() ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Debug())
}

func (l leaveSignOffSettingDo) WithContext(ctx context.Context) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l leaveSignOffSettingDo) ReadDB() ILeaveSignOffSettingDo {
	return l.Clauses(dbresolver.Read)
}

func (l leaveSignOffSettingDo) WriteDB() ILeaveSignOffSettingDo {
	return l.Clauses(dbresolver.Write)
}

func (l leaveSignOffSettingDo) Session(config *gorm.Session) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Session(config))
}

func (l leaveSignOffSettingDo) Clauses(conds ...clause.Expression) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l leaveSignOffSettingDo) Returning(value interface{}, columns ...string) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l leaveSignOffSettingDo) Not(conds ...gen.Condition) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l leaveSignOffSettingDo) Or(conds ...gen.Condition) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l leaveSignOffSettingDo) Select(conds ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l leaveSignOffSettingDo) Where(conds ...gen.Condition) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l leaveSignOffSettingDo) Order(conds ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l leaveSignOffSettingDo) Distinct(cols ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l leaveSignOffSettingDo) Omit(cols ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l leaveSignOffSettingDo) Join(table schema.Tabler, on ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l leaveSignOffSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l leaveSignOffSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l leaveSignOffSettingDo) Group(cols ...field.Expr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l leaveSignOffSettingDo) Having(conds ...gen.Condition) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l leaveSignOffSettingDo) Limit(limit int) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l leaveSignOffSettingDo) Offset(offset int) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l leaveSignOffSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l leaveSignOffSettingDo) Unscoped() ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Unscoped())
}

func (l leaveSignOffSettingDo) Create(values ...*types.LeaveSignOffSetting) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l leaveSignOffSettingDo) CreateInBatches(values []*types.LeaveSignOffSetting, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l leaveSignOffSettingDo) Save(values ...*types.LeaveSignOffSetting) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l leaveSignOffSettingDo) First() (*types.LeaveSignOffSetting, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.LeaveSignOffSetting), nil
	}
}

func (l leaveSignOffSettingDo) Take() (*types.LeaveSignOffSetting, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.LeaveSignOffSetting), nil
	}
}

func (l leaveSignOffSettingDo) Last() (*types.LeaveSignOffSetting, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.LeaveSignOffSetting), nil
	}
}

func (l leaveSignOffSettingDo) Find() ([]*types.LeaveSignOffSetting, error) {
	result, err := l.DO.Find()
	return result.([]*types.LeaveSignOffSetting), err
}

func (l leaveSignOffSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.LeaveSignOffSetting, err error) {
	buf := make([]*types.LeaveSignOffSetting, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l leaveSignOffSettingDo) FindInBatches(result *[]*types.LeaveSignOffSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l leaveSignOffSettingDo) Attrs(attrs ...field.AssignExpr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l leaveSignOffSettingDo) Assign(attrs ...field.AssignExpr) ILeaveSignOffSettingDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l leaveSignOffSettingDo) Joins(fields ...field.RelationField) ILeaveSignOffSettingDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l leaveSignOffSettingDo) Preload(fields ...field.RelationField) ILeaveSignOffSettingDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l leaveSignOffSettingDo) FirstOrInit() (*types.LeaveSignOffSetting, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.LeaveSignOffSetting), nil
	}
}

func (l leaveSignOffSettingDo) FirstOrCreate() (*types.LeaveSignOffSetting, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.LeaveSignOffSetting), nil
	}
}

func (l leaveSignOffSettingDo) FindByPage(offset int, limit int) (result []*types.LeaveSignOffSetting, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l leaveSignOffSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l leaveSignOffSettingDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l leaveSignOffSettingDo) Delete(models ...*types.LeaveSignOffSetting) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *leaveSignOffSettingDo) withDO(do gen.Dao) *leaveSignOffSettingDo {
	l.DO = *do.(*gen.DO)
	return l
}
