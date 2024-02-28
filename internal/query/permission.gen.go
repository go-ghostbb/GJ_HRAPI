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

func newPermission(db *gorm.DB, opts ...gen.DOOption) permission {
	_permission := permission{}

	_permission.permissionDo.UseDB(db, opts...)
	_permission.permissionDo.UseModel(&types.Permission{})

	tableName := _permission.permissionDo.TableName()
	_permission.ALL = field.NewAsterisk(tableName)
	_permission.ID = field.NewUint(tableName, "id")
	_permission.CreatedAt = field.NewTime(tableName, "created_at")
	_permission.UpdatedAt = field.NewTime(tableName, "updated_at")
	_permission.DeletedAt = field.NewField(tableName, "deleted_at")
	_permission.Perm = field.NewString(tableName, "perm")
	_permission.Name = field.NewString(tableName, "name")
	_permission.Status = field.NewBool(tableName, "status")
	_permission.Remark = field.NewString(tableName, "remark")
	_permission.Roles = permissionManyToManyRoles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Roles", "types.Role"),
		Employees: struct {
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
			}
		}{
			RelationField: field.NewRelation("Roles.Employees", "types.Employee"),
			Department: struct {
				field.RelationField
				Manager struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Roles.Employees.Department", "types.Department"),
				Manager: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Roles.Employees.Department.Manager", "types.Employee"),
				},
			},
			LoginInformation: struct {
				field.RelationField
				Employee struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Roles.Employees.LoginInformation", "types.LoginInformation"),
				Employee: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Roles.Employees.LoginInformation.Employee", "types.Employee"),
				},
			},
			Roles: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Roles.Employees.Roles", "types.Role"),
			},
		},
		Permissions: struct {
			field.RelationField
			Roles struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Roles.Permissions", "types.Permission"),
			Roles: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Roles.Permissions.Roles", "types.Role"),
			},
		},
		Menus: struct {
			field.RelationField
			Roles struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Roles.Menus", "types.Menu"),
			Roles: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Roles.Menus.Roles", "types.Role"),
			},
		},
	}

	_permission.fillFieldMap()

	return _permission
}

type permission struct {
	permissionDo permissionDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Perm      field.String
	Name      field.String
	Status    field.Bool
	Remark    field.String
	Roles     permissionManyToManyRoles

	fieldMap map[string]field.Expr
}

func (p permission) Table(newTableName string) *permission {
	p.permissionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p permission) As(alias string) *permission {
	p.permissionDo.DO = *(p.permissionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *permission) updateTableName(table string) *permission {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Perm = field.NewString(table, "perm")
	p.Name = field.NewString(table, "name")
	p.Status = field.NewBool(table, "status")
	p.Remark = field.NewString(table, "remark")

	p.fillFieldMap()

	return p
}

func (p *permission) WithContext(ctx context.Context) IPermissionDo {
	return p.permissionDo.WithContext(ctx)
}

func (p permission) TableName() string { return p.permissionDo.TableName() }

func (p permission) Alias() string { return p.permissionDo.Alias() }

func (p permission) Columns(cols ...field.Expr) gen.Columns { return p.permissionDo.Columns(cols...) }

func (p *permission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *permission) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["perm"] = p.Perm
	p.fieldMap["name"] = p.Name
	p.fieldMap["status"] = p.Status
	p.fieldMap["remark"] = p.Remark

}

func (p permission) clone(db *gorm.DB) permission {
	p.permissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p permission) replaceDB(db *gorm.DB) permission {
	p.permissionDo.ReplaceDB(db)
	return p
}

type permissionManyToManyRoles struct {
	db *gorm.DB

	field.RelationField

	Employees struct {
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
		}
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

func (a permissionManyToManyRoles) Where(conds ...field.Expr) *permissionManyToManyRoles {
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

func (a permissionManyToManyRoles) WithContext(ctx context.Context) *permissionManyToManyRoles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a permissionManyToManyRoles) Session(session *gorm.Session) *permissionManyToManyRoles {
	a.db = a.db.Session(session)
	return &a
}

func (a permissionManyToManyRoles) Model(m *types.Permission) *permissionManyToManyRolesTx {
	return &permissionManyToManyRolesTx{a.db.Model(m).Association(a.Name())}
}

type permissionManyToManyRolesTx struct{ tx *gorm.Association }

func (a permissionManyToManyRolesTx) Find() (result []*types.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a permissionManyToManyRolesTx) Append(values ...*types.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a permissionManyToManyRolesTx) Replace(values ...*types.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a permissionManyToManyRolesTx) Delete(values ...*types.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a permissionManyToManyRolesTx) Clear() error {
	return a.tx.Clear()
}

func (a permissionManyToManyRolesTx) Count() int64 {
	return a.tx.Count()
}

type permissionDo struct{ gen.DO }

type IPermissionDo interface {
	gen.SubQuery
	Debug() IPermissionDo
	WithContext(ctx context.Context) IPermissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPermissionDo
	WriteDB() IPermissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPermissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPermissionDo
	Not(conds ...gen.Condition) IPermissionDo
	Or(conds ...gen.Condition) IPermissionDo
	Select(conds ...field.Expr) IPermissionDo
	Where(conds ...gen.Condition) IPermissionDo
	Order(conds ...field.Expr) IPermissionDo
	Distinct(cols ...field.Expr) IPermissionDo
	Omit(cols ...field.Expr) IPermissionDo
	Join(table schema.Tabler, on ...field.Expr) IPermissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPermissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPermissionDo
	Group(cols ...field.Expr) IPermissionDo
	Having(conds ...gen.Condition) IPermissionDo
	Limit(limit int) IPermissionDo
	Offset(offset int) IPermissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPermissionDo
	Unscoped() IPermissionDo
	Create(values ...*types.Permission) error
	CreateInBatches(values []*types.Permission, batchSize int) error
	Save(values ...*types.Permission) error
	First() (*types.Permission, error)
	Take() (*types.Permission, error)
	Last() (*types.Permission, error)
	Find() ([]*types.Permission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.Permission, err error)
	FindInBatches(result *[]*types.Permission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.Permission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPermissionDo
	Assign(attrs ...field.AssignExpr) IPermissionDo
	Joins(fields ...field.RelationField) IPermissionDo
	Preload(fields ...field.RelationField) IPermissionDo
	FirstOrInit() (*types.Permission, error)
	FirstOrCreate() (*types.Permission, error)
	FindByPage(offset int, limit int) (result []*types.Permission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPermissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p permissionDo) Debug() IPermissionDo {
	return p.withDO(p.DO.Debug())
}

func (p permissionDo) WithContext(ctx context.Context) IPermissionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p permissionDo) ReadDB() IPermissionDo {
	return p.Clauses(dbresolver.Read)
}

func (p permissionDo) WriteDB() IPermissionDo {
	return p.Clauses(dbresolver.Write)
}

func (p permissionDo) Session(config *gorm.Session) IPermissionDo {
	return p.withDO(p.DO.Session(config))
}

func (p permissionDo) Clauses(conds ...clause.Expression) IPermissionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p permissionDo) Returning(value interface{}, columns ...string) IPermissionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p permissionDo) Not(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p permissionDo) Or(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p permissionDo) Select(conds ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p permissionDo) Where(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p permissionDo) Order(conds ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p permissionDo) Distinct(cols ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p permissionDo) Omit(cols ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p permissionDo) Join(table schema.Tabler, on ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p permissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p permissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p permissionDo) Group(cols ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p permissionDo) Having(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p permissionDo) Limit(limit int) IPermissionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p permissionDo) Offset(offset int) IPermissionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p permissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPermissionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p permissionDo) Unscoped() IPermissionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p permissionDo) Create(values ...*types.Permission) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p permissionDo) CreateInBatches(values []*types.Permission, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p permissionDo) Save(values ...*types.Permission) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p permissionDo) First() (*types.Permission, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.Permission), nil
	}
}

func (p permissionDo) Take() (*types.Permission, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.Permission), nil
	}
}

func (p permissionDo) Last() (*types.Permission, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.Permission), nil
	}
}

func (p permissionDo) Find() ([]*types.Permission, error) {
	result, err := p.DO.Find()
	return result.([]*types.Permission), err
}

func (p permissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.Permission, err error) {
	buf := make([]*types.Permission, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p permissionDo) FindInBatches(result *[]*types.Permission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p permissionDo) Attrs(attrs ...field.AssignExpr) IPermissionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p permissionDo) Assign(attrs ...field.AssignExpr) IPermissionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p permissionDo) Joins(fields ...field.RelationField) IPermissionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p permissionDo) Preload(fields ...field.RelationField) IPermissionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p permissionDo) FirstOrInit() (*types.Permission, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.Permission), nil
	}
}

func (p permissionDo) FirstOrCreate() (*types.Permission, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.Permission), nil
	}
}

func (p permissionDo) FindByPage(offset int, limit int) (result []*types.Permission, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p permissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p permissionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p permissionDo) Delete(models ...*types.Permission) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *permissionDo) withDO(do gen.Dao) *permissionDo {
	p.DO = *do.(*gen.DO)
	return p
}
