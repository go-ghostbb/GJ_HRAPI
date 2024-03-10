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

func newVacationGroupOvertimeRate(db *gorm.DB, opts ...gen.DOOption) vacationGroupOvertimeRate {
	_vacationGroupOvertimeRate := vacationGroupOvertimeRate{}

	_vacationGroupOvertimeRate.vacationGroupOvertimeRateDo.UseDB(db, opts...)
	_vacationGroupOvertimeRate.vacationGroupOvertimeRateDo.UseModel(&types.VacationGroupOvertimeRate{})

	tableName := _vacationGroupOvertimeRate.vacationGroupOvertimeRateDo.TableName()
	_vacationGroupOvertimeRate.ALL = field.NewAsterisk(tableName)
	_vacationGroupOvertimeRate.ID = field.NewUint(tableName, "id")
	_vacationGroupOvertimeRate.CreatedAt = field.NewTime(tableName, "created_at")
	_vacationGroupOvertimeRate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_vacationGroupOvertimeRate.DeletedAt = field.NewField(tableName, "deleted_at")
	_vacationGroupOvertimeRate.VacationGroupID = field.NewUint(tableName, "vacation_group_id")
	_vacationGroupOvertimeRate.Hours = field.NewUint(tableName, "hours")
	_vacationGroupOvertimeRate.Multiply = field.NewFloat32(tableName, "multiply")
	_vacationGroupOvertimeRate.Level = field.NewUint(tableName, "level")
	_vacationGroupOvertimeRate.VacationGroup = vacationGroupOvertimeRateBelongsToVacationGroup{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("VacationGroup", "types.LeaveGroup"),
		Leave: struct {
			field.RelationField
			LeaveGroup struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("VacationGroup.Leave", "types.Leave"),
			LeaveGroup: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("VacationGroup.Leave.LeaveGroup", "types.LeaveGroup"),
			},
		},
		LeaveGroupCondition: struct {
			field.RelationField
			LeaveGroup struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("VacationGroup.LeaveGroupCondition", "types.LeaveGroupCondition"),
			LeaveGroup: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("VacationGroup.LeaveGroupCondition.LeaveGroup", "types.LeaveGroup"),
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
			RelationField: field.NewRelation("VacationGroup.Employee", "types.Employee"),
			Department: struct {
				field.RelationField
				Manager struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("VacationGroup.Employee.Department", "types.Department"),
				Manager: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("VacationGroup.Employee.Department.Manager", "types.Employee"),
				},
			},
			LoginInformation: struct {
				field.RelationField
				Employee struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("VacationGroup.Employee.LoginInformation", "types.LoginInformation"),
				Employee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("VacationGroup.Employee.LoginInformation.Employee", "types.Employee"),
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
				RelationField: field.NewRelation("VacationGroup.Employee.Roles", "types.Role"),
				Employees: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("VacationGroup.Employee.Roles.Employees", "types.Employee"),
				},
				Permissions: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("VacationGroup.Employee.Roles.Permissions", "types.Permission"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("VacationGroup.Employee.Roles.Permissions.Roles", "types.Role"),
					},
				},
				Menus: struct {
					field.RelationField
					Roles struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("VacationGroup.Employee.Roles.Menus", "types.Menu"),
					Roles: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("VacationGroup.Employee.Roles.Menus.Roles", "types.Role"),
					},
				},
			},
		},
	}

	_vacationGroupOvertimeRate.fillFieldMap()

	return _vacationGroupOvertimeRate
}

type vacationGroupOvertimeRate struct {
	vacationGroupOvertimeRateDo vacationGroupOvertimeRateDo

	ALL             field.Asterisk
	ID              field.Uint
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field
	VacationGroupID field.Uint
	Hours           field.Uint
	Multiply        field.Float32
	Level           field.Uint
	VacationGroup   vacationGroupOvertimeRateBelongsToVacationGroup

	fieldMap map[string]field.Expr
}

func (v vacationGroupOvertimeRate) Table(newTableName string) *vacationGroupOvertimeRate {
	v.vacationGroupOvertimeRateDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vacationGroupOvertimeRate) As(alias string) *vacationGroupOvertimeRate {
	v.vacationGroupOvertimeRateDo.DO = *(v.vacationGroupOvertimeRateDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vacationGroupOvertimeRate) updateTableName(table string) *vacationGroupOvertimeRate {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewUint(table, "id")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")
	v.VacationGroupID = field.NewUint(table, "vacation_group_id")
	v.Hours = field.NewUint(table, "hours")
	v.Multiply = field.NewFloat32(table, "multiply")
	v.Level = field.NewUint(table, "level")

	v.fillFieldMap()

	return v
}

func (v *vacationGroupOvertimeRate) WithContext(ctx context.Context) IVacationGroupOvertimeRateDo {
	return v.vacationGroupOvertimeRateDo.WithContext(ctx)
}

func (v vacationGroupOvertimeRate) TableName() string {
	return v.vacationGroupOvertimeRateDo.TableName()
}

func (v vacationGroupOvertimeRate) Alias() string { return v.vacationGroupOvertimeRateDo.Alias() }

func (v vacationGroupOvertimeRate) Columns(cols ...field.Expr) gen.Columns {
	return v.vacationGroupOvertimeRateDo.Columns(cols...)
}

func (v *vacationGroupOvertimeRate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vacationGroupOvertimeRate) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 9)
	v.fieldMap["id"] = v.ID
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
	v.fieldMap["vacation_group_id"] = v.VacationGroupID
	v.fieldMap["hours"] = v.Hours
	v.fieldMap["multiply"] = v.Multiply
	v.fieldMap["level"] = v.Level

}

func (v vacationGroupOvertimeRate) clone(db *gorm.DB) vacationGroupOvertimeRate {
	v.vacationGroupOvertimeRateDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vacationGroupOvertimeRate) replaceDB(db *gorm.DB) vacationGroupOvertimeRate {
	v.vacationGroupOvertimeRateDo.ReplaceDB(db)
	return v
}

type vacationGroupOvertimeRateBelongsToVacationGroup struct {
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

func (a vacationGroupOvertimeRateBelongsToVacationGroup) Where(conds ...field.Expr) *vacationGroupOvertimeRateBelongsToVacationGroup {
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

func (a vacationGroupOvertimeRateBelongsToVacationGroup) WithContext(ctx context.Context) *vacationGroupOvertimeRateBelongsToVacationGroup {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vacationGroupOvertimeRateBelongsToVacationGroup) Session(session *gorm.Session) *vacationGroupOvertimeRateBelongsToVacationGroup {
	a.db = a.db.Session(session)
	return &a
}

func (a vacationGroupOvertimeRateBelongsToVacationGroup) Model(m *types.VacationGroupOvertimeRate) *vacationGroupOvertimeRateBelongsToVacationGroupTx {
	return &vacationGroupOvertimeRateBelongsToVacationGroupTx{a.db.Model(m).Association(a.Name())}
}

type vacationGroupOvertimeRateBelongsToVacationGroupTx struct{ tx *gorm.Association }

func (a vacationGroupOvertimeRateBelongsToVacationGroupTx) Find() (result *types.LeaveGroup, err error) {
	return result, a.tx.Find(&result)
}

func (a vacationGroupOvertimeRateBelongsToVacationGroupTx) Append(values ...*types.LeaveGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vacationGroupOvertimeRateBelongsToVacationGroupTx) Replace(values ...*types.LeaveGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vacationGroupOvertimeRateBelongsToVacationGroupTx) Delete(values ...*types.LeaveGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vacationGroupOvertimeRateBelongsToVacationGroupTx) Clear() error {
	return a.tx.Clear()
}

func (a vacationGroupOvertimeRateBelongsToVacationGroupTx) Count() int64 {
	return a.tx.Count()
}

type vacationGroupOvertimeRateDo struct{ gen.DO }

type IVacationGroupOvertimeRateDo interface {
	gen.SubQuery
	Debug() IVacationGroupOvertimeRateDo
	WithContext(ctx context.Context) IVacationGroupOvertimeRateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVacationGroupOvertimeRateDo
	WriteDB() IVacationGroupOvertimeRateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVacationGroupOvertimeRateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVacationGroupOvertimeRateDo
	Not(conds ...gen.Condition) IVacationGroupOvertimeRateDo
	Or(conds ...gen.Condition) IVacationGroupOvertimeRateDo
	Select(conds ...field.Expr) IVacationGroupOvertimeRateDo
	Where(conds ...gen.Condition) IVacationGroupOvertimeRateDo
	Order(conds ...field.Expr) IVacationGroupOvertimeRateDo
	Distinct(cols ...field.Expr) IVacationGroupOvertimeRateDo
	Omit(cols ...field.Expr) IVacationGroupOvertimeRateDo
	Join(table schema.Tabler, on ...field.Expr) IVacationGroupOvertimeRateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVacationGroupOvertimeRateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVacationGroupOvertimeRateDo
	Group(cols ...field.Expr) IVacationGroupOvertimeRateDo
	Having(conds ...gen.Condition) IVacationGroupOvertimeRateDo
	Limit(limit int) IVacationGroupOvertimeRateDo
	Offset(offset int) IVacationGroupOvertimeRateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVacationGroupOvertimeRateDo
	Unscoped() IVacationGroupOvertimeRateDo
	Create(values ...*types.VacationGroupOvertimeRate) error
	CreateInBatches(values []*types.VacationGroupOvertimeRate, batchSize int) error
	Save(values ...*types.VacationGroupOvertimeRate) error
	First() (*types.VacationGroupOvertimeRate, error)
	Take() (*types.VacationGroupOvertimeRate, error)
	Last() (*types.VacationGroupOvertimeRate, error)
	Find() ([]*types.VacationGroupOvertimeRate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.VacationGroupOvertimeRate, err error)
	FindInBatches(result *[]*types.VacationGroupOvertimeRate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.VacationGroupOvertimeRate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVacationGroupOvertimeRateDo
	Assign(attrs ...field.AssignExpr) IVacationGroupOvertimeRateDo
	Joins(fields ...field.RelationField) IVacationGroupOvertimeRateDo
	Preload(fields ...field.RelationField) IVacationGroupOvertimeRateDo
	FirstOrInit() (*types.VacationGroupOvertimeRate, error)
	FirstOrCreate() (*types.VacationGroupOvertimeRate, error)
	FindByPage(offset int, limit int) (result []*types.VacationGroupOvertimeRate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVacationGroupOvertimeRateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v vacationGroupOvertimeRateDo) Debug() IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Debug())
}

func (v vacationGroupOvertimeRateDo) WithContext(ctx context.Context) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vacationGroupOvertimeRateDo) ReadDB() IVacationGroupOvertimeRateDo {
	return v.Clauses(dbresolver.Read)
}

func (v vacationGroupOvertimeRateDo) WriteDB() IVacationGroupOvertimeRateDo {
	return v.Clauses(dbresolver.Write)
}

func (v vacationGroupOvertimeRateDo) Session(config *gorm.Session) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Session(config))
}

func (v vacationGroupOvertimeRateDo) Clauses(conds ...clause.Expression) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vacationGroupOvertimeRateDo) Returning(value interface{}, columns ...string) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vacationGroupOvertimeRateDo) Not(conds ...gen.Condition) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vacationGroupOvertimeRateDo) Or(conds ...gen.Condition) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vacationGroupOvertimeRateDo) Select(conds ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vacationGroupOvertimeRateDo) Where(conds ...gen.Condition) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vacationGroupOvertimeRateDo) Order(conds ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vacationGroupOvertimeRateDo) Distinct(cols ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vacationGroupOvertimeRateDo) Omit(cols ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vacationGroupOvertimeRateDo) Join(table schema.Tabler, on ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vacationGroupOvertimeRateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vacationGroupOvertimeRateDo) RightJoin(table schema.Tabler, on ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vacationGroupOvertimeRateDo) Group(cols ...field.Expr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vacationGroupOvertimeRateDo) Having(conds ...gen.Condition) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vacationGroupOvertimeRateDo) Limit(limit int) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vacationGroupOvertimeRateDo) Offset(offset int) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vacationGroupOvertimeRateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vacationGroupOvertimeRateDo) Unscoped() IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vacationGroupOvertimeRateDo) Create(values ...*types.VacationGroupOvertimeRate) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vacationGroupOvertimeRateDo) CreateInBatches(values []*types.VacationGroupOvertimeRate, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vacationGroupOvertimeRateDo) Save(values ...*types.VacationGroupOvertimeRate) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vacationGroupOvertimeRateDo) First() (*types.VacationGroupOvertimeRate, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroupOvertimeRate), nil
	}
}

func (v vacationGroupOvertimeRateDo) Take() (*types.VacationGroupOvertimeRate, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroupOvertimeRate), nil
	}
}

func (v vacationGroupOvertimeRateDo) Last() (*types.VacationGroupOvertimeRate, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroupOvertimeRate), nil
	}
}

func (v vacationGroupOvertimeRateDo) Find() ([]*types.VacationGroupOvertimeRate, error) {
	result, err := v.DO.Find()
	return result.([]*types.VacationGroupOvertimeRate), err
}

func (v vacationGroupOvertimeRateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.VacationGroupOvertimeRate, err error) {
	buf := make([]*types.VacationGroupOvertimeRate, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vacationGroupOvertimeRateDo) FindInBatches(result *[]*types.VacationGroupOvertimeRate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vacationGroupOvertimeRateDo) Attrs(attrs ...field.AssignExpr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vacationGroupOvertimeRateDo) Assign(attrs ...field.AssignExpr) IVacationGroupOvertimeRateDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vacationGroupOvertimeRateDo) Joins(fields ...field.RelationField) IVacationGroupOvertimeRateDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vacationGroupOvertimeRateDo) Preload(fields ...field.RelationField) IVacationGroupOvertimeRateDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vacationGroupOvertimeRateDo) FirstOrInit() (*types.VacationGroupOvertimeRate, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroupOvertimeRate), nil
	}
}

func (v vacationGroupOvertimeRateDo) FirstOrCreate() (*types.VacationGroupOvertimeRate, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroupOvertimeRate), nil
	}
}

func (v vacationGroupOvertimeRateDo) FindByPage(offset int, limit int) (result []*types.VacationGroupOvertimeRate, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vacationGroupOvertimeRateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vacationGroupOvertimeRateDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vacationGroupOvertimeRateDo) Delete(models ...*types.VacationGroupOvertimeRate) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vacationGroupOvertimeRateDo) withDO(do gen.Dao) *vacationGroupOvertimeRateDo {
	v.DO = *do.(*gen.DO)
	return v
}
