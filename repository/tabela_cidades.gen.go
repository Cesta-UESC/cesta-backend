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

func newCidades(db *gorm.DB, opts ...gen.DOOption) cidades {
	_cidades := cidades{}

	_cidades.cidadesDo.UseDB(db, opts...)
	_cidades.cidadesDo.UseModel(&model.Cidades{})

	tableName := _cidades.cidadesDo.TableName()
	_cidades.ALL = field.NewAsterisk(tableName)
	_cidades.CidadeID = field.NewInt32(tableName, "cidade_id")
	_cidades.CidadeNome = field.NewString(tableName, "cidade_nome")
	_cidades.CidadeData = field.NewTime(tableName, "cidade_data")

	_cidades.fillFieldMap()

	return _cidades
}

type cidades struct {
	cidadesDo

	ALL        field.Asterisk
	CidadeID   field.Int32
	CidadeNome field.String
	CidadeData field.Time

	fieldMap map[string]field.Expr
}

func (c cidades) Table(newTableName string) *cidades {
	c.cidadesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cidades) As(alias string) *cidades {
	c.cidadesDo.DO = *(c.cidadesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cidades) updateTableName(table string) *cidades {
	c.ALL = field.NewAsterisk(table)
	c.CidadeID = field.NewInt32(table, "cidade_id")
	c.CidadeNome = field.NewString(table, "cidade_nome")
	c.CidadeData = field.NewTime(table, "cidade_data")

	c.fillFieldMap()

	return c
}

func (c *cidades) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cidades) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["cidade_id"] = c.CidadeID
	c.fieldMap["cidade_nome"] = c.CidadeNome
	c.fieldMap["cidade_data"] = c.CidadeData
}

func (c cidades) clone(db *gorm.DB) cidades {
	c.cidadesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cidades) replaceDB(db *gorm.DB) cidades {
	c.cidadesDo.ReplaceDB(db)
	return c
}

type cidadesDo struct{ gen.DO }

type ICidadesDo interface {
	gen.SubQuery
	Debug() ICidadesDo
	WithContext(ctx context.Context) ICidadesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICidadesDo
	WriteDB() ICidadesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICidadesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICidadesDo
	Not(conds ...gen.Condition) ICidadesDo
	Or(conds ...gen.Condition) ICidadesDo
	Select(conds ...field.Expr) ICidadesDo
	Where(conds ...gen.Condition) ICidadesDo
	Order(conds ...field.Expr) ICidadesDo
	Distinct(cols ...field.Expr) ICidadesDo
	Omit(cols ...field.Expr) ICidadesDo
	Join(table schema.Tabler, on ...field.Expr) ICidadesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICidadesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICidadesDo
	Group(cols ...field.Expr) ICidadesDo
	Having(conds ...gen.Condition) ICidadesDo
	Limit(limit int) ICidadesDo
	Offset(offset int) ICidadesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICidadesDo
	Unscoped() ICidadesDo
	Create(values ...*model.Cidades) error
	CreateInBatches(values []*model.Cidades, batchSize int) error
	Save(values ...*model.Cidades) error
	First() (*model.Cidades, error)
	Take() (*model.Cidades, error)
	Last() (*model.Cidades, error)
	Find() ([]*model.Cidades, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cidades, err error)
	FindInBatches(result *[]*model.Cidades, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Cidades) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICidadesDo
	Assign(attrs ...field.AssignExpr) ICidadesDo
	Joins(fields ...field.RelationField) ICidadesDo
	Preload(fields ...field.RelationField) ICidadesDo
	FirstOrInit() (*model.Cidades, error)
	FirstOrCreate() (*model.Cidades, error)
	FindByPage(offset int, limit int) (result []*model.Cidades, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICidadesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cidadesDo) Debug() ICidadesDo {
	return c.withDO(c.DO.Debug())
}

func (c cidadesDo) WithContext(ctx context.Context) ICidadesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cidadesDo) ReadDB() ICidadesDo {
	return c.Clauses(dbresolver.Read)
}

func (c cidadesDo) WriteDB() ICidadesDo {
	return c.Clauses(dbresolver.Write)
}

func (c cidadesDo) Session(config *gorm.Session) ICidadesDo {
	return c.withDO(c.DO.Session(config))
}

func (c cidadesDo) Clauses(conds ...clause.Expression) ICidadesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cidadesDo) Returning(value interface{}, columns ...string) ICidadesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cidadesDo) Not(conds ...gen.Condition) ICidadesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cidadesDo) Or(conds ...gen.Condition) ICidadesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cidadesDo) Select(conds ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cidadesDo) Where(conds ...gen.Condition) ICidadesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cidadesDo) Order(conds ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cidadesDo) Distinct(cols ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cidadesDo) Omit(cols ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cidadesDo) Join(table schema.Tabler, on ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cidadesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cidadesDo) RightJoin(table schema.Tabler, on ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cidadesDo) Group(cols ...field.Expr) ICidadesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cidadesDo) Having(conds ...gen.Condition) ICidadesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cidadesDo) Limit(limit int) ICidadesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cidadesDo) Offset(offset int) ICidadesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cidadesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICidadesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cidadesDo) Unscoped() ICidadesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cidadesDo) Create(values ...*model.Cidades) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cidadesDo) CreateInBatches(values []*model.Cidades, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cidadesDo) Save(values ...*model.Cidades) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cidadesDo) First() (*model.Cidades, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cidades), nil
	}
}

func (c cidadesDo) Take() (*model.Cidades, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cidades), nil
	}
}

func (c cidadesDo) Last() (*model.Cidades, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cidades), nil
	}
}

func (c cidadesDo) Find() ([]*model.Cidades, error) {
	result, err := c.DO.Find()
	return result.([]*model.Cidades), err
}

func (c cidadesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cidades, err error) {
	buf := make([]*model.Cidades, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cidadesDo) FindInBatches(result *[]*model.Cidades, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cidadesDo) Attrs(attrs ...field.AssignExpr) ICidadesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cidadesDo) Assign(attrs ...field.AssignExpr) ICidadesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cidadesDo) Joins(fields ...field.RelationField) ICidadesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cidadesDo) Preload(fields ...field.RelationField) ICidadesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cidadesDo) FirstOrInit() (*model.Cidades, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cidades), nil
	}
}

func (c cidadesDo) FirstOrCreate() (*model.Cidades, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cidades), nil
	}
}

func (c cidadesDo) FindByPage(offset int, limit int) (result []*model.Cidades, count int64, err error) {
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

func (c cidadesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cidadesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cidadesDo) Delete(models ...*model.Cidades) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cidadesDo) withDO(do gen.Dao) *cidadesDo {
	c.DO = *do.(*gen.DO)
	return c
}
