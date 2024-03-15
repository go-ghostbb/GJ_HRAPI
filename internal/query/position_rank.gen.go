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

func newPositionRank(db *gorm.DB, opts ...gen.DOOption) positionRank {
	_positionRank := positionRank{}

	_positionRank.positionRankDo.UseDB(db, opts...)
	_positionRank.positionRankDo.UseModel(&types.PositionRank{})

	tableName := _positionRank.positionRankDo.TableName()
	_positionRank.ALL = field.NewAsterisk(tableName)
	_positionRank.ID = field.NewUint(tableName, "id")
	_positionRank.CreatedAt = field.NewTime(tableName, "created_at")
	_positionRank.UpdatedAt = field.NewTime(tableName, "updated_at")
	_positionRank.DeletedAt = field.NewField(tableName, "deleted_at")
	_positionRank.Code = field.NewString(tableName, "code")
	_positionRank.Name = field.NewString(tableName, "name")
	_positionRank.Remark = field.NewString(tableName, "remark")
	_positionRank.Grade = positionRankHasManyGrade{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Grade", "types.PositionGrade"),
		Rank: struct {
			field.RelationField
			Grade struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Grade.Rank", "types.PositionRank"),
			Grade: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Grade.Rank.Grade", "types.PositionGrade"),
			},
		},
	}

	_positionRank.fillFieldMap()

	return _positionRank
}

type positionRank struct {
	positionRankDo positionRankDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Code      field.String
	Name      field.String
	Remark    field.String
	Grade     positionRankHasManyGrade

	fieldMap map[string]field.Expr
}

func (p positionRank) Table(newTableName string) *positionRank {
	p.positionRankDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p positionRank) As(alias string) *positionRank {
	p.positionRankDo.DO = *(p.positionRankDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *positionRank) updateTableName(table string) *positionRank {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Code = field.NewString(table, "code")
	p.Name = field.NewString(table, "name")
	p.Remark = field.NewString(table, "remark")

	p.fillFieldMap()

	return p
}

func (p *positionRank) WithContext(ctx context.Context) IPositionRankDo {
	return p.positionRankDo.WithContext(ctx)
}

func (p positionRank) TableName() string { return p.positionRankDo.TableName() }

func (p positionRank) Alias() string { return p.positionRankDo.Alias() }

func (p positionRank) Columns(cols ...field.Expr) gen.Columns {
	return p.positionRankDo.Columns(cols...)
}

func (p *positionRank) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *positionRank) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["code"] = p.Code
	p.fieldMap["name"] = p.Name
	p.fieldMap["remark"] = p.Remark

}

func (p positionRank) clone(db *gorm.DB) positionRank {
	p.positionRankDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p positionRank) replaceDB(db *gorm.DB) positionRank {
	p.positionRankDo.ReplaceDB(db)
	return p
}

type positionRankHasManyGrade struct {
	db *gorm.DB

	field.RelationField

	Rank struct {
		field.RelationField
		Grade struct {
			field.RelationField
		}
	}
}

func (a positionRankHasManyGrade) Where(conds ...field.Expr) *positionRankHasManyGrade {
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

func (a positionRankHasManyGrade) WithContext(ctx context.Context) *positionRankHasManyGrade {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a positionRankHasManyGrade) Session(session *gorm.Session) *positionRankHasManyGrade {
	a.db = a.db.Session(session)
	return &a
}

func (a positionRankHasManyGrade) Model(m *types.PositionRank) *positionRankHasManyGradeTx {
	return &positionRankHasManyGradeTx{a.db.Model(m).Association(a.Name())}
}

type positionRankHasManyGradeTx struct{ tx *gorm.Association }

func (a positionRankHasManyGradeTx) Find() (result []*types.PositionGrade, err error) {
	return result, a.tx.Find(&result)
}

func (a positionRankHasManyGradeTx) Append(values ...*types.PositionGrade) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a positionRankHasManyGradeTx) Replace(values ...*types.PositionGrade) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a positionRankHasManyGradeTx) Delete(values ...*types.PositionGrade) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a positionRankHasManyGradeTx) Clear() error {
	return a.tx.Clear()
}

func (a positionRankHasManyGradeTx) Count() int64 {
	return a.tx.Count()
}

type positionRankDo struct{ gen.DO }

type IPositionRankDo interface {
	gen.SubQuery
	Debug() IPositionRankDo
	WithContext(ctx context.Context) IPositionRankDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPositionRankDo
	WriteDB() IPositionRankDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPositionRankDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPositionRankDo
	Not(conds ...gen.Condition) IPositionRankDo
	Or(conds ...gen.Condition) IPositionRankDo
	Select(conds ...field.Expr) IPositionRankDo
	Where(conds ...gen.Condition) IPositionRankDo
	Order(conds ...field.Expr) IPositionRankDo
	Distinct(cols ...field.Expr) IPositionRankDo
	Omit(cols ...field.Expr) IPositionRankDo
	Join(table schema.Tabler, on ...field.Expr) IPositionRankDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPositionRankDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPositionRankDo
	Group(cols ...field.Expr) IPositionRankDo
	Having(conds ...gen.Condition) IPositionRankDo
	Limit(limit int) IPositionRankDo
	Offset(offset int) IPositionRankDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPositionRankDo
	Unscoped() IPositionRankDo
	Create(values ...*types.PositionRank) error
	CreateInBatches(values []*types.PositionRank, batchSize int) error
	Save(values ...*types.PositionRank) error
	First() (*types.PositionRank, error)
	Take() (*types.PositionRank, error)
	Last() (*types.PositionRank, error)
	Find() ([]*types.PositionRank, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.PositionRank, err error)
	FindInBatches(result *[]*types.PositionRank, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.PositionRank) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPositionRankDo
	Assign(attrs ...field.AssignExpr) IPositionRankDo
	Joins(fields ...field.RelationField) IPositionRankDo
	Preload(fields ...field.RelationField) IPositionRankDo
	FirstOrInit() (*types.PositionRank, error)
	FirstOrCreate() (*types.PositionRank, error)
	FindByPage(offset int, limit int) (result []*types.PositionRank, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPositionRankDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p positionRankDo) Debug() IPositionRankDo {
	return p.withDO(p.DO.Debug())
}

func (p positionRankDo) WithContext(ctx context.Context) IPositionRankDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p positionRankDo) ReadDB() IPositionRankDo {
	return p.Clauses(dbresolver.Read)
}

func (p positionRankDo) WriteDB() IPositionRankDo {
	return p.Clauses(dbresolver.Write)
}

func (p positionRankDo) Session(config *gorm.Session) IPositionRankDo {
	return p.withDO(p.DO.Session(config))
}

func (p positionRankDo) Clauses(conds ...clause.Expression) IPositionRankDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p positionRankDo) Returning(value interface{}, columns ...string) IPositionRankDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p positionRankDo) Not(conds ...gen.Condition) IPositionRankDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p positionRankDo) Or(conds ...gen.Condition) IPositionRankDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p positionRankDo) Select(conds ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p positionRankDo) Where(conds ...gen.Condition) IPositionRankDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p positionRankDo) Order(conds ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p positionRankDo) Distinct(cols ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p positionRankDo) Omit(cols ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p positionRankDo) Join(table schema.Tabler, on ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p positionRankDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p positionRankDo) RightJoin(table schema.Tabler, on ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p positionRankDo) Group(cols ...field.Expr) IPositionRankDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p positionRankDo) Having(conds ...gen.Condition) IPositionRankDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p positionRankDo) Limit(limit int) IPositionRankDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p positionRankDo) Offset(offset int) IPositionRankDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p positionRankDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPositionRankDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p positionRankDo) Unscoped() IPositionRankDo {
	return p.withDO(p.DO.Unscoped())
}

func (p positionRankDo) Create(values ...*types.PositionRank) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p positionRankDo) CreateInBatches(values []*types.PositionRank, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p positionRankDo) Save(values ...*types.PositionRank) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p positionRankDo) First() (*types.PositionRank, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionRank), nil
	}
}

func (p positionRankDo) Take() (*types.PositionRank, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionRank), nil
	}
}

func (p positionRankDo) Last() (*types.PositionRank, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionRank), nil
	}
}

func (p positionRankDo) Find() ([]*types.PositionRank, error) {
	result, err := p.DO.Find()
	return result.([]*types.PositionRank), err
}

func (p positionRankDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.PositionRank, err error) {
	buf := make([]*types.PositionRank, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p positionRankDo) FindInBatches(result *[]*types.PositionRank, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p positionRankDo) Attrs(attrs ...field.AssignExpr) IPositionRankDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p positionRankDo) Assign(attrs ...field.AssignExpr) IPositionRankDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p positionRankDo) Joins(fields ...field.RelationField) IPositionRankDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p positionRankDo) Preload(fields ...field.RelationField) IPositionRankDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p positionRankDo) FirstOrInit() (*types.PositionRank, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionRank), nil
	}
}

func (p positionRankDo) FirstOrCreate() (*types.PositionRank, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionRank), nil
	}
}

func (p positionRankDo) FindByPage(offset int, limit int) (result []*types.PositionRank, count int64, err error) {
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

func (p positionRankDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p positionRankDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p positionRankDo) Delete(models ...*types.PositionRank) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *positionRankDo) withDO(do gen.Dao) *positionRankDo {
	p.DO = *do.(*gen.DO)
	return p
}
