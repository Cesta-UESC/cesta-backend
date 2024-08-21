package repositories

import (
	"context"

	"github.com/Cesta-UESC/cesta-backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newColetas(db *gorm.DB, opts ...gen.DOOption) coletas {
	_coletas := coletas{}

	_coletas.coletasDo.UseDB(db, opts...)
	_coletas.coletasDo.UseModel(&models.Coletas{})

	tableName := _coletas.coletasDo.TableName()
	_coletas.ALL = field.NewAsterisk(tableName)
	_coletas.ColetaID = field.NewInt32(tableName, "coleta_id")
	_coletas.EstabelecimentoID = field.NewInt32(tableName, "estabelecimento_id")
	_coletas.PesquisaID = field.NewInt32(tableName, "pesquisa_id")
	_coletas.ColetaData = field.NewTime(tableName, "coleta_data")
	_coletas.ColetaPrecoCesta = field.NewFloat64(tableName, "coleta_preco_cesta")
	_coletas.ColetaFechada = field.NewInt32(tableName, "coleta_fechada")

	_coletas.fillFieldMap()

	return _coletas
}

type coletas struct {
	coletasDo

	ALL               field.Asterisk
	ColetaID          field.Int32
	EstabelecimentoID field.Int32
	PesquisaID        field.Int32
	ColetaData        field.Time
	ColetaPrecoCesta  field.Float64
	ColetaFechada     field.Int32

	fieldMap map[string]field.Expr
}

func (c coletas) Table(newTableName string) *coletas {
	c.coletasDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c coletas) As(alias string) *coletas {
	c.coletasDo.DO = *(c.coletasDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *coletas) updateTableName(table string) *coletas {
	c.ALL = field.NewAsterisk(table)
	c.ColetaID = field.NewInt32(table, "coleta_id")
	c.EstabelecimentoID = field.NewInt32(table, "estabelecimento_id")
	c.PesquisaID = field.NewInt32(table, "pesquisa_id")
	c.ColetaData = field.NewTime(table, "coleta_data")
	c.ColetaPrecoCesta = field.NewFloat64(table, "coleta_preco_cesta")
	c.ColetaFechada = field.NewInt32(table, "coleta_fechada")

	c.fillFieldMap()

	return c
}

func (c *coletas) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *coletas) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["coleta_id"] = c.ColetaID
	c.fieldMap["estabelecimento_id"] = c.EstabelecimentoID
	c.fieldMap["pesquisa_id"] = c.PesquisaID
	c.fieldMap["coleta_data"] = c.ColetaData
	c.fieldMap["coleta_preco_cesta"] = c.ColetaPrecoCesta
	c.fieldMap["coleta_fechada"] = c.ColetaFechada
}

func (c coletas) clone(db *gorm.DB) coletas {
	c.coletasDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c coletas) replaceDB(db *gorm.DB) coletas {
	c.coletasDo.ReplaceDB(db)
	return c
}

type coletasDo struct{ gen.DO }

type IColetasDo interface {
	gen.SubQuery
	Debug() IColetasDo
	WithContext(ctx context.Context) IColetasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IColetasDo
	WriteDB() IColetasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IColetasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IColetasDo
	Not(conds ...gen.Condition) IColetasDo
	Or(conds ...gen.Condition) IColetasDo
	Select(conds ...field.Expr) IColetasDo
	Where(conds ...gen.Condition) IColetasDo
	Order(conds ...field.Expr) IColetasDo
	Distinct(cols ...field.Expr) IColetasDo
	Omit(cols ...field.Expr) IColetasDo
	Join(table schema.Tabler, on ...field.Expr) IColetasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IColetasDo
	RightJoin(table schema.Tabler, on ...field.Expr) IColetasDo
	Group(cols ...field.Expr) IColetasDo
	Having(conds ...gen.Condition) IColetasDo
	Limit(limit int) IColetasDo
	Offset(offset int) IColetasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IColetasDo
	Unscoped() IColetasDo
	Create(values ...*models.Coletas) error
	CreateInBatches(values []*models.Coletas, batchSize int) error
	Save(values ...*models.Coletas) error
	First() (*models.Coletas, error)
	Take() (*models.Coletas, error)
	Last() (*models.Coletas, error)
	Find() ([]*models.Coletas, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Coletas, err error)
	FindInBatches(result *[]*models.Coletas, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Coletas) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IColetasDo
	Assign(attrs ...field.AssignExpr) IColetasDo
	Joins(fields ...field.RelationField) IColetasDo
	Preload(fields ...field.RelationField) IColetasDo
	FirstOrInit() (*models.Coletas, error)
	FirstOrCreate() (*models.Coletas, error)
	FindByPage(offset int, limit int) (result []*models.Coletas, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IColetasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c coletasDo) Debug() IColetasDo {
	return c.withDO(c.DO.Debug())
}

func (c coletasDo) WithContext(ctx context.Context) IColetasDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c coletasDo) ReadDB() IColetasDo {
	return c.Clauses(dbresolver.Read)
}

func (c coletasDo) WriteDB() IColetasDo {
	return c.Clauses(dbresolver.Write)
}

func (c coletasDo) Session(config *gorm.Session) IColetasDo {
	return c.withDO(c.DO.Session(config))
}

func (c coletasDo) Clauses(conds ...clause.Expression) IColetasDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c coletasDo) Returning(value interface{}, columns ...string) IColetasDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c coletasDo) Not(conds ...gen.Condition) IColetasDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c coletasDo) Or(conds ...gen.Condition) IColetasDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c coletasDo) Select(conds ...field.Expr) IColetasDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c coletasDo) Where(conds ...gen.Condition) IColetasDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c coletasDo) Order(conds ...field.Expr) IColetasDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c coletasDo) Distinct(cols ...field.Expr) IColetasDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c coletasDo) Omit(cols ...field.Expr) IColetasDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c coletasDo) Join(table schema.Tabler, on ...field.Expr) IColetasDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c coletasDo) LeftJoin(table schema.Tabler, on ...field.Expr) IColetasDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c coletasDo) RightJoin(table schema.Tabler, on ...field.Expr) IColetasDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c coletasDo) Group(cols ...field.Expr) IColetasDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c coletasDo) Having(conds ...gen.Condition) IColetasDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c coletasDo) Limit(limit int) IColetasDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c coletasDo) Offset(offset int) IColetasDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c coletasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IColetasDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c coletasDo) Unscoped() IColetasDo {
	return c.withDO(c.DO.Unscoped())
}

func (c coletasDo) Create(values ...*models.Coletas) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c coletasDo) CreateInBatches(values []*models.Coletas, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c coletasDo) Save(values ...*models.Coletas) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c coletasDo) First() (*models.Coletas, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Coletas), nil
	}
}

func (c coletasDo) Take() (*models.Coletas, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Coletas), nil
	}
}

func (c coletasDo) Last() (*models.Coletas, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Coletas), nil
	}
}

func (c coletasDo) Find() ([]*models.Coletas, error) {
	result, err := c.DO.Find()
	return result.([]*models.Coletas), err
}

func (c coletasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Coletas, err error) {
	buf := make([]*models.Coletas, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c coletasDo) FindInBatches(result *[]*models.Coletas, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c coletasDo) Attrs(attrs ...field.AssignExpr) IColetasDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c coletasDo) Assign(attrs ...field.AssignExpr) IColetasDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c coletasDo) Joins(fields ...field.RelationField) IColetasDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c coletasDo) Preload(fields ...field.RelationField) IColetasDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c coletasDo) FirstOrInit() (*models.Coletas, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Coletas), nil
	}
}

func (c coletasDo) FirstOrCreate() (*models.Coletas, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Coletas), nil
	}
}

func (c coletasDo) FindByPage(offset int, limit int) (result []*models.Coletas, count int64, err error) {
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

func (c coletasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c coletasDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c coletasDo) Delete(models ...*models.Coletas) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *coletasDo) withDO(do gen.Dao) *coletasDo {
	c.DO = *do.(*gen.DO)
	return c
}
