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

func newCalcSalaryEmployee(db *gorm.DB, opts ...gen.DOOption) calcSalaryEmployee {
	_calcSalaryEmployee := calcSalaryEmployee{}

	_calcSalaryEmployee.calcSalaryEmployeeDo.UseDB(db, opts...)
	_calcSalaryEmployee.calcSalaryEmployeeDo.UseModel(&types.CalcSalaryEmployee{})

	tableName := _calcSalaryEmployee.calcSalaryEmployeeDo.TableName()
	_calcSalaryEmployee.ALL = field.NewAsterisk(tableName)
	_calcSalaryEmployee.ID = field.NewUint(tableName, "id")
	_calcSalaryEmployee.CreatedAt = field.NewTime(tableName, "created_at")
	_calcSalaryEmployee.UpdatedAt = field.NewTime(tableName, "updated_at")
	_calcSalaryEmployee.DeletedAt = field.NewField(tableName, "deleted_at")
	_calcSalaryEmployee.EmployeeID = field.NewUint(tableName, "employee_id")
	_calcSalaryEmployee.CalcSalaryID = field.NewUint(tableName, "calc_salary_id")
	_calcSalaryEmployee.Salary = field.NewFloat32(tableName, "salary")
	_calcSalaryEmployee.HourlySalary = field.NewFloat32(tableName, "hourly_salary")
	_calcSalaryEmployee.SalaryAdd = calcSalaryEmployeeHasManySalaryAdd{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SalaryAdd", "types.CalcSalaryAdd"),
		CalcSalaryEmployee: struct {
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
			CalcSalary struct {
				field.RelationField
				Founder struct {
					field.RelationField
				}
				CalcSalaryEmployee struct {
					field.RelationField
				}
			}
			SalaryAdd struct {
				field.RelationField
			}
			SalaryReduce struct {
				field.RelationField
				CalcSalaryEmployee struct {
					field.RelationField
				}
				SalaryReduceItem struct {
					field.RelationField
					Employee struct {
						field.RelationField
					}
				}
			}
		}{
			RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee", "types.CalcSalaryEmployee"),
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
				RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee", "types.Employee"),
				Department: struct {
					field.RelationField
					Manager struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Department", "types.Department"),
					Manager: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Department.Manager", "types.Employee"),
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
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Rank", "types.PositionRank"),
					Grade: struct {
						field.RelationField
						Rank struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Rank.Grade", "types.PositionGrade"),
						Rank: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Rank.Grade.Rank", "types.PositionRank"),
						},
					},
				},
				Grade: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Grade", "types.PositionGrade"),
				},
				LoginInformation: struct {
					field.RelationField
					Employee struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.LoginInformation", "types.LoginInformation"),
					Employee: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.LoginInformation.Employee", "types.Employee"),
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
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Roles", "types.Role"),
					Employees: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Roles.Employees", "types.Employee"),
					},
					Permissions: struct {
						field.RelationField
						Roles struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Roles.Permissions", "types.Permission"),
						Roles: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Roles.Permissions.Roles", "types.Role"),
						},
					},
					Menus: struct {
						field.RelationField
						Roles struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Roles.Menus", "types.Menu"),
						Roles: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.Employee.Roles.Menus.Roles", "types.Role"),
						},
					},
				},
			},
			CalcSalary: struct {
				field.RelationField
				Founder struct {
					field.RelationField
				}
				CalcSalaryEmployee struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.CalcSalary", "types.CalcSalary"),
				Founder: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.CalcSalary.Founder", "types.Employee"),
				},
				CalcSalaryEmployee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.CalcSalary.CalcSalaryEmployee", "types.CalcSalaryEmployee"),
				},
			},
			SalaryAdd: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.SalaryAdd", "types.CalcSalaryAdd"),
			},
			SalaryReduce: struct {
				field.RelationField
				CalcSalaryEmployee struct {
					field.RelationField
				}
				SalaryReduceItem struct {
					field.RelationField
					Employee struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.SalaryReduce", "types.CalcSalaryReduce"),
				CalcSalaryEmployee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.SalaryReduce.CalcSalaryEmployee", "types.CalcSalaryEmployee"),
				},
				SalaryReduceItem: struct {
					field.RelationField
					Employee struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.SalaryReduce.SalaryReduceItem", "types.SalaryReduceItem"),
					Employee: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SalaryAdd.CalcSalaryEmployee.SalaryReduce.SalaryReduceItem.Employee", "types.Employee"),
					},
				},
			},
		},
		SalaryAddItem: struct {
			field.RelationField
			Employee struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("SalaryAdd.SalaryAddItem", "types.SalaryAddItem"),
			Employee: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SalaryAdd.SalaryAddItem.Employee", "types.Employee"),
			},
		},
	}

	_calcSalaryEmployee.SalaryReduce = calcSalaryEmployeeHasManySalaryReduce{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SalaryReduce", "types.CalcSalaryReduce"),
	}

	_calcSalaryEmployee.Employee = calcSalaryEmployeeBelongsToEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Employee", "types.Employee"),
	}

	_calcSalaryEmployee.CalcSalary = calcSalaryEmployeeBelongsToCalcSalary{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("CalcSalary", "types.CalcSalary"),
	}

	_calcSalaryEmployee.fillFieldMap()

	return _calcSalaryEmployee
}

type calcSalaryEmployee struct {
	calcSalaryEmployeeDo calcSalaryEmployeeDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	EmployeeID   field.Uint
	CalcSalaryID field.Uint
	Salary       field.Float32
	HourlySalary field.Float32
	SalaryAdd    calcSalaryEmployeeHasManySalaryAdd

	SalaryReduce calcSalaryEmployeeHasManySalaryReduce

	Employee calcSalaryEmployeeBelongsToEmployee

	CalcSalary calcSalaryEmployeeBelongsToCalcSalary

	fieldMap map[string]field.Expr
}

func (c calcSalaryEmployee) Table(newTableName string) *calcSalaryEmployee {
	c.calcSalaryEmployeeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c calcSalaryEmployee) As(alias string) *calcSalaryEmployee {
	c.calcSalaryEmployeeDo.DO = *(c.calcSalaryEmployeeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *calcSalaryEmployee) updateTableName(table string) *calcSalaryEmployee {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.EmployeeID = field.NewUint(table, "employee_id")
	c.CalcSalaryID = field.NewUint(table, "calc_salary_id")
	c.Salary = field.NewFloat32(table, "salary")
	c.HourlySalary = field.NewFloat32(table, "hourly_salary")

	c.fillFieldMap()

	return c
}

func (c *calcSalaryEmployee) WithContext(ctx context.Context) ICalcSalaryEmployeeDo {
	return c.calcSalaryEmployeeDo.WithContext(ctx)
}

func (c calcSalaryEmployee) TableName() string { return c.calcSalaryEmployeeDo.TableName() }

func (c calcSalaryEmployee) Alias() string { return c.calcSalaryEmployeeDo.Alias() }

func (c calcSalaryEmployee) Columns(cols ...field.Expr) gen.Columns {
	return c.calcSalaryEmployeeDo.Columns(cols...)
}

func (c *calcSalaryEmployee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *calcSalaryEmployee) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["employee_id"] = c.EmployeeID
	c.fieldMap["calc_salary_id"] = c.CalcSalaryID
	c.fieldMap["salary"] = c.Salary
	c.fieldMap["hourly_salary"] = c.HourlySalary

}

func (c calcSalaryEmployee) clone(db *gorm.DB) calcSalaryEmployee {
	c.calcSalaryEmployeeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c calcSalaryEmployee) replaceDB(db *gorm.DB) calcSalaryEmployee {
	c.calcSalaryEmployeeDo.ReplaceDB(db)
	return c
}

type calcSalaryEmployeeHasManySalaryAdd struct {
	db *gorm.DB

	field.RelationField

	CalcSalaryEmployee struct {
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
		CalcSalary struct {
			field.RelationField
			Founder struct {
				field.RelationField
			}
			CalcSalaryEmployee struct {
				field.RelationField
			}
		}
		SalaryAdd struct {
			field.RelationField
		}
		SalaryReduce struct {
			field.RelationField
			CalcSalaryEmployee struct {
				field.RelationField
			}
			SalaryReduceItem struct {
				field.RelationField
				Employee struct {
					field.RelationField
				}
			}
		}
	}
	SalaryAddItem struct {
		field.RelationField
		Employee struct {
			field.RelationField
		}
	}
}

func (a calcSalaryEmployeeHasManySalaryAdd) Where(conds ...field.Expr) *calcSalaryEmployeeHasManySalaryAdd {
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

func (a calcSalaryEmployeeHasManySalaryAdd) WithContext(ctx context.Context) *calcSalaryEmployeeHasManySalaryAdd {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a calcSalaryEmployeeHasManySalaryAdd) Session(session *gorm.Session) *calcSalaryEmployeeHasManySalaryAdd {
	a.db = a.db.Session(session)
	return &a
}

func (a calcSalaryEmployeeHasManySalaryAdd) Model(m *types.CalcSalaryEmployee) *calcSalaryEmployeeHasManySalaryAddTx {
	return &calcSalaryEmployeeHasManySalaryAddTx{a.db.Model(m).Association(a.Name())}
}

type calcSalaryEmployeeHasManySalaryAddTx struct{ tx *gorm.Association }

func (a calcSalaryEmployeeHasManySalaryAddTx) Find() (result []*types.CalcSalaryAdd, err error) {
	return result, a.tx.Find(&result)
}

func (a calcSalaryEmployeeHasManySalaryAddTx) Append(values ...*types.CalcSalaryAdd) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a calcSalaryEmployeeHasManySalaryAddTx) Replace(values ...*types.CalcSalaryAdd) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a calcSalaryEmployeeHasManySalaryAddTx) Delete(values ...*types.CalcSalaryAdd) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a calcSalaryEmployeeHasManySalaryAddTx) Clear() error {
	return a.tx.Clear()
}

func (a calcSalaryEmployeeHasManySalaryAddTx) Count() int64 {
	return a.tx.Count()
}

type calcSalaryEmployeeHasManySalaryReduce struct {
	db *gorm.DB

	field.RelationField
}

func (a calcSalaryEmployeeHasManySalaryReduce) Where(conds ...field.Expr) *calcSalaryEmployeeHasManySalaryReduce {
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

func (a calcSalaryEmployeeHasManySalaryReduce) WithContext(ctx context.Context) *calcSalaryEmployeeHasManySalaryReduce {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a calcSalaryEmployeeHasManySalaryReduce) Session(session *gorm.Session) *calcSalaryEmployeeHasManySalaryReduce {
	a.db = a.db.Session(session)
	return &a
}

func (a calcSalaryEmployeeHasManySalaryReduce) Model(m *types.CalcSalaryEmployee) *calcSalaryEmployeeHasManySalaryReduceTx {
	return &calcSalaryEmployeeHasManySalaryReduceTx{a.db.Model(m).Association(a.Name())}
}

type calcSalaryEmployeeHasManySalaryReduceTx struct{ tx *gorm.Association }

func (a calcSalaryEmployeeHasManySalaryReduceTx) Find() (result []*types.CalcSalaryReduce, err error) {
	return result, a.tx.Find(&result)
}

func (a calcSalaryEmployeeHasManySalaryReduceTx) Append(values ...*types.CalcSalaryReduce) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a calcSalaryEmployeeHasManySalaryReduceTx) Replace(values ...*types.CalcSalaryReduce) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a calcSalaryEmployeeHasManySalaryReduceTx) Delete(values ...*types.CalcSalaryReduce) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a calcSalaryEmployeeHasManySalaryReduceTx) Clear() error {
	return a.tx.Clear()
}

func (a calcSalaryEmployeeHasManySalaryReduceTx) Count() int64 {
	return a.tx.Count()
}

type calcSalaryEmployeeBelongsToEmployee struct {
	db *gorm.DB

	field.RelationField
}

func (a calcSalaryEmployeeBelongsToEmployee) Where(conds ...field.Expr) *calcSalaryEmployeeBelongsToEmployee {
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

func (a calcSalaryEmployeeBelongsToEmployee) WithContext(ctx context.Context) *calcSalaryEmployeeBelongsToEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a calcSalaryEmployeeBelongsToEmployee) Session(session *gorm.Session) *calcSalaryEmployeeBelongsToEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a calcSalaryEmployeeBelongsToEmployee) Model(m *types.CalcSalaryEmployee) *calcSalaryEmployeeBelongsToEmployeeTx {
	return &calcSalaryEmployeeBelongsToEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type calcSalaryEmployeeBelongsToEmployeeTx struct{ tx *gorm.Association }

func (a calcSalaryEmployeeBelongsToEmployeeTx) Find() (result *types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a calcSalaryEmployeeBelongsToEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a calcSalaryEmployeeBelongsToEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a calcSalaryEmployeeBelongsToEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a calcSalaryEmployeeBelongsToEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a calcSalaryEmployeeBelongsToEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type calcSalaryEmployeeBelongsToCalcSalary struct {
	db *gorm.DB

	field.RelationField
}

func (a calcSalaryEmployeeBelongsToCalcSalary) Where(conds ...field.Expr) *calcSalaryEmployeeBelongsToCalcSalary {
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

func (a calcSalaryEmployeeBelongsToCalcSalary) WithContext(ctx context.Context) *calcSalaryEmployeeBelongsToCalcSalary {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a calcSalaryEmployeeBelongsToCalcSalary) Session(session *gorm.Session) *calcSalaryEmployeeBelongsToCalcSalary {
	a.db = a.db.Session(session)
	return &a
}

func (a calcSalaryEmployeeBelongsToCalcSalary) Model(m *types.CalcSalaryEmployee) *calcSalaryEmployeeBelongsToCalcSalaryTx {
	return &calcSalaryEmployeeBelongsToCalcSalaryTx{a.db.Model(m).Association(a.Name())}
}

type calcSalaryEmployeeBelongsToCalcSalaryTx struct{ tx *gorm.Association }

func (a calcSalaryEmployeeBelongsToCalcSalaryTx) Find() (result *types.CalcSalary, err error) {
	return result, a.tx.Find(&result)
}

func (a calcSalaryEmployeeBelongsToCalcSalaryTx) Append(values ...*types.CalcSalary) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a calcSalaryEmployeeBelongsToCalcSalaryTx) Replace(values ...*types.CalcSalary) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a calcSalaryEmployeeBelongsToCalcSalaryTx) Delete(values ...*types.CalcSalary) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a calcSalaryEmployeeBelongsToCalcSalaryTx) Clear() error {
	return a.tx.Clear()
}

func (a calcSalaryEmployeeBelongsToCalcSalaryTx) Count() int64 {
	return a.tx.Count()
}

type calcSalaryEmployeeDo struct{ gen.DO }

type ICalcSalaryEmployeeDo interface {
	gen.SubQuery
	Debug() ICalcSalaryEmployeeDo
	WithContext(ctx context.Context) ICalcSalaryEmployeeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICalcSalaryEmployeeDo
	WriteDB() ICalcSalaryEmployeeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICalcSalaryEmployeeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICalcSalaryEmployeeDo
	Not(conds ...gen.Condition) ICalcSalaryEmployeeDo
	Or(conds ...gen.Condition) ICalcSalaryEmployeeDo
	Select(conds ...field.Expr) ICalcSalaryEmployeeDo
	Where(conds ...gen.Condition) ICalcSalaryEmployeeDo
	Order(conds ...field.Expr) ICalcSalaryEmployeeDo
	Distinct(cols ...field.Expr) ICalcSalaryEmployeeDo
	Omit(cols ...field.Expr) ICalcSalaryEmployeeDo
	Join(table schema.Tabler, on ...field.Expr) ICalcSalaryEmployeeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICalcSalaryEmployeeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICalcSalaryEmployeeDo
	Group(cols ...field.Expr) ICalcSalaryEmployeeDo
	Having(conds ...gen.Condition) ICalcSalaryEmployeeDo
	Limit(limit int) ICalcSalaryEmployeeDo
	Offset(offset int) ICalcSalaryEmployeeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICalcSalaryEmployeeDo
	Unscoped() ICalcSalaryEmployeeDo
	Create(values ...*types.CalcSalaryEmployee) error
	CreateInBatches(values []*types.CalcSalaryEmployee, batchSize int) error
	Save(values ...*types.CalcSalaryEmployee) error
	First() (*types.CalcSalaryEmployee, error)
	Take() (*types.CalcSalaryEmployee, error)
	Last() (*types.CalcSalaryEmployee, error)
	Find() ([]*types.CalcSalaryEmployee, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.CalcSalaryEmployee, err error)
	FindInBatches(result *[]*types.CalcSalaryEmployee, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.CalcSalaryEmployee) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICalcSalaryEmployeeDo
	Assign(attrs ...field.AssignExpr) ICalcSalaryEmployeeDo
	Joins(fields ...field.RelationField) ICalcSalaryEmployeeDo
	Preload(fields ...field.RelationField) ICalcSalaryEmployeeDo
	FirstOrInit() (*types.CalcSalaryEmployee, error)
	FirstOrCreate() (*types.CalcSalaryEmployee, error)
	FindByPage(offset int, limit int) (result []*types.CalcSalaryEmployee, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICalcSalaryEmployeeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c calcSalaryEmployeeDo) Debug() ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Debug())
}

func (c calcSalaryEmployeeDo) WithContext(ctx context.Context) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c calcSalaryEmployeeDo) ReadDB() ICalcSalaryEmployeeDo {
	return c.Clauses(dbresolver.Read)
}

func (c calcSalaryEmployeeDo) WriteDB() ICalcSalaryEmployeeDo {
	return c.Clauses(dbresolver.Write)
}

func (c calcSalaryEmployeeDo) Session(config *gorm.Session) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Session(config))
}

func (c calcSalaryEmployeeDo) Clauses(conds ...clause.Expression) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c calcSalaryEmployeeDo) Returning(value interface{}, columns ...string) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c calcSalaryEmployeeDo) Not(conds ...gen.Condition) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c calcSalaryEmployeeDo) Or(conds ...gen.Condition) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c calcSalaryEmployeeDo) Select(conds ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c calcSalaryEmployeeDo) Where(conds ...gen.Condition) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c calcSalaryEmployeeDo) Order(conds ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c calcSalaryEmployeeDo) Distinct(cols ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c calcSalaryEmployeeDo) Omit(cols ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c calcSalaryEmployeeDo) Join(table schema.Tabler, on ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c calcSalaryEmployeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c calcSalaryEmployeeDo) RightJoin(table schema.Tabler, on ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c calcSalaryEmployeeDo) Group(cols ...field.Expr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c calcSalaryEmployeeDo) Having(conds ...gen.Condition) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c calcSalaryEmployeeDo) Limit(limit int) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c calcSalaryEmployeeDo) Offset(offset int) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c calcSalaryEmployeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c calcSalaryEmployeeDo) Unscoped() ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c calcSalaryEmployeeDo) Create(values ...*types.CalcSalaryEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c calcSalaryEmployeeDo) CreateInBatches(values []*types.CalcSalaryEmployee, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c calcSalaryEmployeeDo) Save(values ...*types.CalcSalaryEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c calcSalaryEmployeeDo) First() (*types.CalcSalaryEmployee, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.CalcSalaryEmployee), nil
	}
}

func (c calcSalaryEmployeeDo) Take() (*types.CalcSalaryEmployee, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.CalcSalaryEmployee), nil
	}
}

func (c calcSalaryEmployeeDo) Last() (*types.CalcSalaryEmployee, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.CalcSalaryEmployee), nil
	}
}

func (c calcSalaryEmployeeDo) Find() ([]*types.CalcSalaryEmployee, error) {
	result, err := c.DO.Find()
	return result.([]*types.CalcSalaryEmployee), err
}

func (c calcSalaryEmployeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.CalcSalaryEmployee, err error) {
	buf := make([]*types.CalcSalaryEmployee, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c calcSalaryEmployeeDo) FindInBatches(result *[]*types.CalcSalaryEmployee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c calcSalaryEmployeeDo) Attrs(attrs ...field.AssignExpr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c calcSalaryEmployeeDo) Assign(attrs ...field.AssignExpr) ICalcSalaryEmployeeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c calcSalaryEmployeeDo) Joins(fields ...field.RelationField) ICalcSalaryEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c calcSalaryEmployeeDo) Preload(fields ...field.RelationField) ICalcSalaryEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c calcSalaryEmployeeDo) FirstOrInit() (*types.CalcSalaryEmployee, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.CalcSalaryEmployee), nil
	}
}

func (c calcSalaryEmployeeDo) FirstOrCreate() (*types.CalcSalaryEmployee, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.CalcSalaryEmployee), nil
	}
}

func (c calcSalaryEmployeeDo) FindByPage(offset int, limit int) (result []*types.CalcSalaryEmployee, count int64, err error) {
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

func (c calcSalaryEmployeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c calcSalaryEmployeeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c calcSalaryEmployeeDo) Delete(models ...*types.CalcSalaryEmployee) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *calcSalaryEmployeeDo) withDO(do gen.Dao) *calcSalaryEmployeeDo {
	c.DO = *do.(*gen.DO)
	return c
}
