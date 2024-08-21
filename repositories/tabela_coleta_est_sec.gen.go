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

func newColetaEstSec(db *gorm.DB, opts ...gen.DOOption) coletaEstSec {
	_coletaEstSec := coletaEstSec{}

	_coletaEstSec.coletaEstSecDo.UseDB(db, opts...)
	_coletaEstSec.coletaEstSecDo.UseModel(&models.ColetaEstSec{})

	tableName := _coletaEstSec.coletaEstSecDo.TableName()
	_coletaEstSec.ALL = field.NewAsterisk(tableName)
	_coletaEstSec.ColetaEstSecID = field.NewInt32(tableName, "coleta_est_sec_id")
	_coletaEstSec.ColetaID = field.NewInt32(tableName, "coleta_id")
	_coletaEstSec.EstHasSecID = field.NewInt32(tableName, "est_has_sec_id")
	_coletaEstSec.ProdutoID = field.NewInt32(tableName, "produto_id")

	_coletaEstSec.fillFieldMap()

	return _coletaEstSec
}

type coletaEstSec struct {
	coletaEstSecDo

	ALL            field.Asterisk
	ColetaEstSecID field.Int32
	ColetaID       field.Int32
	EstHasSecID    field.Int32
	ProdutoID      field.Int32

	fieldMap map[string]field.Expr
}

func (c coletaEstSec) Table(newTableName string) *coletaEstSec {
	c.coletaEstSecDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c coletaEstSec) As(alias string) *coletaEstSec {
	c.coletaEstSecDo.DO = *(c.coletaEstSecDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *coletaEstSec) updateTableName(table string) *coletaEstSec {
	c.ALL = field.NewAsterisk(table)
	c.ColetaEstSecID = field.NewInt32(table, "coleta_est_sec_id")
	c.ColetaID = field.NewInt32(table, "coleta_id")
	c.EstHasSecID = field.NewInt32(table, "est_has_sec_id")
	c.ProdutoID = field.NewInt32(table, "produto_id")

	c.fillFieldMap()

	return c
}

func (c *coletaEstSec) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *coletaEstSec) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["coleta_est_sec_id"] = c.ColetaEstSecID
	c.fieldMap["coleta_id"] = c.ColetaID
	c.fieldMap["est_has_sec_id"] = c.EstHasSecID
	c.fieldMap["produto_id"] = c.ProdutoID
}

func (c coletaEstSec) clone(db *gorm.DB) coletaEstSec {
	c.coletaEstSecDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c coletaEstSec) replaceDB(db *gorm.DB) coletaEstSec {
	c.coletaEstSecDo.ReplaceDB(db)
	return c
}

type coletaEstSecDo struct{ gen.DO }

type IColetaEstSecDo interface {
	gen.SubQuery
	Debug() IColetaEstSecDo
	WithContext(ctx context.Context) IColetaEstSecDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IColetaEstSecDo
	WriteDB() IColetaEstSecDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IColetaEstSecDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IColetaEstSecDo
	Not(conds ...gen.Condition) IColetaEstSecDo
	Or(conds ...gen.Condition) IColetaEstSecDo
	Select(conds ...field.Expr) IColetaEstSecDo
	Where(conds ...gen.Condition) IColetaEstSecDo
	Order(conds ...field.Expr) IColetaEstSecDo
	Distinct(cols ...field.Expr) IColetaEstSecDo
	Omit(cols ...field.Expr) IColetaEstSecDo
	Join(table schema.Tabler, on ...field.Expr) IColetaEstSecDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IColetaEstSecDo
	RightJoin(table schema.Tabler, on ...field.Expr) IColetaEstSecDo
	Group(cols ...field.Expr) IColetaEstSecDo
	Having(conds ...gen.Condition) IColetaEstSecDo
	Limit(limit int) IColetaEstSecDo
	Offset(offset int) IColetaEstSecDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IColetaEstSecDo
	Unscoped() IColetaEstSecDo
	Create(values ...*models.ColetaEstSec) error
	CreateInBatches(values []*models.ColetaEstSec, batchSize int) error
	Save(values ...*models.ColetaEstSec) error
	First() (*models.ColetaEstSec, error)
	Take() (*models.ColetaEstSec, error)
	Last() (*models.ColetaEstSec, error)
	Find() ([]*models.ColetaEstSec, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ColetaEstSec, err error)
	FindInBatches(result *[]*models.ColetaEstSec, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ColetaEstSec) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IColetaEstSecDo
	Assign(attrs ...field.AssignExpr) IColetaEstSecDo
	Joins(fields ...field.RelationField) IColetaEstSecDo
	Preload(fields ...field.RelationField) IColetaEstSecDo
	FirstOrInit() (*models.ColetaEstSec, error)
	FirstOrCreate() (*models.ColetaEstSec, error)
	FindByPage(offset int, limit int) (result []*models.ColetaEstSec, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IColetaEstSecDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c coletaEstSecDo) Debug() IColetaEstSecDo {
	return c.withDO(c.DO.Debug())
}

func (c coletaEstSecDo) WithContext(ctx context.Context) IColetaEstSecDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c coletaEstSecDo) ReadDB() IColetaEstSecDo {
	return c.Clauses(dbresolver.Read)
}

func (c coletaEstSecDo) WriteDB() IColetaEstSecDo {
	return c.Clauses(dbresolver.Write)
}

func (c coletaEstSecDo) Session(config *gorm.Session) IColetaEstSecDo {
	return c.withDO(c.DO.Session(config))
}

func (c coletaEstSecDo) Clauses(conds ...clause.Expression) IColetaEstSecDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c coletaEstSecDo) Returning(value interface{}, columns ...string) IColetaEstSecDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c coletaEstSecDo) Not(conds ...gen.Condition) IColetaEstSecDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c coletaEstSecDo) Or(conds ...gen.Condition) IColetaEstSecDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c coletaEstSecDo) Select(conds ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c coletaEstSecDo) Where(conds ...gen.Condition) IColetaEstSecDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c coletaEstSecDo) Order(conds ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c coletaEstSecDo) Distinct(cols ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c coletaEstSecDo) Omit(cols ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c coletaEstSecDo) Join(table schema.Tabler, on ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c coletaEstSecDo) LeftJoin(table schema.Tabler, on ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c coletaEstSecDo) RightJoin(table schema.Tabler, on ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c coletaEstSecDo) Group(cols ...field.Expr) IColetaEstSecDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c coletaEstSecDo) Having(conds ...gen.Condition) IColetaEstSecDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c coletaEstSecDo) Limit(limit int) IColetaEstSecDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c coletaEstSecDo) Offset(offset int) IColetaEstSecDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c coletaEstSecDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IColetaEstSecDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c coletaEstSecDo) Unscoped() IColetaEstSecDo {
	return c.withDO(c.DO.Unscoped())
}

func (c coletaEstSecDo) Create(values ...*models.ColetaEstSec) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c coletaEstSecDo) CreateInBatches(values []*models.ColetaEstSec, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c coletaEstSecDo) Save(values ...*models.ColetaEstSec) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c coletaEstSecDo) First() (*models.ColetaEstSec, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ColetaEstSec), nil
	}
}

func (c coletaEstSecDo) Take() (*models.ColetaEstSec, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ColetaEstSec), nil
	}
}

func (c coletaEstSecDo) Last() (*models.ColetaEstSec, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ColetaEstSec), nil
	}
}

func (c coletaEstSecDo) Find() ([]*models.ColetaEstSec, error) {
	result, err := c.DO.Find()
	return result.([]*models.ColetaEstSec), err
}

func (c coletaEstSecDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ColetaEstSec, err error) {
	buf := make([]*models.ColetaEstSec, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c coletaEstSecDo) FindInBatches(result *[]*models.ColetaEstSec, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c coletaEstSecDo) Attrs(attrs ...field.AssignExpr) IColetaEstSecDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c coletaEstSecDo) Assign(attrs ...field.AssignExpr) IColetaEstSecDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c coletaEstSecDo) Joins(fields ...field.RelationField) IColetaEstSecDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c coletaEstSecDo) Preload(fields ...field.RelationField) IColetaEstSecDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c coletaEstSecDo) FirstOrInit() (*models.ColetaEstSec, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ColetaEstSec), nil
	}
}

func (c coletaEstSecDo) FirstOrCreate() (*models.ColetaEstSec, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ColetaEstSec), nil
	}
}

func (c coletaEstSecDo) FindByPage(offset int, limit int) (result []*models.ColetaEstSec, count int64, err error) {
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

func (c coletaEstSecDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c coletaEstSecDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c coletaEstSecDo) Delete(models ...*models.ColetaEstSec) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *coletaEstSecDo) withDO(do gen.Dao) *coletaEstSecDo {
	c.DO = *do.(*gen.DO)
	return c
}
