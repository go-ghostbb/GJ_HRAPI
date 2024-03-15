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

func newPositionGrade(db *gorm.DB, opts ...gen.DOOption) positionGrade {
	_positionGrade := positionGrade{}

	_positionGrade.positionGradeDo.UseDB(db, opts...)
	_positionGrade.positionGradeDo.UseModel(&types.PositionGrade{})

	tableName := _positionGrade.positionGradeDo.TableName()
	_positionGrade.ALL = field.NewAsterisk(tableName)
	_positionGrade.ID = field.NewUint(tableName, "id")
	_positionGrade.CreatedAt = field.NewTime(tableName, "created_at")
	_positionGrade.UpdatedAt = field.NewTime(tableName, "updated_at")
	_positionGrade.DeletedAt = field.NewField(tableName, "deleted_at")
	_positionGrade.Code = field.NewString(tableName, "code")
	_positionGrade.Name = field.NewString(tableName, "name")
	_positionGrade.Remark = field.NewString(tableName, "remark")
	_positionGrade.RankID = field.NewUint(tableName, "rank_id")
	_positionGrade.Rank = positionGradeBelongsToRank{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Rank", "types.PositionRank"),
		Grade: struct {
			field.RelationField
			Rank struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Rank.Grade", "types.PositionGrade"),
			Rank: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Rank.Grade.Rank", "types.PositionRank"),
			},
		},
	}

	_positionGrade.fillFieldMap()

	return _positionGrade
}

type positionGrade struct {
	positionGradeDo positionGradeDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Code      field.String
	Name      field.String
	Remark    field.String
	RankID    field.Uint
	Rank      positionGradeBelongsToRank

	fieldMap map[string]field.Expr
}

func (p positionGrade) Table(newTableName string) *positionGrade {
	p.positionGradeDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p positionGrade) As(alias string) *positionGrade {
	p.positionGradeDo.DO = *(p.positionGradeDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *positionGrade) updateTableName(table string) *positionGrade {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Code = field.NewString(table, "code")
	p.Name = field.NewString(table, "name")
	p.Remark = field.NewString(table, "remark")
	p.RankID = field.NewUint(table, "rank_id")

	p.fillFieldMap()

	return p
}

func (p *positionGrade) WithContext(ctx context.Context) IPositionGradeDo {
	return p.positionGradeDo.WithContext(ctx)
}

func (p positionGrade) TableName() string { return p.positionGradeDo.TableName() }

func (p positionGrade) Alias() string { return p.positionGradeDo.Alias() }

func (p positionGrade) Columns(cols ...field.Expr) gen.Columns {
	return p.positionGradeDo.Columns(cols...)
}

func (p *positionGrade) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *positionGrade) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["code"] = p.Code
	p.fieldMap["name"] = p.Name
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["rank_id"] = p.RankID

}

func (p positionGrade) clone(db *gorm.DB) positionGrade {
	p.positionGradeDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p positionGrade) replaceDB(db *gorm.DB) positionGrade {
	p.positionGradeDo.ReplaceDB(db)
	return p
}

type positionGradeBelongsToRank struct {
	db *gorm.DB

	field.RelationField

	Grade struct {
		field.RelationField
		Rank struct {
			field.RelationField
		}
	}
}

func (a positionGradeBelongsToRank) Where(conds ...field.Expr) *positionGradeBelongsToRank {
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

func (a positionGradeBelongsToRank) WithContext(ctx context.Context) *positionGradeBelongsToRank {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a positionGradeBelongsToRank) Session(session *gorm.Session) *positionGradeBelongsToRank {
	a.db = a.db.Session(session)
	return &a
}

func (a positionGradeBelongsToRank) Model(m *types.PositionGrade) *positionGradeBelongsToRankTx {
	return &positionGradeBelongsToRankTx{a.db.Model(m).Association(a.Name())}
}

type positionGradeBelongsToRankTx struct{ tx *gorm.Association }

func (a positionGradeBelongsToRankTx) Find() (result *types.PositionRank, err error) {
	return result, a.tx.Find(&result)
}

func (a positionGradeBelongsToRankTx) Append(values ...*types.PositionRank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a positionGradeBelongsToRankTx) Replace(values ...*types.PositionRank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a positionGradeBelongsToRankTx) Delete(values ...*types.PositionRank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a positionGradeBelongsToRankTx) Clear() error {
	return a.tx.Clear()
}

func (a positionGradeBelongsToRankTx) Count() int64 {
	return a.tx.Count()
}

type positionGradeDo struct{ gen.DO }

type IPositionGradeDo interface {
	gen.SubQuery
	Debug() IPositionGradeDo
	WithContext(ctx context.Context) IPositionGradeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPositionGradeDo
	WriteDB() IPositionGradeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPositionGradeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPositionGradeDo
	Not(conds ...gen.Condition) IPositionGradeDo
	Or(conds ...gen.Condition) IPositionGradeDo
	Select(conds ...field.Expr) IPositionGradeDo
	Where(conds ...gen.Condition) IPositionGradeDo
	Order(conds ...field.Expr) IPositionGradeDo
	Distinct(cols ...field.Expr) IPositionGradeDo
	Omit(cols ...field.Expr) IPositionGradeDo
	Join(table schema.Tabler, on ...field.Expr) IPositionGradeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPositionGradeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPositionGradeDo
	Group(cols ...field.Expr) IPositionGradeDo
	Having(conds ...gen.Condition) IPositionGradeDo
	Limit(limit int) IPositionGradeDo
	Offset(offset int) IPositionGradeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPositionGradeDo
	Unscoped() IPositionGradeDo
	Create(values ...*types.PositionGrade) error
	CreateInBatches(values []*types.PositionGrade, batchSize int) error
	Save(values ...*types.PositionGrade) error
	First() (*types.PositionGrade, error)
	Take() (*types.PositionGrade, error)
	Last() (*types.PositionGrade, error)
	Find() ([]*types.PositionGrade, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.PositionGrade, err error)
	FindInBatches(result *[]*types.PositionGrade, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*types.PositionGrade) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPositionGradeDo
	Assign(attrs ...field.AssignExpr) IPositionGradeDo
	Joins(fields ...field.RelationField) IPositionGradeDo
	Preload(fields ...field.RelationField) IPositionGradeDo
	FirstOrInit() (*types.PositionGrade, error)
	FirstOrCreate() (*types.PositionGrade, error)
	FindByPage(offset int, limit int) (result []*types.PositionGrade, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPositionGradeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p positionGradeDo) Debug() IPositionGradeDo {
	return p.withDO(p.DO.Debug())
}

func (p positionGradeDo) WithContext(ctx context.Context) IPositionGradeDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p positionGradeDo) ReadDB() IPositionGradeDo {
	return p.Clauses(dbresolver.Read)
}

func (p positionGradeDo) WriteDB() IPositionGradeDo {
	return p.Clauses(dbresolver.Write)
}

func (p positionGradeDo) Session(config *gorm.Session) IPositionGradeDo {
	return p.withDO(p.DO.Session(config))
}

func (p positionGradeDo) Clauses(conds ...clause.Expression) IPositionGradeDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p positionGradeDo) Returning(value interface{}, columns ...string) IPositionGradeDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p positionGradeDo) Not(conds ...gen.Condition) IPositionGradeDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p positionGradeDo) Or(conds ...gen.Condition) IPositionGradeDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p positionGradeDo) Select(conds ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p positionGradeDo) Where(conds ...gen.Condition) IPositionGradeDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p positionGradeDo) Order(conds ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p positionGradeDo) Distinct(cols ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p positionGradeDo) Omit(cols ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p positionGradeDo) Join(table schema.Tabler, on ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p positionGradeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p positionGradeDo) RightJoin(table schema.Tabler, on ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p positionGradeDo) Group(cols ...field.Expr) IPositionGradeDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p positionGradeDo) Having(conds ...gen.Condition) IPositionGradeDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p positionGradeDo) Limit(limit int) IPositionGradeDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p positionGradeDo) Offset(offset int) IPositionGradeDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p positionGradeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPositionGradeDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p positionGradeDo) Unscoped() IPositionGradeDo {
	return p.withDO(p.DO.Unscoped())
}

func (p positionGradeDo) Create(values ...*types.PositionGrade) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p positionGradeDo) CreateInBatches(values []*types.PositionGrade, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p positionGradeDo) Save(values ...*types.PositionGrade) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p positionGradeDo) First() (*types.PositionGrade, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionGrade), nil
	}
}

func (p positionGradeDo) Take() (*types.PositionGrade, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionGrade), nil
	}
}

func (p positionGradeDo) Last() (*types.PositionGrade, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionGrade), nil
	}
}

func (p positionGradeDo) Find() ([]*types.PositionGrade, error) {
	result, err := p.DO.Find()
	return result.([]*types.PositionGrade), err
}

func (p positionGradeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*types.PositionGrade, err error) {
	buf := make([]*types.PositionGrade, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p positionGradeDo) FindInBatches(result *[]*types.PositionGrade, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p positionGradeDo) Attrs(attrs ...field.AssignExpr) IPositionGradeDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p positionGradeDo) Assign(attrs ...field.AssignExpr) IPositionGradeDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p positionGradeDo) Joins(fields ...field.RelationField) IPositionGradeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p positionGradeDo) Preload(fields ...field.RelationField) IPositionGradeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p positionGradeDo) FirstOrInit() (*types.PositionGrade, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionGrade), nil
	}
}

func (p positionGradeDo) FirstOrCreate() (*types.PositionGrade, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*types.PositionGrade), nil
	}
}

func (p positionGradeDo) FindByPage(offset int, limit int) (result []*types.PositionGrade, count int64, err error) {
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

func (p positionGradeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p positionGradeDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p positionGradeDo) Delete(models ...*types.PositionGrade) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *positionGradeDo) withDO(do gen.Dao) *positionGradeDo {
	p.DO = *do.(*gen.DO)
	return p
}
