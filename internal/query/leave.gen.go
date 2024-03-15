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

func newLeave(db *gorm.DB, opts ...gen.DOOption) leave {
	_leave := leave{}

	_leave.leaveDo.UseDB(db, opts...)
	_leave.leaveDo.UseModel(&types.Leave{})

	tableName := _leave.leaveDo.TableName()
	_leave.ALL = field.NewAsterisk(tableName)
	_leave.ID = field.NewUint(tableName, "id")
	_leave.CreatedAt = field.NewTime(tableName, "created_at")
	_leave.UpdatedAt = field.NewTime(tableName, "updated_at")
	_leave.DeletedAt = field.NewField(tableName, "deleted_at")
	_leave.Code = field.NewString(tableName, "code")
	_leave.Name = field.NewString(tableName, "name")
	_leave.Status = field.NewBool(tableName, "status")
	_leave.Default = field.NewUint(tableName, "default")
	_leave.Minimum = field.NewUint(tableName, "minimum")
	_leave.Pay = field.NewField(tableName, "pay")
	_leave.Cycle = field.NewField(tableName, "cycle")
	_leave.Remark = field.NewString(tableName, "remark")
	_leave.Deferrable = field.NewBool(tableName, "deferrable")
	_leave.DeferrableDays = field.NewUint(tableName, "deferrable_days")
	_leave.CustomizableDuration = field.NewBool(tableName, "customizable_duration")
	_leave.Duration = field.NewUint(tableName, "duration")
	_leave.LeaveGroup = leaveHasManyLeaveGroup{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("LeaveGroup", "types.LeaveGroup"),
		Leave: struct {
			field.RelationField
			LeaveGroup struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("LeaveGroup.Leave", "types.Leave"),
			LeaveGroup: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("LeaveGroup.Leave.LeaveGroup", "types.LeaveGroup"),
			},
		},
		LeaveGroupCondition: struct {
			field.RelationField
			LeaveGroup struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("LeaveGroup.LeaveGroupCondition", "types.LeaveGroupCondition"),
			LeaveGroup: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("LeaveGroup.LeaveGroupCondition.LeaveGroup", "types.LeaveGroup"),
			},
		},
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
			RelationField: field.NewRelation("LeaveGroup.Employee", "types.Employee"),
			Department: struct {
				field.RelationField
				Manager struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("LeaveGroup.Employee.Department", "types.Department"),
				Manager: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("LeaveGroup.Employee.Department.Manager", "types.Employee"),
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
				RelationField: field.NewRelation("LeaveGroup.Employee.Rank", "types.PositionRank"),
				Grade: struct {
					field.RelationField
					Rank struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("LeaveGroup.Employee.Rank.Grade", "types.PositionGrade"),
					Rank: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("LeaveGroup.Employee.Rank.Grade.Rank", "types.PositionRank"),
					},
				},
			},
			Grade: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("LeaveGroup.Employee.Grade", "types.PositionGrade"),
			},
			LoginInformation: struct {
				field.RelationField
				Employee struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("LeaveGroup.Employee.LoginInformation", "types.LoginInformation"),
				Employee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("LeaveGroup.Employee.LoginInformation.Employee", "types.Employee"),
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
				RelationField: field.NewRelation("LeaveGroup.Employee.Roles", "types.Role"),
				Employees: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("LeaveGroup.Employee.Roles.Employees", "types.Employee"),
				},
				Permissions: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("LeaveGroup.Employee.Roles.Permissions", "types.Permission"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("LeaveGroup.Employee.Roles.Permissions.Roles", "types.Role"),
					},
				},
				Menus: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("LeaveGroup.Employee.Roles.Menus", "types.Menu"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("LeaveGroup.Employee.Roles.Menus.Roles", "types.Role"),
					},
				},
			},
		},
	}

	_leave.fillFieldMap()

	return _leave
}

type leave struct {
	leaveDo leaveDo

	ALL                  field.Asterisk
	ID                   field.Uint
	CreatedAt            field.Time
	UpdatedAt            field.Time
	DeletedAt            field.Field
	Code                 field.String
	Name                 field.String
	Status               field.Bool
	Default              field.Uint
	Minimum              field.Uint
	Pay                  field.Field
	Cycle                field.Field
	Remark               field.String
	Deferrable           field.Bool
	DeferrableDays       field.Uint
	CustomizableDuration field.Bool
	Duration             field.Uint
	LeaveGroup           leaveHasManyLeaveGroup

	fieldMap map[string]field.Expr
}

func (l leave) Table(newTableName string) *leave {
	l.leaveDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l leave) As(alias string) *leave {
	l.leaveDo.DO = *(l.leaveDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *leave) updateTableName(table string) *leave {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewUint(table, "id")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.DeletedAt = field.NewField(table, "deleted_at")
	l.Code = field.NewString(table, "code")
	l.Name = field.NewString(table, "name")
	l.Status = field.NewBool(table, "status")
	l.Default = field.NewUint(table, "default")
	l.Minimum = field.NewUint(table, "minimum")
	l.Pay = field.NewField(table, "pay")
	l.Cycle = field.NewField(table, "cycle")
	l.Remark = field.NewString(table, "remark")
	l.Deferrable = field.NewBool(table, "deferrable")
	l.DeferrableDays = field.NewUint(table, "deferrable_days")
	l.CustomizableDuration = field.NewBool(table, "customizable_duration")
	l.Duration = field.NewUint(table, "duration")

	l.fillFieldMap()

	return l
}

func (l *leave) WithContext(ctx context.Context) ILeaveDo { return l.leaveDo.WithContext(ctx) }

func (l leave) TableName() string { return l.leaveDo.TableName() }

func (l leave) Alias() string { return l.leaveDo.Alias() }

func (l leave) Columns(cols ...field.Expr) gen.Columns { return l.leaveDo.Columns(cols...) }

func (l *leave) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *leave) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 17)
	l.fieldMap["id"] = l.ID
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["deleted_at"] = l.DeletedAt
	l.fieldMap["code"] = l.Code
	l.fieldMap["name"] = l.Name
	l.fieldMap["status"] = l.Status
	l.fieldMap["default"] = l.Default
	l.fieldMap["minimum"] = l.Minimum
	l.fieldMap["pay"] = l.Pay
	l.fieldMap["cycle"] = l.Cycle
	l.fieldMap["remark"] = l.Remark
	l.fieldMap["deferrable"] = l.Deferrable
	l.fieldMap["deferrable_days"] = l.DeferrableDays
	l.fieldMap["customizable_duration"] = l.CustomizableDuration
	l.fieldMap["duration"] = l.Duration

}

func (l leave) clone(db *gorm.DB) leave {
	l.leaveDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l leave) replaceDB(db *gorm.DB) leave {
	l.leaveDo.ReplaceDB(db)
	return l
}

type leaveHasManyLeaveGroup struct {
	db *gorm.DB

	field.RelationField

	Leave struct {
		field.RelationField
		LeaveGroup struct {
			field.RelationField
		}
	}
	LeaveGroupCondition struct {
		field.RelationField
		LeaveGroup struct {
			field.RelationField
		}
	}
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
}

func (a leaveHasManyLeaveGroup) Where(conds ...field.Expr) *leaveHasManyLeaveGroup {
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

func (a leaveHasManyLeaveGroup) WithContext(ctx context.Context) *leaveHasManyLeaveGroup {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a leaveHasManyLeaveGroup) Session(session *gorm.Session) *leaveHasManyLeaveGroup {
	a.db = a.db.Session(session)
	return &a
}

func (a leaveHasManyLeaveGroup) Model(m *types.Leave) *leaveHasManyLeaveGroupTx {
	return &leaveHasManyLeaveGroupTx{a.db.Model(m).Association(a.Name())}
}

type leaveHasManyLeaveGroupTx struct{ tx *gorm.Association }

func (a leaveHasManyLeaveGroupTx) Find() (result []*types.LeaveGroup, err error) {
	return result, a.tx.Find(&result)
}

func (a leaveHasManyLeaveGroupTx) Append(values ...*types.LeaveGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a leaveHasManyLeaveGroupTx) Replace(values ...*types.LeaveGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a leaveHasManyLeaveGroupTx) Delete(values ...*types.LeaveGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a leaveHasManyLeaveGroupTx) Clear() error {
	return a.tx.Clear()
}

func (a leaveHasManyLeaveGroupTx) Count() int64 {
	return a.tx.Count()
}

type leaveDo struct{ gen.DO }

type ILeaveDo interface {
	gen.SubQuery
	Debug() ILeaveDo
	WithContext(ctx context.Context) ILeaveDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILeaveDo
	WriteDB() ILeaveDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILeaveDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILeaveDo
	Not(conds ...gen.Condition) ILeaveDo
	Or(conds ...gen.Condition) ILeaveDo
	Select(conds ...field.Expr) ILeaveDo
	Where(conds ...gen.Condition) ILeaveDo
	Order(conds ...field.Expr) ILeaveDo
	Distinct(cols ...field.Expr) ILeaveDo
	Omit(cols ...field.Expr) ILeaveDo
	Join(table schema.Tabler, on ...field.Expr) ILeaveDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILeaveDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILeaveDo
	Group(cols ...field.Expr) ILeaveDo
	Having(conds ...gen.Condition) ILeaveDo
	Limit(limit int) ILeaveDo
	Offset(offset int) ILeaveDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILeaveDo
	Unscoped() ILeaveDo
	Create(values ...*types.Leave) error
	CreateInBatches(values []*types.Leave, batchSize int) error
	Save(values ...*types.Leave) error
	First() (*types.Leave, error)
	Take() (*types.Leave, error)
	Last() (*types.Leave, error)
	Find() ([]*types.Leave, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.Leave, err error)
	FindInBatches(result *[]*types.Leave, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.Leave) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILeaveDo
	Assign(attrs ...field.AssignExpr) ILeaveDo
	Joins(fields ...field.RelationField) ILeaveDo
	Preload(fields ...field.RelationField) ILeaveDo
	FirstOrInit() (*types.Leave, error)
	FirstOrCreate() (*types.Leave, error)
	FindByPage(offset int, limit int) (result []*types.Leave, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILeaveDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l leaveDo) Debug() ILeaveDo {
	return l.withDO(l.DO.Debug())
}

func (l leaveDo) WithContext(ctx context.Context) ILeaveDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l leaveDo) ReadDB() ILeaveDo {
	return l.Clauses(dbresolver.Read)
}

func (l leaveDo) WriteDB() ILeaveDo {
	return l.Clauses(dbresolver.Write)
}

func (l leaveDo) Session(config *gorm.Session) ILeaveDo {
	return l.withDO(l.DO.Session(config))
}

func (l leaveDo) Clauses(conds ...clause.Expression) ILeaveDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l leaveDo) Returning(value interface{}, columns ...string) ILeaveDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l leaveDo) Not(conds ...gen.Condition) ILeaveDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l leaveDo) Or(conds ...gen.Condition) ILeaveDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l leaveDo) Select(conds ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l leaveDo) Where(conds ...gen.Condition) ILeaveDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l leaveDo) Order(conds ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l leaveDo) Distinct(cols ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l leaveDo) Omit(cols ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l leaveDo) Join(table schema.Tabler, on ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l leaveDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l leaveDo) RightJoin(table schema.Tabler, on ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l leaveDo) Group(cols ...field.Expr) ILeaveDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l leaveDo) Having(conds ...gen.Condition) ILeaveDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l leaveDo) Limit(limit int) ILeaveDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l leaveDo) Offset(offset int) ILeaveDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l leaveDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILeaveDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l leaveDo) Unscoped() ILeaveDo {
	return l.withDO(l.DO.Unscoped())
}

func (l leaveDo) Create(values ...*types.Leave) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l leaveDo) CreateInBatches(values []*types.Leave, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l leaveDo) Save(values ...*types.Leave) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l leaveDo) First() (*types.Leave, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.Leave), nil
	}
}

func (l leaveDo) Take() (*types.Leave, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.Leave), nil
	}
}

func (l leaveDo) Last() (*types.Leave, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.Leave), nil
	}
}

func (l leaveDo) Find() ([]*types.Leave, error) {
	result, err := l.DO.Find()
	return result.([]*types.Leave), err
}

func (l leaveDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.Leave, err error) {
	buf := make([]*types.Leave, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l leaveDo) FindInBatches(result *[]*types.Leave, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l leaveDo) Attrs(attrs ...field.AssignExpr) ILeaveDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l leaveDo) Assign(attrs ...field.AssignExpr) ILeaveDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l leaveDo) Joins(fields ...field.RelationField) ILeaveDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l leaveDo) Preload(fields ...field.RelationField) ILeaveDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l leaveDo) FirstOrInit() (*types.Leave, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.Leave), nil
	}
}

func (l leaveDo) FirstOrCreate() (*types.Leave, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.Leave), nil
	}
}

func (l leaveDo) FindByPage(offset int, limit int) (result []*types.Leave, count int64, err error) {
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

func (l leaveDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l leaveDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l leaveDo) Delete(models ...*types.Leave) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *leaveDo) withDO(do gen.Dao) *leaveDo {
	l.DO = *do.(*gen.DO)
	return l
}
