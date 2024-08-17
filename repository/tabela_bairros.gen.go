package repository

import (
	"context"

	"github.com/Cesta-UESC/cesta-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newBairros(db *gorm.DB, opts ...gen.DOOption) bairros {
	_bairros := bairros{}

	_bairros.bairrosDo.UseDB(db, opts...)
	_bairros.bairrosDo.UseModel(&model.Bairros{})

	tableName := _bairros.bairrosDo.TableName()
	_bairros.ALL = field.NewAsterisk(tableName)
	_bairros.BairroID = field.NewInt32(tableName, "bairro_id")
	_bairros.BairroNome = field.NewString(tableName, "bairro_nome")
	_bairros.CidadeID = field.NewInt32(tableName, "cidade_id")

	_bairros.fillFieldMap()

	return _bairros
}

type bairros struct {
	bairrosDo

	ALL        field.Asterisk
	BairroID   field.Int32
	BairroNome field.String
	CidadeID   field.Int32

	fieldMap map[string]field.Expr
}

func (b bairros) Table(newTableName string) *bairros {
	b.bairrosDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bairros) As(alias string) *bairros {
	b.bairrosDo.DO = *(b.bairrosDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bairros) updateTableName(table string) *bairros {
	b.ALL = field.NewAsterisk(table)
	b.BairroID = field.NewInt32(table, "bairro_id")
	b.BairroNome = field.NewString(table, "bairro_nome")
	b.CidadeID = field.NewInt32(table, "cidade_id")

	b.fillFieldMap()

	return b
}

func (b *bairros) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bairros) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 3)
	b.fieldMap["bairro_id"] = b.BairroID
	b.fieldMap["bairro_nome"] = b.BairroNome
	b.fieldMap["cidade_id"] = b.CidadeID
}

func (b bairros) clone(db *gorm.DB) bairros {
	b.bairrosDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bairros) replaceDB(db *gorm.DB) bairros {
	b.bairrosDo.ReplaceDB(db)
	return b
}

type bairrosDo struct{ gen.DO }

type IBairrosDo interface {
	gen.SubQuery
	Debug() IBairrosDo
	WithContext(ctx context.Context) IBairrosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBairrosDo
	WriteDB() IBairrosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBairrosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBairrosDo
	Not(conds ...gen.Condition) IBairrosDo
	Or(conds ...gen.Condition) IBairrosDo
	Select(conds ...field.Expr) IBairrosDo
	Where(conds ...gen.Condition) IBairrosDo
	Order(conds ...field.Expr) IBairrosDo
	Distinct(cols ...field.Expr) IBairrosDo
	Omit(cols ...field.Expr) IBairrosDo
	Join(table schema.Tabler, on ...field.Expr) IBairrosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBairrosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBairrosDo
	Group(cols ...field.Expr) IBairrosDo
	Having(conds ...gen.Condition) IBairrosDo
	Limit(limit int) IBairrosDo
	Offset(offset int) IBairrosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBairrosDo
	Unscoped() IBairrosDo
	Create(values ...*model.Bairros) error
	CreateInBatches(values []*model.Bairros, batchSize int) error
	Save(values ...*model.Bairros) error
	First() (*model.Bairros, error)
	Take() (*model.Bairros, error)
	Last() (*model.Bairros, error)
	Find() ([]*model.Bairros, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bairros, err error)
	FindInBatches(result *[]*model.Bairros, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Bairros) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBairrosDo
	Assign(attrs ...field.AssignExpr) IBairrosDo
	Joins(fields ...field.RelationField) IBairrosDo
	Preload(fields ...field.RelationField) IBairrosDo
	FirstOrInit() (*model.Bairros, error)
	FirstOrCreate() (*model.Bairros, error)
	FindByPage(offset int, limit int) (result []*model.Bairros, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBairrosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bairrosDo) Debug() IBairrosDo {
	return b.withDO(b.DO.Debug())
}

func (b bairrosDo) WithContext(ctx context.Context) IBairrosDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bairrosDo) ReadDB() IBairrosDo {
	return b.Clauses(dbresolver.Read)
}

func (b bairrosDo) WriteDB() IBairrosDo {
	return b.Clauses(dbresolver.Write)
}

func (b bairrosDo) Session(config *gorm.Session) IBairrosDo {
	return b.withDO(b.DO.Session(config))
}

func (b bairrosDo) Clauses(conds ...clause.Expression) IBairrosDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bairrosDo) Returning(value interface{}, columns ...string) IBairrosDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bairrosDo) Not(conds ...gen.Condition) IBairrosDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bairrosDo) Or(conds ...gen.Condition) IBairrosDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bairrosDo) Select(conds ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bairrosDo) Where(conds ...gen.Condition) IBairrosDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bairrosDo) Order(conds ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bairrosDo) Distinct(cols ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bairrosDo) Omit(cols ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bairrosDo) Join(table schema.Tabler, on ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bairrosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bairrosDo) RightJoin(table schema.Tabler, on ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bairrosDo) Group(cols ...field.Expr) IBairrosDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bairrosDo) Having(conds ...gen.Condition) IBairrosDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bairrosDo) Limit(limit int) IBairrosDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bairrosDo) Offset(offset int) IBairrosDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bairrosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBairrosDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bairrosDo) Unscoped() IBairrosDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bairrosDo) Create(values ...*model.Bairros) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bairrosDo) CreateInBatches(values []*model.Bairros, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bairrosDo) Save(values ...*model.Bairros) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bairrosDo) First() (*model.Bairros, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bairros), nil
	}
}

func (b bairrosDo) Take() (*model.Bairros, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bairros), nil
	}
}

func (b bairrosDo) Last() (*model.Bairros, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bairros), nil
	}
}

func (b bairrosDo) Find() ([]*model.Bairros, error) {
	result, err := b.DO.Find()
	return result.([]*model.Bairros), err
}

func (b bairrosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bairros, err error) {
	buf := make([]*model.Bairros, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bairrosDo) FindInBatches(result *[]*model.Bairros, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bairrosDo) Attrs(attrs ...field.AssignExpr) IBairrosDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bairrosDo) Assign(attrs ...field.AssignExpr) IBairrosDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bairrosDo) Joins(fields ...field.RelationField) IBairrosDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bairrosDo) Preload(fields ...field.RelationField) IBairrosDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bairrosDo) FirstOrInit() (*model.Bairros, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bairros), nil
	}
}

func (b bairrosDo) FirstOrCreate() (*model.Bairros, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bairros), nil
	}
}

func (b bairrosDo) FindByPage(offset int, limit int) (result []*model.Bairros, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bairrosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bairrosDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bairrosDo) Delete(models ...*model.Bairros) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bairrosDo) withDO(do gen.Dao) *bairrosDo {
	b.DO = *do.(*gen.DO)
	return b
}
