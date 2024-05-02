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

func newOvertimeSignOffFlow(db *gorm.DB, opts ...gen.DOOption) overtimeSignOffFlow {
	_overtimeSignOffFlow := overtimeSignOffFlow{}

	_overtimeSignOffFlow.overtimeSignOffFlowDo.UseDB(db, opts...)
	_overtimeSignOffFlow.overtimeSignOffFlowDo.UseModel(&types.OvertimeSignOffFlow{})

	tableName := _overtimeSignOffFlow.overtimeSignOffFlowDo.TableName()
	_overtimeSignOffFlow.ALL = field.NewAsterisk(tableName)
	_overtimeSignOffFlow.ID = field.NewUint(tableName, "id")
	_overtimeSignOffFlow.CreatedAt = field.NewTime(tableName, "created_at")
	_overtimeSignOffFlow.UpdatedAt = field.NewTime(tableName, "updated_at")
	_overtimeSignOffFlow.DeletedAt = field.NewField(tableName, "deleted_at")
	_overtimeSignOffFlow.OvertimeRequestFormID = field.NewUint(tableName, "overtime_request_form_id")
	_overtimeSignOffFlow.SignOffEmployeeID = field.NewUint(tableName, "sign_off_employee_id")
	_overtimeSignOffFlow.Level = field.NewUint(tableName, "level")
	_overtimeSignOffFlow.SignType = field.NewField(tableName, "sign_type")
	_overtimeSignOffFlow.Notify = field.NewField(tableName, "notify")
	_overtimeSignOffFlow.Remark = field.NewString(tableName, "remark")
	_overtimeSignOffFlow.Comment = field.NewString(tableName, "comment")
	_overtimeSignOffFlow.Status = field.NewField(tableName, "status")
	_overtimeSignOffFlow.SignDate = field.NewTime(tableName, "sign_date")
	_overtimeSignOffFlow.UUID = field.NewString(tableName, "uuid")
	_overtimeSignOffFlow.Locale = field.NewField(tableName, "locale")
	_overtimeSignOffFlow.OvertimeRequestForm = overtimeSignOffFlowBelongsToOvertimeRequestForm{
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
	}

	_overtimeSignOffFlow.SignOffEmployee = overtimeSignOffFlowBelongsToSignOffEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SignOffEmployee", "types.Employee"),
	}

	_overtimeSignOffFlow.fillFieldMap()

	return _overtimeSignOffFlow
}

type overtimeSignOffFlow struct {
	overtimeSignOffFlowDo overtimeSignOffFlowDo

	ALL                   field.Asterisk
	ID                    field.Uint
	CreatedAt             field.Time
	UpdatedAt             field.Time
	DeletedAt             field.Field
	OvertimeRequestFormID field.Uint
	SignOffEmployeeID     field.Uint
	Level                 field.Uint
	SignType              field.Field
	Notify                field.Field
	Remark                field.String
	Comment               field.String
	Status                field.Field
	SignDate              field.Time
	UUID                  field.String
	Locale                field.Field
	OvertimeRequestForm   overtimeSignOffFlowBelongsToOvertimeRequestForm

	SignOffEmployee overtimeSignOffFlowBelongsToSignOffEmployee

	fieldMap map[string]field.Expr
}

func (o overtimeSignOffFlow) Table(newTableName string) *overtimeSignOffFlow {
	o.overtimeSignOffFlowDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o overtimeSignOffFlow) As(alias string) *overtimeSignOffFlow {
	o.overtimeSignOffFlowDo.DO = *(o.overtimeSignOffFlowDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *overtimeSignOffFlow) updateTableName(table string) *overtimeSignOffFlow {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewUint(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.OvertimeRequestFormID = field.NewUint(table, "overtime_request_form_id")
	o.SignOffEmployeeID = field.NewUint(table, "sign_off_employee_id")
	o.Level = field.NewUint(table, "level")
	o.SignType = field.NewField(table, "sign_type")
	o.Notify = field.NewField(table, "notify")
	o.Remark = field.NewString(table, "remark")
	o.Comment = field.NewString(table, "comment")
	o.Status = field.NewField(table, "status")
	o.SignDate = field.NewTime(table, "sign_date")
	o.UUID = field.NewString(table, "uuid")
	o.Locale = field.NewField(table, "locale")

	o.fillFieldMap()

	return o
}

func (o *overtimeSignOffFlow) WithContext(ctx context.Context) IOvertimeSignOffFlowDo {
	return o.overtimeSignOffFlowDo.WithContext(ctx)
}

func (o overtimeSignOffFlow) TableName() string { return o.overtimeSignOffFlowDo.TableName() }

func (o overtimeSignOffFlow) Alias() string { return o.overtimeSignOffFlowDo.Alias() }

func (o overtimeSignOffFlow) Columns(cols ...field.Expr) gen.Columns {
	return o.overtimeSignOffFlowDo.Columns(cols...)
}

func (o *overtimeSignOffFlow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *overtimeSignOffFlow) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 17)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["overtime_request_form_id"] = o.OvertimeRequestFormID
	o.fieldMap["sign_off_employee_id"] = o.SignOffEmployeeID
	o.fieldMap["level"] = o.Level
	o.fieldMap["sign_type"] = o.SignType
	o.fieldMap["notify"] = o.Notify
	o.fieldMap["remark"] = o.Remark
	o.fieldMap["comment"] = o.Comment
	o.fieldMap["status"] = o.Status
	o.fieldMap["sign_date"] = o.SignDate
	o.fieldMap["uuid"] = o.UUID
	o.fieldMap["locale"] = o.Locale

}

func (o overtimeSignOffFlow) clone(db *gorm.DB) overtimeSignOffFlow {
	o.overtimeSignOffFlowDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o overtimeSignOffFlow) replaceDB(db *gorm.DB) overtimeSignOffFlow {
	o.overtimeSignOffFlowDo.ReplaceDB(db)
	return o
}

type overtimeSignOffFlowBelongsToOvertimeRequestForm struct {
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
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestForm) Where(conds ...field.Expr) *overtimeSignOffFlowBelongsToOvertimeRequestForm {
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

func (a overtimeSignOffFlowBelongsToOvertimeRequestForm) WithContext(ctx context.Context) *overtimeSignOffFlowBelongsToOvertimeRequestForm {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestForm) Session(session *gorm.Session) *overtimeSignOffFlowBelongsToOvertimeRequestForm {
	a.db = a.db.Session(session)
	return &a
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestForm) Model(m *types.OvertimeSignOffFlow) *overtimeSignOffFlowBelongsToOvertimeRequestFormTx {
	return &overtimeSignOffFlowBelongsToOvertimeRequestFormTx{a.db.Model(m).Association(a.Name())}
}

type overtimeSignOffFlowBelongsToOvertimeRequestFormTx struct{ tx *gorm.Association }

func (a overtimeSignOffFlowBelongsToOvertimeRequestFormTx) Find() (result *types.OvertimeRequestForm, err error) {
	return result, a.tx.Find(&result)
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestFormTx) Append(values ...*types.OvertimeRequestForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestFormTx) Replace(values ...*types.OvertimeRequestForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestFormTx) Delete(values ...*types.OvertimeRequestForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestFormTx) Clear() error {
	return a.tx.Clear()
}

func (a overtimeSignOffFlowBelongsToOvertimeRequestFormTx) Count() int64 {
	return a.tx.Count()
}

type overtimeSignOffFlowBelongsToSignOffEmployee struct {
	db *gorm.DB

	field.RelationField
}

func (a overtimeSignOffFlowBelongsToSignOffEmployee) Where(conds ...field.Expr) *overtimeSignOffFlowBelongsToSignOffEmployee {
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

func (a overtimeSignOffFlowBelongsToSignOffEmployee) WithContext(ctx context.Context) *overtimeSignOffFlowBelongsToSignOffEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a overtimeSignOffFlowBelongsToSignOffEmployee) Session(session *gorm.Session) *overtimeSignOffFlowBelongsToSignOffEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a overtimeSignOffFlowBelongsToSignOffEmployee) Model(m *types.OvertimeSignOffFlow) *overtimeSignOffFlowBelongsToSignOffEmployeeTx {
	return &overtimeSignOffFlowBelongsToSignOffEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type overtimeSignOffFlowBelongsToSignOffEmployeeTx struct{ tx *gorm.Association }

func (a overtimeSignOffFlowBelongsToSignOffEmployeeTx) Find() (result *types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a overtimeSignOffFlowBelongsToSignOffEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a overtimeSignOffFlowBelongsToSignOffEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a overtimeSignOffFlowBelongsToSignOffEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a overtimeSignOffFlowBelongsToSignOffEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a overtimeSignOffFlowBelongsToSignOffEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type overtimeSignOffFlowDo struct{ gen.DO }

type IOvertimeSignOffFlowDo interface {
	gen.SubQuery
	Debug() IOvertimeSignOffFlowDo
	WithContext(ctx context.Context) IOvertimeSignOffFlowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOvertimeSignOffFlowDo
	WriteDB() IOvertimeSignOffFlowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOvertimeSignOffFlowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOvertimeSignOffFlowDo
	Not(conds ...gen.Condition) IOvertimeSignOffFlowDo
	Or(conds ...gen.Condition) IOvertimeSignOffFlowDo
	Select(conds ...field.Expr) IOvertimeSignOffFlowDo
	Where(conds ...gen.Condition) IOvertimeSignOffFlowDo
	Order(conds ...field.Expr) IOvertimeSignOffFlowDo
	Distinct(cols ...field.Expr) IOvertimeSignOffFlowDo
	Omit(cols ...field.Expr) IOvertimeSignOffFlowDo
	Join(table schema.Tabler, on ...field.Expr) IOvertimeSignOffFlowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOvertimeSignOffFlowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOvertimeSignOffFlowDo
	Group(cols ...field.Expr) IOvertimeSignOffFlowDo
	Having(conds ...gen.Condition) IOvertimeSignOffFlowDo
	Limit(limit int) IOvertimeSignOffFlowDo
	Offset(offset int) IOvertimeSignOffFlowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOvertimeSignOffFlowDo
	Unscoped() IOvertimeSignOffFlowDo
	Create(values ...*types.OvertimeSignOffFlow) error
	CreateInBatches(values []*types.OvertimeSignOffFlow, batchSize int) error
	Save(values ...*types.OvertimeSignOffFlow) error
	First() (*types.OvertimeSignOffFlow, error)
	Take() (*types.OvertimeSignOffFlow, error)
	Last() (*types.OvertimeSignOffFlow, error)
	Find() ([]*types.OvertimeSignOffFlow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.OvertimeSignOffFlow, err error)
	FindInBatches(result *[]*types.OvertimeSignOffFlow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.OvertimeSignOffFlow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOvertimeSignOffFlowDo
	Assign(attrs ...field.AssignExpr) IOvertimeSignOffFlowDo
	Joins(fields ...field.RelationField) IOvertimeSignOffFlowDo
	Preload(fields ...field.RelationField) IOvertimeSignOffFlowDo
	FirstOrInit() (*types.OvertimeSignOffFlow, error)
	FirstOrCreate() (*types.OvertimeSignOffFlow, error)
	FindByPage(offset int, limit int) (result []*types.OvertimeSignOffFlow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOvertimeSignOffFlowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o overtimeSignOffFlowDo) Debug() IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Debug())
}

func (o overtimeSignOffFlowDo) WithContext(ctx context.Context) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o overtimeSignOffFlowDo) ReadDB() IOvertimeSignOffFlowDo {
	return o.Clauses(dbresolver.Read)
}

func (o overtimeSignOffFlowDo) WriteDB() IOvertimeSignOffFlowDo {
	return o.Clauses(dbresolver.Write)
}

func (o overtimeSignOffFlowDo) Session(config *gorm.Session) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Session(config))
}

func (o overtimeSignOffFlowDo) Clauses(conds ...clause.Expression) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o overtimeSignOffFlowDo) Returning(value interface{}, columns ...string) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o overtimeSignOffFlowDo) Not(conds ...gen.Condition) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o overtimeSignOffFlowDo) Or(conds ...gen.Condition) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o overtimeSignOffFlowDo) Select(conds ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o overtimeSignOffFlowDo) Where(conds ...gen.Condition) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o overtimeSignOffFlowDo) Order(conds ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o overtimeSignOffFlowDo) Distinct(cols ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o overtimeSignOffFlowDo) Omit(cols ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o overtimeSignOffFlowDo) Join(table schema.Tabler, on ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o overtimeSignOffFlowDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o overtimeSignOffFlowDo) RightJoin(table schema.Tabler, on ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o overtimeSignOffFlowDo) Group(cols ...field.Expr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o overtimeSignOffFlowDo) Having(conds ...gen.Condition) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o overtimeSignOffFlowDo) Limit(limit int) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o overtimeSignOffFlowDo) Offset(offset int) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o overtimeSignOffFlowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o overtimeSignOffFlowDo) Unscoped() IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Unscoped())
}

func (o overtimeSignOffFlowDo) Create(values ...*types.OvertimeSignOffFlow) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o overtimeSignOffFlowDo) CreateInBatches(values []*types.OvertimeSignOffFlow, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o overtimeSignOffFlowDo) Save(values ...*types.OvertimeSignOffFlow) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o overtimeSignOffFlowDo) First() (*types.OvertimeSignOffFlow, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeSignOffFlow), nil
	}
}

func (o overtimeSignOffFlowDo) Take() (*types.OvertimeSignOffFlow, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeSignOffFlow), nil
	}
}

func (o overtimeSignOffFlowDo) Last() (*types.OvertimeSignOffFlow, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeSignOffFlow), nil
	}
}

func (o overtimeSignOffFlowDo) Find() ([]*types.OvertimeSignOffFlow, error) {
	result, err := o.DO.Find()
	return result.([]*types.OvertimeSignOffFlow), err
}

func (o overtimeSignOffFlowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.OvertimeSignOffFlow, err error) {
	buf := make([]*types.OvertimeSignOffFlow, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o overtimeSignOffFlowDo) FindInBatches(result *[]*types.OvertimeSignOffFlow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o overtimeSignOffFlowDo) Attrs(attrs ...field.AssignExpr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o overtimeSignOffFlowDo) Assign(attrs ...field.AssignExpr) IOvertimeSignOffFlowDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o overtimeSignOffFlowDo) Joins(fields ...field.RelationField) IOvertimeSignOffFlowDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o overtimeSignOffFlowDo) Preload(fields ...field.RelationField) IOvertimeSignOffFlowDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o overtimeSignOffFlowDo) FirstOrInit() (*types.OvertimeSignOffFlow, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeSignOffFlow), nil
	}
}

func (o overtimeSignOffFlowDo) FirstOrCreate() (*types.OvertimeSignOffFlow, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.OvertimeSignOffFlow), nil
	}
}

func (o overtimeSignOffFlowDo) FindByPage(offset int, limit int) (result []*types.OvertimeSignOffFlow, count int64, err error) {
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

func (o overtimeSignOffFlowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o overtimeSignOffFlowDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o overtimeSignOffFlowDo) Delete(models ...*types.OvertimeSignOffFlow) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *overtimeSignOffFlowDo) withDO(do gen.Dao) *overtimeSignOffFlowDo {
	o.DO = *do.(*gen.DO)
	return o
}
