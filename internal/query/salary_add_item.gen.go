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

func newSalaryAddItem(db *gorm.DB, opts ...gen.DOOption) salaryAddItem {
	_salaryAddItem := salaryAddItem{}

	_salaryAddItem.salaryAddItemDo.UseDB(db, opts...)
	_salaryAddItem.salaryAddItemDo.UseModel(&types.SalaryAddItem{})

	tableName := _salaryAddItem.salaryAddItemDo.TableName()
	_salaryAddItem.ALL = field.NewAsterisk(tableName)
	_salaryAddItem.ID = field.NewUint(tableName, "id")
	_salaryAddItem.CreatedAt = field.NewTime(tableName, "created_at")
	_salaryAddItem.UpdatedAt = field.NewTime(tableName, "updated_at")
	_salaryAddItem.DeletedAt = field.NewField(tableName, "deleted_at")
	_salaryAddItem.Name = field.NewString(tableName, "name")
	_salaryAddItem.SalaryType = field.NewField(tableName, "salary_type")
	_salaryAddItem.IncomeTax = field.NewBool(tableName, "income_tax")
	_salaryAddItem.Benefits = field.NewBool(tableName, "benefits")
	_salaryAddItem.Premiums = field.NewBool(tableName, "premiums")
	_salaryAddItem.Amount = field.NewFloat32(tableName, "amount")
	_salaryAddItem.CalcType = field.NewField(tableName, "calc_type")
	_salaryAddItem.MonthCalc = field.NewBool(tableName, "month_calc")
	_salaryAddItem.Unit = field.NewField(tableName, "unit")
	_salaryAddItem.Operator = field.NewField(tableName, "operator")
	_salaryAddItem.Argument = field.NewFloat32(tableName, "argument")
	_salaryAddItem.Employee = salaryAddItemManyToManyEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Employee", "types.Employee"),
		Department: struct {
			field.RelationField
			Manager struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Employee.Department", "types.Department"),
			Manager: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Employee.Department.Manager", "types.Employee"),
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
			RelationField: field.NewRelation("Employee.Rank", "types.PositionRank"),
			Grade: struct {
				field.RelationField
				Rank struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Employee.Rank.Grade", "types.PositionGrade"),
				Rank: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Employee.Rank.Grade.Rank", "types.PositionRank"),
				},
			},
		},
		Grade: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Employee.Grade", "types.PositionGrade"),
		},
		LoginInformation: struct {
			field.RelationField
			Employee struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Employee.LoginInformation", "types.LoginInformation"),
			Employee: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Employee.LoginInformation.Employee", "types.Employee"),
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
			RelationField: field.NewRelation("Employee.Roles", "types.Role"),
			Employees: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Employee.Roles.Employees", "types.Employee"),
			},
			Permissions: struct {
				field.RelationField
				Roles struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Employee.Roles.Permissions", "types.Permission"),
				Roles: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Employee.Roles.Permissions.Roles", "types.Role"),
				},
			},
			Menus: struct {
				field.RelationField
				Roles struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Employee.Roles.Menus", "types.Menu"),
				Roles: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Employee.Roles.Menus.Roles", "types.Role"),
				},
			},
		},
	}

	_salaryAddItem.fillFieldMap()

	return _salaryAddItem
}

type salaryAddItem struct {
	salaryAddItemDo salaryAddItemDo

	ALL        field.Asterisk
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Name       field.String
	SalaryType field.Field
	IncomeTax  field.Bool
	Benefits   field.Bool
	Premiums   field.Bool
	Amount     field.Float32
	CalcType   field.Field
	MonthCalc  field.Bool
	Unit       field.Field
	Operator   field.Field
	Argument   field.Float32
	Employee   salaryAddItemManyToManyEmployee

	fieldMap map[string]field.Expr
}

func (s salaryAddItem) Table(newTableName string) *salaryAddItem {
	s.salaryAddItemDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s salaryAddItem) As(alias string) *salaryAddItem {
	s.salaryAddItemDo.DO = *(s.salaryAddItemDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *salaryAddItem) updateTableName(table string) *salaryAddItem {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Name = field.NewString(table, "name")
	s.SalaryType = field.NewField(table, "salary_type")
	s.IncomeTax = field.NewBool(table, "income_tax")
	s.Benefits = field.NewBool(table, "benefits")
	s.Premiums = field.NewBool(table, "premiums")
	s.Amount = field.NewFloat32(table, "amount")
	s.CalcType = field.NewField(table, "calc_type")
	s.MonthCalc = field.NewBool(table, "month_calc")
	s.Unit = field.NewField(table, "unit")
	s.Operator = field.NewField(table, "operator")
	s.Argument = field.NewFloat32(table, "argument")

	s.fillFieldMap()

	return s
}

func (s *salaryAddItem) WithContext(ctx context.Context) ISalaryAddItemDo {
	return s.salaryAddItemDo.WithContext(ctx)
}

func (s salaryAddItem) TableName() string { return s.salaryAddItemDo.TableName() }

func (s salaryAddItem) Alias() string { return s.salaryAddItemDo.Alias() }

func (s salaryAddItem) Columns(cols ...field.Expr) gen.Columns {
	return s.salaryAddItemDo.Columns(cols...)
}

func (s *salaryAddItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *salaryAddItem) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["salary_type"] = s.SalaryType
	s.fieldMap["income_tax"] = s.IncomeTax
	s.fieldMap["benefits"] = s.Benefits
	s.fieldMap["premiums"] = s.Premiums
	s.fieldMap["amount"] = s.Amount
	s.fieldMap["calc_type"] = s.CalcType
	s.fieldMap["month_calc"] = s.MonthCalc
	s.fieldMap["unit"] = s.Unit
	s.fieldMap["operator"] = s.Operator
	s.fieldMap["argument"] = s.Argument

}

func (s salaryAddItem) clone(db *gorm.DB) salaryAddItem {
	s.salaryAddItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s salaryAddItem) replaceDB(db *gorm.DB) salaryAddItem {
	s.salaryAddItemDo.ReplaceDB(db)
	return s
}

type salaryAddItemManyToManyEmployee struct {
	db *gorm.DB

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

func (a salaryAddItemManyToManyEmployee) Where(conds ...field.Expr) *salaryAddItemManyToManyEmployee {
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

func (a salaryAddItemManyToManyEmployee) WithContext(ctx context.Context) *salaryAddItemManyToManyEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a salaryAddItemManyToManyEmployee) Session(session *gorm.Session) *salaryAddItemManyToManyEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a salaryAddItemManyToManyEmployee) Model(m *types.SalaryAddItem) *salaryAddItemManyToManyEmployeeTx {
	return &salaryAddItemManyToManyEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type salaryAddItemManyToManyEmployeeTx struct{ tx *gorm.Association }

func (a salaryAddItemManyToManyEmployeeTx) Find() (result []*types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a salaryAddItemManyToManyEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a salaryAddItemManyToManyEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a salaryAddItemManyToManyEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a salaryAddItemManyToManyEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a salaryAddItemManyToManyEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type salaryAddItemDo struct{ gen.DO }

type ISalaryAddItemDo interface {
	gen.SubQuery
	Debug() ISalaryAddItemDo
	WithContext(ctx context.Context) ISalaryAddItemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISalaryAddItemDo
	WriteDB() ISalaryAddItemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISalaryAddItemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISalaryAddItemDo
	Not(conds ...gen.Condition) ISalaryAddItemDo
	Or(conds ...gen.Condition) ISalaryAddItemDo
	Select(conds ...field.Expr) ISalaryAddItemDo
	Where(conds ...gen.Condition) ISalaryAddItemDo
	Order(conds ...field.Expr) ISalaryAddItemDo
	Distinct(cols ...field.Expr) ISalaryAddItemDo
	Omit(cols ...field.Expr) ISalaryAddItemDo
	Join(table schema.Tabler, on ...field.Expr) ISalaryAddItemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISalaryAddItemDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISalaryAddItemDo
	Group(cols ...field.Expr) ISalaryAddItemDo
	Having(conds ...gen.Condition) ISalaryAddItemDo
	Limit(limit int) ISalaryAddItemDo
	Offset(offset int) ISalaryAddItemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISalaryAddItemDo
	Unscoped() ISalaryAddItemDo
	Create(values ...*types.SalaryAddItem) error
	CreateInBatches(values []*types.SalaryAddItem, batchSize int) error
	Save(values ...*types.SalaryAddItem) error
	First() (*types.SalaryAddItem, error)
	Take() (*types.SalaryAddItem, error)
	Last() (*types.SalaryAddItem, error)
	Find() ([]*types.SalaryAddItem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.SalaryAddItem, err error)
	FindInBatches(result *[]*types.SalaryAddItem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.SalaryAddItem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISalaryAddItemDo
	Assign(attrs ...field.AssignExpr) ISalaryAddItemDo
	Joins(fields ...field.RelationField) ISalaryAddItemDo
	Preload(fields ...field.RelationField) ISalaryAddItemDo
	FirstOrInit() (*types.SalaryAddItem, error)
	FirstOrCreate() (*types.SalaryAddItem, error)
	FindByPage(offset int, limit int) (result []*types.SalaryAddItem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISalaryAddItemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	QueryByEmployeeIDMultiply(empIDs string, dateOnly1 string, dateOnly2 string) (result []map[string]interface{}, err error)
	QueryByEmployeeID(empID uint, keyword string) (result []map[string]interface{}, err error)
}

// select * from FN_C_SalaryAddItemMultiply (@empIDs, @dateOnly1, @dateOnly2)
func (s salaryAddItemDo) QueryByEmployeeIDMultiply(empIDs string, dateOnly1 string, dateOnly2 string) (result []map[string]interface{}, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, empIDs)
	params = append(params, dateOnly1)
	params = append(params, dateOnly2)
	generateSQL.WriteString("select * from FN_C_SalaryAddItemMultiply (?, ?, ?) ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select sa.*, sae.use_custom, sae.custom_amount from salary_add_item_employee sae
//
//	join salary_add_item sa on (sa.id = sae.salary_add_item_id)
//	where sae.employee_id = @empID and sa.name like @keyword
func (s salaryAddItemDo) QueryByEmployeeID(empID uint, keyword string) (result []map[string]interface{}, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, empID)
	params = append(params, keyword)
	generateSQL.WriteString("select sa.*, sae.use_custom, sae.custom_amount from salary_add_item_employee sae join salary_add_item sa on (sa.id = sae.salary_add_item_id) where sae.employee_id = ? and sa.name like ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s salaryAddItemDo) Debug() ISalaryAddItemDo {
	return s.withDO(s.DO.Debug())
}

func (s salaryAddItemDo) WithContext(ctx context.Context) ISalaryAddItemDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s salaryAddItemDo) ReadDB() ISalaryAddItemDo {
	return s.Clauses(dbresolver.Read)
}

func (s salaryAddItemDo) WriteDB() ISalaryAddItemDo {
	return s.Clauses(dbresolver.Write)
}

func (s salaryAddItemDo) Session(config *gorm.Session) ISalaryAddItemDo {
	return s.withDO(s.DO.Session(config))
}

func (s salaryAddItemDo) Clauses(conds ...clause.Expression) ISalaryAddItemDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s salaryAddItemDo) Returning(value interface{}, columns ...string) ISalaryAddItemDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s salaryAddItemDo) Not(conds ...gen.Condition) ISalaryAddItemDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s salaryAddItemDo) Or(conds ...gen.Condition) ISalaryAddItemDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s salaryAddItemDo) Select(conds ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s salaryAddItemDo) Where(conds ...gen.Condition) ISalaryAddItemDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s salaryAddItemDo) Order(conds ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s salaryAddItemDo) Distinct(cols ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s salaryAddItemDo) Omit(cols ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s salaryAddItemDo) Join(table schema.Tabler, on ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s salaryAddItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s salaryAddItemDo) RightJoin(table schema.Tabler, on ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s salaryAddItemDo) Group(cols ...field.Expr) ISalaryAddItemDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s salaryAddItemDo) Having(conds ...gen.Condition) ISalaryAddItemDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s salaryAddItemDo) Limit(limit int) ISalaryAddItemDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s salaryAddItemDo) Offset(offset int) ISalaryAddItemDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s salaryAddItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISalaryAddItemDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s salaryAddItemDo) Unscoped() ISalaryAddItemDo {
	return s.withDO(s.DO.Unscoped())
}

func (s salaryAddItemDo) Create(values ...*types.SalaryAddItem) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s salaryAddItemDo) CreateInBatches(values []*types.SalaryAddItem, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s salaryAddItemDo) Save(values ...*types.SalaryAddItem) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s salaryAddItemDo) First() (*types.SalaryAddItem, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.SalaryAddItem), nil
	}
}

func (s salaryAddItemDo) Take() (*types.SalaryAddItem, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.SalaryAddItem), nil
	}
}

func (s salaryAddItemDo) Last() (*types.SalaryAddItem, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.SalaryAddItem), nil
	}
}

func (s salaryAddItemDo) Find() ([]*types.SalaryAddItem, error) {
	result, err := s.DO.Find()
	return result.([]*types.SalaryAddItem), err
}

func (s salaryAddItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.SalaryAddItem, err error) {
	buf := make([]*types.SalaryAddItem, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s salaryAddItemDo) FindInBatches(result *[]*types.SalaryAddItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s salaryAddItemDo) Attrs(attrs ...field.AssignExpr) ISalaryAddItemDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s salaryAddItemDo) Assign(attrs ...field.AssignExpr) ISalaryAddItemDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s salaryAddItemDo) Joins(fields ...field.RelationField) ISalaryAddItemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s salaryAddItemDo) Preload(fields ...field.RelationField) ISalaryAddItemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s salaryAddItemDo) FirstOrInit() (*types.SalaryAddItem, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.SalaryAddItem), nil
	}
}

func (s salaryAddItemDo) FirstOrCreate() (*types.SalaryAddItem, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.SalaryAddItem), nil
	}
}

func (s salaryAddItemDo) FindByPage(offset int, limit int) (result []*types.SalaryAddItem, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s salaryAddItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s salaryAddItemDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s salaryAddItemDo) Delete(models ...*types.SalaryAddItem) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *salaryAddItemDo) withDO(do gen.Dao) *salaryAddItemDo {
	s.DO = *do.(*gen.DO)
	return s
}
