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

func newOvertimeRequestFormRate(db *gorm.DB, opts ...gen.DOOption) overtimeRequestFormRate {
	_overtimeRequestFormRate := overtimeRequestFormRate{}

	_overtimeRequestFormRate.overtimeRequestFormRateDo.UseDB(db, opts...)
	_overtimeRequestFormRate.overtimeRequestFormRateDo.UseModel(&types.OvertimeRequestFormRate{})

	tableName := _overtimeRequestFormRate.overtimeRequestFormRateDo.TableName()
	_overtimeRequestFormRate.ALL = field.NewAsterisk(tableName)
	_overtimeRequestFormRate.ID = field.NewUint(tableName, "id")
	_overtimeRequestFormRate.CreatedAt = field.NewTime(tableName, "created_at")
	_overtimeRequestFormRate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_overtimeRequestFormRate.DeletedAt = field.NewField(tableName, "deleted_at")
	_overtimeRequestFormRate.OvertimeRequestFormID = field.NewUint(tableName, "overtime_request_form_id")
	_overtimeRequestFormRate.Hours = field.NewUint(tableName, "hours")
	_overtimeRequestFormRate.Multiply = field.NewFloat32(tableName, "multiply")
	_overtimeRequestFormRate.Level = field.NewUint(tableName, "level")
	_overtimeRequestFormRate.IsFill = field.NewBool(tableName, "is_fill")
	_overtimeRequestFormRate.Fill = field.NewUint(tableName, "fill")
	_overtimeRequestFormRate.OvertimeRequestForm = overtimeRequestFormRateBelongsToOvertimeRequestForm{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("OvertimeRequestForm", "types.OvertimeRequestForm"),
		Vacation: struct {
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
				}
				Employee struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("OvertimeRequestForm.Vacation", "types.Vacation"),
			Schedule: struct {
				field.RelationField
				Vacation struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("OvertimeRequestForm.Vacation.Schedule", "types.VacationSchedule"),
				Vacation: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("OvertimeRequestForm.Vacation.Schedule.Vacation", "types.Vacation"),
				},
			},
			VacationGroup: struct {
				field.RelationField
				Vacation struct {
					field.RelationField
				}
				VacationGroupOvertimeRate struct {
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
				}
				Employee struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup", "types.VacationGroup"),
				Vacation: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.Vacation", "types.Vacation"),
				},
				VacationGroupOvertimeRate: struct {
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
				}{
					RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate", "types.VacationGroupOvertimeRate"),
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
					}{
						RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup", "types.LeaveGroup"),
						Leave: struct {
							field.RelationField
							LeaveGroup struct {
								field.RelationField
							}
						}{
							RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Leave", "types.Leave"),
							LeaveGroup: struct {
								field.RelationField
							}{
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Leave.LeaveGroup", "types.LeaveGroup"),
							},
						},
						LeaveGroupCondition: struct {
							field.RelationField
							LeaveGroup struct {
								field.RelationField
							}
						}{
							RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.LeaveGroupCondition", "types.LeaveGroupCondition"),
							LeaveGroup: struct {
								field.RelationField
							}{
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.LeaveGroupCondition.LeaveGroup", "types.LeaveGroup"),
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
							RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee", "types.Employee"),
							Department: struct {
								field.RelationField
								Manager struct {
									field.RelationField
								}
							}{
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Department", "types.Department"),
								Manager: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Department.Manager", "types.Employee"),
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
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Rank", "types.PositionRank"),
								Grade: struct {
									field.RelationField
									Rank struct {
										field.RelationField
									}
								}{
									RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Rank.Grade", "types.PositionGrade"),
									Rank: struct {
										field.RelationField
									}{
										RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Rank.Grade.Rank", "types.PositionRank"),
									},
								},
							},
							Grade: struct {
								field.RelationField
							}{
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Grade", "types.PositionGrade"),
							},
							LoginInformation: struct {
								field.RelationField
								Employee struct {
									field.RelationField
								}
							}{
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.LoginInformation", "types.LoginInformation"),
								Employee: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.LoginInformation.Employee", "types.Employee"),
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
								RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Roles", "types.Role"),
								Employees: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Employees", "types.Employee"),
								},
								Permissions: struct {
									field.RelationField
									Roles struct {
										field.RelationField
									}
								}{
									RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Permissions", "types.Permission"),
									Roles: struct {
										field.RelationField
									}{
										RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Permissions.Roles", "types.Role"),
									},
								},
								Menus: struct {
									field.RelationField
									Roles struct {
										field.RelationField
									}
								}{
									RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Menus", "types.Menu"),
									Roles: struct {
										field.RelationField
									}{
										RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.VacationGroupOvertimeRate.VacationGroup.Employee.Roles.Menus.Roles", "types.Role"),
									},
								},
							},
						},
					},
				},
				Employee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("OvertimeRequestForm.Vacation.VacationGroup.Employee", "types.Employee"),
				},
			},
		},
		Employee: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("OvertimeRequestForm.Employee", "types.Employee"),
		},
		Department: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("OvertimeRequestForm.Department", "types.Department"),
		},
		SignOffFlow: struct {
			field.RelationField
			OvertimeRequestForm struct {
				field.RelationField
			}
			SignOffEmployee struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("OvertimeRequestForm.SignOffFlow", "types.OvertimeSignOffFlow"),
			OvertimeRequestForm: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("OvertimeRequestForm.SignOffFlow.OvertimeRequestForm", "types.OvertimeRequestForm"),
			},
			SignOffEmployee: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("OvertimeRequestForm.SignOffFlow.SignOffEmployee", "types.Employee"),
			},
		},
		Rate: struct {
			field.RelationField
			OvertimeRequestForm struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("OvertimeRequestForm.Rate", "types.OvertimeRequestFormRate"),
			OvertimeRequestForm: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("OvertimeRequestForm.Rate.OvertimeRequestForm", "types.OvertimeRequestForm"),
			},
		},
	}

	_overtimeRequestFormRate.fillFieldMap()

	return _overtimeRequestFormRate
}

type overtimeRequestFormRate struct {
	overtimeRequestFormRateDo overtimeRequestFormRateDo

	ALL                   field.Asterisk
	ID                    field.Uint
	CreatedAt             field.Time
	UpdatedAt             field.Time
	DeletedAt             field.Field
	OvertimeRequestFormID field.Uint
	Hours                 field.Uint
	Multiply              field.Float32
	Level                 field.Uint
	IsFill                field.Bool
	Fill                  field.Uint
	OvertimeRequestForm   overtimeRequestFormRateBelongsToOvertimeRequestForm

	fieldMap map[string]field.Expr
}

func (o overtimeRequestFormRate) Table(newTableName string) *overtimeRequestFormRate {
	o.overtimeRequestFormRateDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o overtimeRequestFormRate) As(alias string) *overtimeRequestFormRate {
	o.overtimeRequestFormRateDo.DO = *(o.overtimeRequestFormRateDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *overtimeRequestFormRate) updateTableName(table string) *overtimeRequestFormRate {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewUint(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.OvertimeRequestFormID = field.NewUint(table, "overtime_request_form_id")
	o.Hours = field.NewUint(table, "hours")
	o.Multiply = field.NewFloat32(table, "multiply")
	o.Level = field.NewUint(table, "level")
	o.IsFill = field.NewBool(table, "is_fill")
	o.Fill = field.NewUint(table, "fill")

	o.fillFieldMap()

	return o
}

func (o *overtimeRequestFormRate) WithContext(ctx context.Context) IOvertimeRequestFormRateDo {
	return o.overtimeRequestFormRateDo.WithContext(ctx)
}

func (o overtimeRequestFormRate) TableName() string { return o.overtimeRequestFormRateDo.TableName() }

func (o overtimeRequestFormRate) Alias() string { return o.overtimeRequestFormRateDo.Alias() }

func (o overtimeRequestFormRate) Columns(cols ...field.Expr) gen.Columns {
	return o.overtimeRequestFormRateDo.Columns(cols...)
}

func (o *overtimeRequestFormRate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *overtimeRequestFormRate) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 11)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["overtime_request_form_id"] = o.OvertimeRequestFormID
	o.fieldMap["hours"] = o.Hours
	o.fieldMap["multiply"] = o.Multiply
	o.fieldMap["level"] = o.Level
	o.fieldMap["is_fill"] = o.IsFill
	o.fieldMap["fill"] = o.Fill

}

func (o overtimeRequestFormRate) clone(db *gorm.DB) overtimeRequestFormRate {
	o.overtimeRequestFormRateDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o overtimeRequestFormRate) replaceDB(db *gorm.DB) overtimeRequestFormRate {
	o.overtimeRequestFormRateDo.ReplaceDB(db)
	return o
}

type overtimeRequestFormRateBelongsToOvertimeRequestForm struct {
	db *gorm.DB

	field.RelationField

	Vacation struct {
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
			}
			Employee struct {
				field.RelationField
			}
		}
	}
	Employee struct {
		field.RelationField
	}
	Department struct {
		field.RelationField
	}
	SignOffFlow struct {
		field.RelationField
		OvertimeRequestForm struct {
			field.RelationField
		}
		SignOffEmployee struct {
			field.RelationField
		}
	}
	Rate struct {
		field.RelationField
		OvertimeRequestForm struct {
			field.RelationField
		}
	}
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestForm) Where(conds ...field.Expr) *overtimeRequestFormRateBelongsToOvertimeRequestForm {
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

func (a overtimeRequestFormRateBelongsToOvertimeRequestForm) WithContext(ctx context.Context) *overtimeRequestFormRateBelongsToOvertimeRequestForm {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestForm) Session(session *gorm.Session) *overtimeRequestFormRateBelongsToOvertimeRequestForm {
	a.db = a.db.Session(session)
	return &a
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestForm) Model(m *types.OvertimeRequestFormRate) *overtimeRequestFormRateBelongsToOvertimeRequestFormTx {
	return &overtimeRequestFormRateBelongsToOvertimeRequestFormTx{a.db.Model(m).Association(a.Name())}
}

type overtimeRequestFormRateBelongsToOvertimeRequestFormTx struct{ tx *gorm.Association }

func (a overtimeRequestFormRateBelongsToOvertimeRequestFormTx) Find() (result *types.OvertimeRequestForm, err error) {
	return result, a.tx.Find(&result)
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestFormTx) Append(values ...*types.OvertimeRequestForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestFormTx) Replace(values ...*types.OvertimeRequestForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestFormTx) Delete(values ...*types.OvertimeRequestForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestFormTx) Clear() error {
	return a.tx.Clear()
}

func (a overtimeRequestFormRateBelongsToOvertimeRequestFormTx) Count() int64 {
	return a.tx.Count()
}

type overtimeRequestFormRateDo struct{ gen.DO }

type IOvertimeRequestFormRateDo interface {
	gen.SubQuery
	Debug() IOvertimeRequestFormRateDo
	WithContext(ctx context.Context) IOvertimeRequestFormRateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOvertimeRequestFormRateDo
	WriteDB() IOvertimeRequestFormRateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOvertimeRequestFormRateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOvertimeRequestFormRateDo
	Not(conds ...gen.Condition) IOvertimeRequestFormRateDo
	Or(conds ...gen.Condition) IOvertimeRequestFormRateDo
	Select(conds ...field.Expr) IOvertimeRequestFormRateDo
	Where(conds ...gen.Condition) IOvertimeRequestFormRateDo
	Order(conds ...field.Expr) IOvertimeRequestFormRateDo
	Distinct(cols ...field.Expr) IOvertimeRequestFormRateDo
	Omit(cols ...field.Expr) IOvertimeRequestFormRateDo
	Join(table schema.Tabler, on ...field.Expr) IOvertimeRequestFormRateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOvertimeRequestFormRateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOvertimeRequestFormRateDo
	Group(cols ...field.Expr) IOvertimeRequestFormRateDo
	Having(conds ...gen.Condition) IOvertimeRequestFormRateDo
	Limit(limit int) IOvertimeRequestFormRateDo
	Offset(offset int) IOvertimeRequestFormRateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOvertimeRequestFormRateDo
	Unscoped() IOvertimeRequestFormRateDo
	Create(values ...*types.OvertimeRequestFormRate) error
	CreateInBatches(values []*types.OvertimeRequestFormRate, batchSize int) error
	Save(values ...*types.OvertimeRequestFormRate) error
	First() (*types.OvertimeRequestFormRate, error)
	Take() (*types.OvertimeRequestFormRate, error)
	Last() (*types.OvertimeRequestFormRate, error)
	Find() ([]*types.OvertimeRequestFormRate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.OvertimeRequestFormRate, err error)
	FindInBatches(result *[]*types.OvertimeRequestFormRate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.OvertimeRequestFormRate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOvertimeRequestFormRateDo
	Assign(attrs ...field.AssignExpr) IOvertimeRequestFormRateDo
	Joins(fields ...field.RelationField) IOvertimeRequestFormRateDo
	Preload(fields ...field.RelationField) IOvertimeRequestFormRateDo
	FirstOrInit() (*types.OvertimeRequestFormRate, error)
	FirstOrCreate() (*types.OvertimeRequestFormRate, error)
	FindByPage(offset int, limit int) (result []*types.OvertimeRequestFormRate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOvertimeRequestFormRateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o overtimeRequestFormRateDo) Debug() IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Debug())
}

func (o overtimeRequestFormRateDo) WithContext(ctx context.Context) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o overtimeRequestFormRateDo) ReadDB() IOvertimeRequestFormRateDo {
	return o.Clauses(dbresolver.Read)
}

func (o overtimeRequestFormRateDo) WriteDB() IOvertimeRequestFormRateDo {
	return o.Clauses(dbresolver.Write)
}

func (o overtimeRequestFormRateDo) Session(config *gorm.Session) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Session(config))
}

func (o overtimeRequestFormRateDo) Clauses(conds ...clause.Expression) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o overtimeRequestFormRateDo) Returning(value interface{}, columns ...string) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o overtimeRequestFormRateDo) Not(conds ...gen.Condition) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o overtimeRequestFormRateDo) Or(conds ...gen.Condition) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o overtimeRequestFormRateDo) Select(conds ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o overtimeRequestFormRateDo) Where(conds ...gen.Condition) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o overtimeRequestFormRateDo) Order(conds ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o overtimeRequestFormRateDo) Distinct(cols ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o overtimeRequestFormRateDo) Omit(cols ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o overtimeRequestFormRateDo) Join(table schema.Tabler, on ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o overtimeRequestFormRateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o overtimeRequestFormRateDo) RightJoin(table schema.Tabler, on ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o overtimeRequestFormRateDo) Group(cols ...field.Expr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o overtimeRequestFormRateDo) Having(conds ...gen.Condition) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o overtimeRequestFormRateDo) Limit(limit int) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o overtimeRequestFormRateDo) Offset(offset int) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o overtimeRequestFormRateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o overtimeRequestFormRateDo) Unscoped() IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Unscoped())
}

func (o overtimeRequestFormRateDo) Create(values ...*types.OvertimeRequestFormRate) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o overtimeRequestFormRateDo) CreateInBatches(values []*types.OvertimeRequestFormRate, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o overtimeRequestFormRateDo) Save(values ...*types.OvertimeRequestFormRate) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o overtimeRequestFormRateDo) First() (*types.OvertimeRequestFormRate, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeRequestFormRate), nil
	}
}

func (o overtimeRequestFormRateDo) Take() (*types.OvertimeRequestFormRate, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeRequestFormRate), nil
	}
}

func (o overtimeRequestFormRateDo) Last() (*types.OvertimeRequestFormRate, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeRequestFormRate), nil
	}
}

func (o overtimeRequestFormRateDo) Find() ([]*types.OvertimeRequestFormRate, error) {
	result, err := o.DO.Find()
	return result.([]*types.OvertimeRequestFormRate), err
}

func (o overtimeRequestFormRateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.OvertimeRequestFormRate, err error) {
	buf := make([]*types.OvertimeRequestFormRate, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o overtimeRequestFormRateDo) FindInBatches(result *[]*types.OvertimeRequestFormRate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o overtimeRequestFormRateDo) Attrs(attrs ...field.AssignExpr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o overtimeRequestFormRateDo) Assign(attrs ...field.AssignExpr) IOvertimeRequestFormRateDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o overtimeRequestFormRateDo) Joins(fields ...field.RelationField) IOvertimeRequestFormRateDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o overtimeRequestFormRateDo) Preload(fields ...field.RelationField) IOvertimeRequestFormRateDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o overtimeRequestFormRateDo) FirstOrInit() (*types.OvertimeRequestFormRate, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeRequestFormRate), nil
	}
}

func (o overtimeRequestFormRateDo) FirstOrCreate() (*types.OvertimeRequestFormRate, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeRequestFormRate), nil
	}
}

func (o overtimeRequestFormRateDo) FindByPage(offset int, limit int) (result []*types.OvertimeRequestFormRate, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o overtimeRequestFormRateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o overtimeRequestFormRateDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o overtimeRequestFormRateDo) Delete(models ...*types.OvertimeRequestFormRate) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *overtimeRequestFormRateDo) withDO(do gen.Dao) *overtimeRequestFormRateDo {
	o.DO = *do.(*gen.DO)
	return o
}