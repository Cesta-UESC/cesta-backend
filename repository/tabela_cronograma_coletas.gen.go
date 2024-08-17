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

func newCronogramaColetas(db *gorm.DB, opts ...gen.DOOption) cronogramaColetas {
	_cronogramaColetas := cronogramaColetas{}

	_cronogramaColetas.cronogramaColetasDo.UseDB(db, opts...)
	_cronogramaColetas.cronogramaColetasDo.UseModel(&model.CronogramaColetas{})

	tableName := _cronogramaColetas.cronogramaColetasDo.TableName()
	_cronogramaColetas.ALL = field.NewAsterisk(tableName)
	_cronogramaColetas.CronogramaID = field.NewInt32(tableName, "cronograma_id")
	_cronogramaColetas.Ano = field.NewInt32(tableName, "ano")

	_cronogramaColetas.fillFieldMap()

	return _cronogramaColetas
}

type cronogramaColetas struct {
	cronogramaColetasDo

	ALL          field.Asterisk
	CronogramaID field.Int32
	Ano          field.Int32

	fieldMap map[string]field.Expr
}

func (c cronogramaColetas) Table(newTableName string) *cronogramaColetas {
	c.cronogramaColetasDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cronogramaColetas) As(alias string) *cronogramaColetas {
	c.cronogramaColetasDo.DO = *(c.cronogramaColetasDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cronogramaColetas) updateTableName(table string) *cronogramaColetas {
	c.ALL = field.NewAsterisk(table)
	c.CronogramaID = field.NewInt32(table, "cronograma_id")
	c.Ano = field.NewInt32(table, "ano")

	c.fillFieldMap()

	return c
}

func (c *cronogramaColetas) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cronogramaColetas) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["cronograma_id"] = c.CronogramaID
	c.fieldMap["ano"] = c.Ano
}

func (c cronogramaColetas) clone(db *gorm.DB) cronogramaColetas {
	c.cronogramaColetasDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cronogramaColetas) replaceDB(db *gorm.DB) cronogramaColetas {
	c.cronogramaColetasDo.ReplaceDB(db)
	return c
}

type cronogramaColetasDo struct{ gen.DO }

type ICronogramaColetasDo interface {
	gen.SubQuery
	Debug() ICronogramaColetasDo
	WithContext(ctx context.Context) ICronogramaColetasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICronogramaColetasDo
	WriteDB() ICronogramaColetasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICronogramaColetasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICronogramaColetasDo
	Not(conds ...gen.Condition) ICronogramaColetasDo
	Or(conds ...gen.Condition) ICronogramaColetasDo
	Select(conds ...field.Expr) ICronogramaColetasDo
	Where(conds ...gen.Condition) ICronogramaColetasDo
	Order(conds ...field.Expr) ICronogramaColetasDo
	Distinct(cols ...field.Expr) ICronogramaColetasDo
	Omit(cols ...field.Expr) ICronogramaColetasDo
	Join(table schema.Tabler, on ...field.Expr) ICronogramaColetasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICronogramaColetasDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICronogramaColetasDo
	Group(cols ...field.Expr) ICronogramaColetasDo
	Having(conds ...gen.Condition) ICronogramaColetasDo
	Limit(limit int) ICronogramaColetasDo
	Offset(offset int) ICronogramaColetasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICronogramaColetasDo
	Unscoped() ICronogramaColetasDo
	Create(values ...*model.CronogramaColetas) error
	CreateInBatches(values []*model.CronogramaColetas, batchSize int) error
	Save(values ...*model.CronogramaColetas) error
	First() (*model.CronogramaColetas, error)
	Take() (*model.CronogramaColetas, error)
	Last() (*model.CronogramaColetas, error)
	Find() ([]*model.CronogramaColetas, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CronogramaColetas, err error)
	FindInBatches(result *[]*model.CronogramaColetas, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CronogramaColetas) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICronogramaColetasDo
	Assign(attrs ...field.AssignExpr) ICronogramaColetasDo
	Joins(fields ...field.RelationField) ICronogramaColetasDo
	Preload(fields ...field.RelationField) ICronogramaColetasDo
	FirstOrInit() (*model.CronogramaColetas, error)
	FirstOrCreate() (*model.CronogramaColetas, error)
	FindByPage(offset int, limit int) (result []*model.CronogramaColetas, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICronogramaColetasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cronogramaColetasDo) Debug() ICronogramaColetasDo {
	return c.withDO(c.DO.Debug())
}

func (c cronogramaColetasDo) WithContext(ctx context.Context) ICronogramaColetasDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cronogramaColetasDo) ReadDB() ICronogramaColetasDo {
	return c.Clauses(dbresolver.Read)
}

func (c cronogramaColetasDo) WriteDB() ICronogramaColetasDo {
	return c.Clauses(dbresolver.Write)
}

func (c cronogramaColetasDo) Session(config *gorm.Session) ICronogramaColetasDo {
	return c.withDO(c.DO.Session(config))
}

func (c cronogramaColetasDo) Clauses(conds ...clause.Expression) ICronogramaColetasDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cronogramaColetasDo) Returning(value interface{}, columns ...string) ICronogramaColetasDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cronogramaColetasDo) Not(conds ...gen.Condition) ICronogramaColetasDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cronogramaColetasDo) Or(conds ...gen.Condition) ICronogramaColetasDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cronogramaColetasDo) Select(conds ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cronogramaColetasDo) Where(conds ...gen.Condition) ICronogramaColetasDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cronogramaColetasDo) Order(conds ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cronogramaColetasDo) Distinct(cols ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cronogramaColetasDo) Omit(cols ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cronogramaColetasDo) Join(table schema.Tabler, on ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cronogramaColetasDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cronogramaColetasDo) RightJoin(table schema.Tabler, on ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cronogramaColetasDo) Group(cols ...field.Expr) ICronogramaColetasDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cronogramaColetasDo) Having(conds ...gen.Condition) ICronogramaColetasDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cronogramaColetasDo) Limit(limit int) ICronogramaColetasDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cronogramaColetasDo) Offset(offset int) ICronogramaColetasDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cronogramaColetasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICronogramaColetasDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cronogramaColetasDo) Unscoped() ICronogramaColetasDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cronogramaColetasDo) Create(values ...*model.CronogramaColetas) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cronogramaColetasDo) CreateInBatches(values []*model.CronogramaColetas, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cronogramaColetasDo) Save(values ...*model.CronogramaColetas) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cronogramaColetasDo) First() (*model.CronogramaColetas, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CronogramaColetas), nil
	}
}

func (c cronogramaColetasDo) Take() (*model.CronogramaColetas, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CronogramaColetas), nil
	}
}

func (c cronogramaColetasDo) Last() (*model.CronogramaColetas, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CronogramaColetas), nil
	}
}

func (c cronogramaColetasDo) Find() ([]*model.CronogramaColetas, error) {
	result, err := c.DO.Find()
	return result.([]*model.CronogramaColetas), err
}

func (c cronogramaColetasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CronogramaColetas, err error) {
	buf := make([]*model.CronogramaColetas, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cronogramaColetasDo) FindInBatches(result *[]*model.CronogramaColetas, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cronogramaColetasDo) Attrs(attrs ...field.AssignExpr) ICronogramaColetasDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cronogramaColetasDo) Assign(attrs ...field.AssignExpr) ICronogramaColetasDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cronogramaColetasDo) Joins(fields ...field.RelationField) ICronogramaColetasDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cronogramaColetasDo) Preload(fields ...field.RelationField) ICronogramaColetasDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cronogramaColetasDo) FirstOrInit() (*model.CronogramaColetas, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CronogramaColetas), nil
	}
}

func (c cronogramaColetasDo) FirstOrCreate() (*model.CronogramaColetas, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CronogramaColetas), nil
	}
}

func (c cronogramaColetasDo) FindByPage(offset int, limit int) (result []*model.CronogramaColetas, count int64, err error) {
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

func (c cronogramaColetasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cronogramaColetasDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cronogramaColetasDo) Delete(models ...*model.CronogramaColetas) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cronogramaColetasDo) withDO(do gen.Dao) *cronogramaColetasDo {
	c.DO = *do.(*gen.DO)
	return c
}
