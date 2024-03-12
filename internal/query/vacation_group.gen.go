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

func newVacationGroup(db *gorm.DB, opts ...gen.DOOption) vacationGroup {
	_vacationGroup := vacationGroup{}

	_vacationGroup.vacationGroupDo.UseDB(db, opts...)
	_vacationGroup.vacationGroupDo.UseModel(&types.VacationGroup{})

	tableName := _vacationGroup.vacationGroupDo.TableName()
	_vacationGroup.ALL = field.NewAsterisk(tableName)
	_vacationGroup.ID = field.NewUint(tableName, "id")
	_vacationGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_vacationGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_vacationGroup.DeletedAt = field.NewField(tableName, "deleted_at")
	_vacationGroup.VacationID = field.NewUint(tableName, "vacation_id")
	_vacationGroup.Name = field.NewString(tableName, "name")
	_vacationGroup.VacationGroupOvertimeRate = vacationGroupHasManyVacationGroupOvertimeRate{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("VacationGroupOvertimeRate", "types.VacationGroupOvertimeRate"),
		VacationGroup: struct {
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
		}{
			RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup", "types.LeaveGroup"),
			Leave: struct {
				field.RelationField
				LeaveGroup struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Leave", "types.Leave"),
				LeaveGroup: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Leave.LeaveGroup", "types.LeaveGroup"),
				},
			},
			LeaveGroupCondition: struct {
				field.RelationField
				LeaveGroup struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.LeaveGroupCondition", "types.LeaveGroupCondition"),
				LeaveGroup: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.LeaveGroupCondition.LeaveGroup", "types.LeaveGroup"),
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
				RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee", "types.Employee"),
				Department: struct {
					field.RelationField
					Manager struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Department", "types.Department"),
					Manager: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Department.Manager", "types.Employee"),
					},
				},
				LoginInformation: struct {
					field.RelationField
					Employee struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.LoginInformation", "types.LoginInformation"),
					Employee: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.LoginInformation.Employee", "types.Employee"),
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
					RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Roles", "types.Role"),
					Employees: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Employees", "types.Employee"),
					},
					Permissions: struct {
						field.RelationField
						Roles struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Permissions", "types.Permission"),
						Roles: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Permissions.Roles", "types.Role"),
						},
					},
					Menus: struct {
						field.RelationField
						Roles struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Menus", "types.Menu"),
						Roles: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Menus.Roles", "types.Role"),
						},
					},
				},
			},
		},
	}

	_vacationGroup.Vacation = vacationGroupBelongsToVacation{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Vacation", "types.Vacation"),
		Schedule: struct {
			field.RelationField
			Vacation struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Vacation.Schedule", "types.VacationSchedule"),
			Vacation: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Vacation.Schedule.Vacation", "types.Vacation"),
			},
		},
		VacationGroup: struct {
			field.RelationField
			Vacation struct {
				field.RelationField
			}
			VacationGroupOvertimeRate struct {
				field.RelationField
			}
			Employee struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Vacation.VacationGroup", "types.VacationGroup"),
			Vacation: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Vacation.VacationGroup.Vacation", "types.Vacation"),
			},
			VacationGroupOvertimeRate: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Vacation.VacationGroup.VacationGroupOvertimeRate", "types.VacationGroupOvertimeRate"),
			},
			Employee: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Vacation.VacationGroup.Employee", "types.Employee"),
			},
		},
	}

	_vacationGroup.Employee = vacationGroupManyToManyEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Employee", "types.Employee"),
	}

	_vacationGroup.fillFieldMap()

	return _vacationGroup
}

type vacationGroup struct {
	vacationGroupDo vacationGroupDo

	ALL                       field.Asterisk
	ID                        field.Uint
	CreatedAt                 field.Time
	UpdatedAt                 field.Time
	DeletedAt                 field.Field
	VacationID                field.Uint
	Name                      field.String
	VacationGroupOvertimeRate vacationGroupHasManyVacationGroupOvertimeRate

	Vacation vacationGroupBelongsToVacation

	Employee vacationGroupManyToManyEmployee

	fieldMap map[string]field.Expr
}

func (v vacationGroup) Table(newTableName string) *vacationGroup {
	v.vacationGroupDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vacationGroup) As(alias string) *vacationGroup {
	v.vacationGroupDo.DO = *(v.vacationGroupDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vacationGroup) updateTableName(table string) *vacationGroup {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewUint(table, "id")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")
	v.VacationID = field.NewUint(table, "vacation_id")
	v.Name = field.NewString(table, "name")

	v.fillFieldMap()

	return v
}

func (v *vacationGroup) WithContext(ctx context.Context) IVacationGroupDo {
	return v.vacationGroupDo.WithContext(ctx)
}

func (v vacationGroup) TableName() string { return v.vacationGroupDo.TableName() }

func (v vacationGroup) Alias() string { return v.vacationGroupDo.Alias() }

func (v vacationGroup) Columns(cols ...field.Expr) gen.Columns {
	return v.vacationGroupDo.Columns(cols...)
}

func (v *vacationGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vacationGroup) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 9)
	v.fieldMap["id"] = v.ID
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
	v.fieldMap["vacation_id"] = v.VacationID
	v.fieldMap["name"] = v.Name

}

func (v vacationGroup) clone(db *gorm.DB) vacationGroup {
	v.vacationGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vacationGroup) replaceDB(db *gorm.DB) vacationGroup {
	v.vacationGroupDo.ReplaceDB(db)
	return v
}

type vacationGroupHasManyVacationGroupOvertimeRate struct {
	db *gorm.DB

	field.RelationField

	VacationGroup struct {
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
}

func (a vacationGroupHasManyVacationGroupOvertimeRate) Where(conds ...field.Expr) *vacationGroupHasManyVacationGroupOvertimeRate {
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

func (a vacationGroupHasManyVacationGroupOvertimeRate) WithContext(ctx context.Context) *vacationGroupHasManyVacationGroupOvertimeRate {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vacationGroupHasManyVacationGroupOvertimeRate) Session(session *gorm.Session) *vacationGroupHasManyVacationGroupOvertimeRate {
	a.db = a.db.Session(session)
	return &a
}

func (a vacationGroupHasManyVacationGroupOvertimeRate) Model(m *types.VacationGroup) *vacationGroupHasManyVacationGroupOvertimeRateTx {
	return &vacationGroupHasManyVacationGroupOvertimeRateTx{a.db.Model(m).Association(a.Name())}
}

type vacationGroupHasManyVacationGroupOvertimeRateTx struct{ tx *gorm.Association }

func (a vacationGroupHasManyVacationGroupOvertimeRateTx) Find() (result []*types.VacationGroupOvertimeRate, err error) {
	return result, a.tx.Find(&result)
}

func (a vacationGroupHasManyVacationGroupOvertimeRateTx) Append(values ...*types.VacationGroupOvertimeRate) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vacationGroupHasManyVacationGroupOvertimeRateTx) Replace(values ...*types.VacationGroupOvertimeRate) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vacationGroupHasManyVacationGroupOvertimeRateTx) Delete(values ...*types.VacationGroupOvertimeRate) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vacationGroupHasManyVacationGroupOvertimeRateTx) Clear() error {
	return a.tx.Clear()
}

func (a vacationGroupHasManyVacationGroupOvertimeRateTx) Count() int64 {
	return a.tx.Count()
}

type vacationGroupBelongsToVacation struct {
	db *gorm.DB

	field.RelationField

	Schedule struct {
		field.RelationField
		Vacation struct {
			field.RelationField
		}
	}
	VacationGroup struct {
		field.RelationField
		Vacation struct {
			field.RelationField
		}
		VacationGroupOvertimeRate struct {
			field.RelationField
		}
		Employee struct {
			field.RelationField
		}
	}
}

func (a vacationGroupBelongsToVacation) Where(conds ...field.Expr) *vacationGroupBelongsToVacation {
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

func (a vacationGroupBelongsToVacation) WithContext(ctx context.Context) *vacationGroupBelongsToVacation {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vacationGroupBelongsToVacation) Session(session *gorm.Session) *vacationGroupBelongsToVacation {
	a.db = a.db.Session(session)
	return &a
}

func (a vacationGroupBelongsToVacation) Model(m *types.VacationGroup) *vacationGroupBelongsToVacationTx {
	return &vacationGroupBelongsToVacationTx{a.db.Model(m).Association(a.Name())}
}

type vacationGroupBelongsToVacationTx struct{ tx *gorm.Association }

func (a vacationGroupBelongsToVacationTx) Find() (result *types.Vacation, err error) {
	return result, a.tx.Find(&result)
}

func (a vacationGroupBelongsToVacationTx) Append(values ...*types.Vacation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vacationGroupBelongsToVacationTx) Replace(values ...*types.Vacation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vacationGroupBelongsToVacationTx) Delete(values ...*types.Vacation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vacationGroupBelongsToVacationTx) Clear() error {
	return a.tx.Clear()
}

func (a vacationGroupBelongsToVacationTx) Count() int64 {
	return a.tx.Count()
}

type vacationGroupManyToManyEmployee struct {
	db *gorm.DB

	field.RelationField
}

func (a vacationGroupManyToManyEmployee) Where(conds ...field.Expr) *vacationGroupManyToManyEmployee {
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

func (a vacationGroupManyToManyEmployee) WithContext(ctx context.Context) *vacationGroupManyToManyEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vacationGroupManyToManyEmployee) Session(session *gorm.Session) *vacationGroupManyToManyEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a vacationGroupManyToManyEmployee) Model(m *types.VacationGroup) *vacationGroupManyToManyEmployeeTx {
	return &vacationGroupManyToManyEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type vacationGroupManyToManyEmployeeTx struct{ tx *gorm.Association }

func (a vacationGroupManyToManyEmployeeTx) Find() (result []*types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a vacationGroupManyToManyEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vacationGroupManyToManyEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vacationGroupManyToManyEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vacationGroupManyToManyEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a vacationGroupManyToManyEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type vacationGroupDo struct{ gen.DO }

type IVacationGroupDo interface {
	gen.SubQuery
	Debug() IVacationGroupDo
	WithContext(ctx context.Context) IVacationGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVacationGroupDo
	WriteDB() IVacationGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVacationGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVacationGroupDo
	Not(conds ...gen.Condition) IVacationGroupDo
	Or(conds ...gen.Condition) IVacationGroupDo
	Select(conds ...field.Expr) IVacationGroupDo
	Where(conds ...gen.Condition) IVacationGroupDo
	Order(conds ...field.Expr) IVacationGroupDo
	Distinct(cols ...field.Expr) IVacationGroupDo
	Omit(cols ...field.Expr) IVacationGroupDo
	Join(table schema.Tabler, on ...field.Expr) IVacationGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVacationGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVacationGroupDo
	Group(cols ...field.Expr) IVacationGroupDo
	Having(conds ...gen.Condition) IVacationGroupDo
	Limit(limit int) IVacationGroupDo
	Offset(offset int) IVacationGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVacationGroupDo
	Unscoped() IVacationGroupDo
	Create(values ...*types.VacationGroup) error
	CreateInBatches(values []*types.VacationGroup, batchSize int) error
	Save(values ...*types.VacationGroup) error
	First() (*types.VacationGroup, error)
	Take() (*types.VacationGroup, error)
	Last() (*types.VacationGroup, error)
	Find() ([]*types.VacationGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.VacationGroup, err error)
	FindInBatches(result *[]*types.VacationGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.VacationGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVacationGroupDo
	Assign(attrs ...field.AssignExpr) IVacationGroupDo
	Joins(fields ...field.RelationField) IVacationGroupDo
	Preload(fields ...field.RelationField) IVacationGroupDo
	FirstOrInit() (*types.VacationGroup, error)
	FirstOrCreate() (*types.VacationGroup, error)
	FindByPage(offset int, limit int) (result []*types.VacationGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVacationGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v vacationGroupDo) Debug() IVacationGroupDo {
	return v.withDO(v.DO.Debug())
}

func (v vacationGroupDo) WithContext(ctx context.Context) IVacationGroupDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vacationGroupDo) ReadDB() IVacationGroupDo {
	return v.Clauses(dbresolver.Read)
}

func (v vacationGroupDo) WriteDB() IVacationGroupDo {
	return v.Clauses(dbresolver.Write)
}

func (v vacationGroupDo) Session(config *gorm.Session) IVacationGroupDo {
	return v.withDO(v.DO.Session(config))
}

func (v vacationGroupDo) Clauses(conds ...clause.Expression) IVacationGroupDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vacationGroupDo) Returning(value interface{}, columns ...string) IVacationGroupDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vacationGroupDo) Not(conds ...gen.Condition) IVacationGroupDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vacationGroupDo) Or(conds ...gen.Condition) IVacationGroupDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vacationGroupDo) Select(conds ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vacationGroupDo) Where(conds ...gen.Condition) IVacationGroupDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vacationGroupDo) Order(conds ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vacationGroupDo) Distinct(cols ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vacationGroupDo) Omit(cols ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vacationGroupDo) Join(table schema.Tabler, on ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vacationGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vacationGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vacationGroupDo) Group(cols ...field.Expr) IVacationGroupDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vacationGroupDo) Having(conds ...gen.Condition) IVacationGroupDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vacationGroupDo) Limit(limit int) IVacationGroupDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vacationGroupDo) Offset(offset int) IVacationGroupDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vacationGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVacationGroupDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vacationGroupDo) Unscoped() IVacationGroupDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vacationGroupDo) Create(values ...*types.VacationGroup) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vacationGroupDo) CreateInBatches(values []*types.VacationGroup, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vacationGroupDo) Save(values ...*types.VacationGroup) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vacationGroupDo) First() (*types.VacationGroup, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroup), nil
	}
}

func (v vacationGroupDo) Take() (*types.VacationGroup, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroup), nil
	}
}

func (v vacationGroupDo) Last() (*types.VacationGroup, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroup), nil
	}
}

func (v vacationGroupDo) Find() ([]*types.VacationGroup, error) {
	result, err := v.DO.Find()
	return result.([]*types.VacationGroup), err
}

func (v vacationGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.VacationGroup, err error) {
	buf := make([]*types.VacationGroup, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vacationGroupDo) FindInBatches(result *[]*types.VacationGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vacationGroupDo) Attrs(attrs ...field.AssignExpr) IVacationGroupDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vacationGroupDo) Assign(attrs ...field.AssignExpr) IVacationGroupDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vacationGroupDo) Joins(fields ...field.RelationField) IVacationGroupDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vacationGroupDo) Preload(fields ...field.RelationField) IVacationGroupDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vacationGroupDo) FirstOrInit() (*types.VacationGroup, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroup), nil
	}
}

func (v vacationGroupDo) FirstOrCreate() (*types.VacationGroup, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.VacationGroup), nil
	}
}

func (v vacationGroupDo) FindByPage(offset int, limit int) (result []*types.VacationGroup, count int64, err error) {
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

func (v vacationGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vacationGroupDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vacationGroupDo) Delete(models ...*types.VacationGroup) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vacationGroupDo) withDO(do gen.Dao) *vacationGroupDo {
	v.DO = *do.(*gen.DO)
	return v
}
