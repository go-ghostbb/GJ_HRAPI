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

func newLoginInformation(db *gorm.DB, opts ...gen.DOOption) loginInformation {
	_loginInformation := loginInformation{}

	_loginInformation.loginInformationDo.UseDB(db, opts...)
	_loginInformation.loginInformationDo.UseModel(&types.LoginInformation{})

	tableName := _loginInformation.loginInformationDo.TableName()
	_loginInformation.ALL = field.NewAsterisk(tableName)
	_loginInformation.ID = field.NewUint(tableName, "id")
	_loginInformation.CreatedAt = field.NewTime(tableName, "created_at")
	_loginInformation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_loginInformation.DeletedAt = field.NewField(tableName, "deleted_at")
	_loginInformation.EmployeeID = field.NewUint(tableName, "employee_id")
	_loginInformation.Status = field.NewBool(tableName, "status")
	_loginInformation.Account = field.NewString(tableName, "account")
	_loginInformation.Password = field.NewString(tableName, "password")
	_loginInformation.Employee = loginInformationBelongsToEmployee{
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

	_loginInformation.fillFieldMap()

	return _loginInformation
}

type loginInformation struct {
	loginInformationDo loginInformationDo

	ALL        field.Asterisk
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	EmployeeID field.Uint
	Status     field.Bool
	Account    field.String
	Password   field.String
	Employee   loginInformationBelongsToEmployee

	fieldMap map[string]field.Expr
}

func (l loginInformation) Table(newTableName string) *loginInformation {
	l.loginInformationDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l loginInformation) As(alias string) *loginInformation {
	l.loginInformationDo.DO = *(l.loginInformationDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *loginInformation) updateTableName(table string) *loginInformation {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewUint(table, "id")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.DeletedAt = field.NewField(table, "deleted_at")
	l.EmployeeID = field.NewUint(table, "employee_id")
	l.Status = field.NewBool(table, "status")
	l.Account = field.NewString(table, "account")
	l.Password = field.NewString(table, "password")

	l.fillFieldMap()

	return l
}

func (l *loginInformation) WithContext(ctx context.Context) ILoginInformationDo {
	return l.loginInformationDo.WithContext(ctx)
}

func (l loginInformation) TableName() string { return l.loginInformationDo.TableName() }

func (l loginInformation) Alias() string { return l.loginInformationDo.Alias() }

func (l loginInformation) Columns(cols ...field.Expr) gen.Columns {
	return l.loginInformationDo.Columns(cols...)
}

func (l *loginInformation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *loginInformation) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 9)
	l.fieldMap["id"] = l.ID
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["deleted_at"] = l.DeletedAt
	l.fieldMap["employee_id"] = l.EmployeeID
	l.fieldMap["status"] = l.Status
	l.fieldMap["account"] = l.Account
	l.fieldMap["password"] = l.Password

}

func (l loginInformation) clone(db *gorm.DB) loginInformation {
	l.loginInformationDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l loginInformation) replaceDB(db *gorm.DB) loginInformation {
	l.loginInformationDo.ReplaceDB(db)
	return l
}

type loginInformationBelongsToEmployee struct {
	db *gorm.DB

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

func (a loginInformationBelongsToEmployee) Where(conds ...field.Expr) *loginInformationBelongsToEmployee {
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

func (a loginInformationBelongsToEmployee) WithContext(ctx context.Context) *loginInformationBelongsToEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a loginInformationBelongsToEmployee) Session(session *gorm.Session) *loginInformationBelongsToEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a loginInformationBelongsToEmployee) Model(m *types.LoginInformation) *loginInformationBelongsToEmployeeTx {
	return &loginInformationBelongsToEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type loginInformationBelongsToEmployeeTx struct{ tx *gorm.Association }

func (a loginInformationBelongsToEmployeeTx) Find() (result *types.Employee, err error) {
	return result, a.tx.Find(&result)
}

func (a loginInformationBelongsToEmployeeTx) Append(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a loginInformationBelongsToEmployeeTx) Replace(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a loginInformationBelongsToEmployeeTx) Delete(values ...*types.Employee) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a loginInformationBelongsToEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a loginInformationBelongsToEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type loginInformationDo struct{ gen.DO }

type ILoginInformationDo interface {
	gen.SubQuery
	Debug() ILoginInformationDo
	WithContext(ctx context.Context) ILoginInformationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILoginInformationDo
	WriteDB() ILoginInformationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILoginInformationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILoginInformationDo
	Not(conds ...gen.Condition) ILoginInformationDo
	Or(conds ...gen.Condition) ILoginInformationDo
	Select(conds ...field.Expr) ILoginInformationDo
	Where(conds ...gen.Condition) ILoginInformationDo
	Order(conds ...field.Expr) ILoginInformationDo
	Distinct(cols ...field.Expr) ILoginInformationDo
	Omit(cols ...field.Expr) ILoginInformationDo
	Join(table schema.Tabler, on ...field.Expr) ILoginInformationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILoginInformationDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILoginInformationDo
	Group(cols ...field.Expr) ILoginInformationDo
	Having(conds ...gen.Condition) ILoginInformationDo
	Limit(limit int) ILoginInformationDo
	Offset(offset int) ILoginInformationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILoginInformationDo
	Unscoped() ILoginInformationDo
	Create(values ...*types.LoginInformation) error
	CreateInBatches(values []*types.LoginInformation, batchSize int) error
	Save(values ...*types.LoginInformation) error
	First() (*types.LoginInformation, error)
	Take() (*types.LoginInformation, error)
	Last() (*types.LoginInformation, error)
	Find() ([]*types.LoginInformation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.LoginInformation, err error)
	FindInBatches(result *[]*types.LoginInformation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.LoginInformation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILoginInformationDo
	Assign(attrs ...field.AssignExpr) ILoginInformationDo
	Joins(fields ...field.RelationField) ILoginInformationDo
	Preload(fields ...field.RelationField) ILoginInformationDo
	FirstOrInit() (*types.LoginInformation, error)
	FirstOrCreate() (*types.LoginInformation, error)
	FindByPage(offset int, limit int) (result []*types.LoginInformation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILoginInformationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l loginInformationDo) Debug() ILoginInformationDo {
	return l.withDO(l.DO.Debug())
}

func (l loginInformationDo) WithContext(ctx context.Context) ILoginInformationDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l loginInformationDo) ReadDB() ILoginInformationDo {
	return l.Clauses(dbresolver.Read)
}

func (l loginInformationDo) WriteDB() ILoginInformationDo {
	return l.Clauses(dbresolver.Write)
}

func (l loginInformationDo) Session(config *gorm.Session) ILoginInformationDo {
	return l.withDO(l.DO.Session(config))
}

func (l loginInformationDo) Clauses(conds ...clause.Expression) ILoginInformationDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l loginInformationDo) Returning(value interface{}, columns ...string) ILoginInformationDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l loginInformationDo) Not(conds ...gen.Condition) ILoginInformationDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l loginInformationDo) Or(conds ...gen.Condition) ILoginInformationDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l loginInformationDo) Select(conds ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l loginInformationDo) Where(conds ...gen.Condition) ILoginInformationDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l loginInformationDo) Order(conds ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l loginInformationDo) Distinct(cols ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l loginInformationDo) Omit(cols ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l loginInformationDo) Join(table schema.Tabler, on ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l loginInformationDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l loginInformationDo) RightJoin(table schema.Tabler, on ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l loginInformationDo) Group(cols ...field.Expr) ILoginInformationDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l loginInformationDo) Having(conds ...gen.Condition) ILoginInformationDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l loginInformationDo) Limit(limit int) ILoginInformationDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l loginInformationDo) Offset(offset int) ILoginInformationDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l loginInformationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILoginInformationDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l loginInformationDo) Unscoped() ILoginInformationDo {
	return l.withDO(l.DO.Unscoped())
}

func (l loginInformationDo) Create(values ...*types.LoginInformation) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l loginInformationDo) CreateInBatches(values []*types.LoginInformation, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l loginInformationDo) Save(values ...*types.LoginInformation) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l loginInformationDo) First() (*types.LoginInformation, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.LoginInformation), nil
	}
}

func (l loginInformationDo) Take() (*types.LoginInformation, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.LoginInformation), nil
	}
}

func (l loginInformationDo) Last() (*types.LoginInformation, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.LoginInformation), nil
	}
}

func (l loginInformationDo) Find() ([]*types.LoginInformation, error) {
	result, err := l.DO.Find()
	return result.([]*types.LoginInformation), err
}

func (l loginInformationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.LoginInformation, err error) {
	buf := make([]*types.LoginInformation, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l loginInformationDo) FindInBatches(result *[]*types.LoginInformation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l loginInformationDo) Attrs(attrs ...field.AssignExpr) ILoginInformationDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l loginInformationDo) Assign(attrs ...field.AssignExpr) ILoginInformationDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l loginInformationDo) Joins(fields ...field.RelationField) ILoginInformationDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l loginInformationDo) Preload(fields ...field.RelationField) ILoginInformationDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l loginInformationDo) FirstOrInit() (*types.LoginInformation, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.LoginInformation), nil
	}
}

func (l loginInformationDo) FirstOrCreate() (*types.LoginInformation, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.LoginInformation), nil
	}
}

func (l loginInformationDo) FindByPage(offset int, limit int) (result []*types.LoginInformation, count int64, err error) {
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

func (l loginInformationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l loginInformationDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l loginInformationDo) Delete(models ...*types.LoginInformation) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *loginInformationDo) withDO(do gen.Dao) *loginInformationDo {
	l.DO = *do.(*gen.DO)
	return l
}
