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

func newProdutosMedidas(db *gorm.DB, opts ...gen.DOOption) produtosMedidas {
	_produtosMedidas := produtosMedidas{}

	_produtosMedidas.produtosMedidasDo.UseDB(db, opts...)
	_produtosMedidas.produtosMedidasDo.UseModel(&model.ProdutosMedidas{})

	tableName := _produtosMedidas.produtosMedidasDo.TableName()
	_produtosMedidas.ALL = field.NewAsterisk(tableName)
	_produtosMedidas.ProdutoID = field.NewInt32(tableName, "produto_id")
	_produtosMedidas.MedidaID = field.NewInt32(tableName, "medida_id")
	_produtosMedidas.MedidaPesquisada = field.NewFloat64(tableName, "medida_pesquisada")

	_produtosMedidas.fillFieldMap()

	return _produtosMedidas
}

type produtosMedidas struct {
	produtosMedidasDo

	ALL              field.Asterisk
	ProdutoID        field.Int32
	MedidaID         field.Int32
	MedidaPesquisada field.Float64

	fieldMap map[string]field.Expr
}

func (p produtosMedidas) Table(newTableName string) *produtosMedidas {
	p.produtosMedidasDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p produtosMedidas) As(alias string) *produtosMedidas {
	p.produtosMedidasDo.DO = *(p.produtosMedidasDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *produtosMedidas) updateTableName(table string) *produtosMedidas {
	p.ALL = field.NewAsterisk(table)
	p.ProdutoID = field.NewInt32(table, "produto_id")
	p.MedidaID = field.NewInt32(table, "medida_id")
	p.MedidaPesquisada = field.NewFloat64(table, "medida_pesquisada")

	p.fillFieldMap()

	return p
}

func (p *produtosMedidas) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *produtosMedidas) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["produto_id"] = p.ProdutoID
	p.fieldMap["medida_id"] = p.MedidaID
	p.fieldMap["medida_pesquisada"] = p.MedidaPesquisada
}

func (p produtosMedidas) clone(db *gorm.DB) produtosMedidas {
	p.produtosMedidasDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p produtosMedidas) replaceDB(db *gorm.DB) produtosMedidas {
	p.produtosMedidasDo.ReplaceDB(db)
	return p
}

type produtosMedidasDo struct{ gen.DO }

type IProdutosMedidasDo interface {
	gen.SubQuery
	Debug() IProdutosMedidasDo
	WithContext(ctx context.Context) IProdutosMedidasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProdutosMedidasDo
	WriteDB() IProdutosMedidasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProdutosMedidasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProdutosMedidasDo
	Not(conds ...gen.Condition) IProdutosMedidasDo
	Or(conds ...gen.Condition) IProdutosMedidasDo
	Select(conds ...field.Expr) IProdutosMedidasDo
	Where(conds ...gen.Condition) IProdutosMedidasDo
	Order(conds ...field.Expr) IProdutosMedidasDo
	Distinct(cols ...field.Expr) IProdutosMedidasDo
	Omit(cols ...field.Expr) IProdutosMedidasDo
	Join(table schema.Tabler, on ...field.Expr) IProdutosMedidasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProdutosMedidasDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProdutosMedidasDo
	Group(cols ...field.Expr) IProdutosMedidasDo
	Having(conds ...gen.Condition) IProdutosMedidasDo
	Limit(limit int) IProdutosMedidasDo
	Offset(offset int) IProdutosMedidasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProdutosMedidasDo
	Unscoped() IProdutosMedidasDo
	Create(values ...*model.ProdutosMedidas) error
	CreateInBatches(values []*model.ProdutosMedidas, batchSize int) error
	Save(values ...*model.ProdutosMedidas) error
	First() (*model.ProdutosMedidas, error)
	Take() (*model.ProdutosMedidas, error)
	Last() (*model.ProdutosMedidas, error)
	Find() ([]*model.ProdutosMedidas, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProdutosMedidas, err error)
	FindInBatches(result *[]*model.ProdutosMedidas, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProdutosMedidas) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProdutosMedidasDo
	Assign(attrs ...field.AssignExpr) IProdutosMedidasDo
	Joins(fields ...field.RelationField) IProdutosMedidasDo
	Preload(fields ...field.RelationField) IProdutosMedidasDo
	FirstOrInit() (*model.ProdutosMedidas, error)
	FirstOrCreate() (*model.ProdutosMedidas, error)
	FindByPage(offset int, limit int) (result []*model.ProdutosMedidas, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProdutosMedidasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p produtosMedidasDo) Debug() IProdutosMedidasDo {
	return p.withDO(p.DO.Debug())
}

func (p produtosMedidasDo) WithContext(ctx context.Context) IProdutosMedidasDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p produtosMedidasDo) ReadDB() IProdutosMedidasDo {
	return p.Clauses(dbresolver.Read)
}

func (p produtosMedidasDo) WriteDB() IProdutosMedidasDo {
	return p.Clauses(dbresolver.Write)
}

func (p produtosMedidasDo) Session(config *gorm.Session) IProdutosMedidasDo {
	return p.withDO(p.DO.Session(config))
}

func (p produtosMedidasDo) Clauses(conds ...clause.Expression) IProdutosMedidasDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p produtosMedidasDo) Returning(value interface{}, columns ...string) IProdutosMedidasDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p produtosMedidasDo) Not(conds ...gen.Condition) IProdutosMedidasDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p produtosMedidasDo) Or(conds ...gen.Condition) IProdutosMedidasDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p produtosMedidasDo) Select(conds ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p produtosMedidasDo) Where(conds ...gen.Condition) IProdutosMedidasDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p produtosMedidasDo) Order(conds ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p produtosMedidasDo) Distinct(cols ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p produtosMedidasDo) Omit(cols ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p produtosMedidasDo) Join(table schema.Tabler, on ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p produtosMedidasDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p produtosMedidasDo) RightJoin(table schema.Tabler, on ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p produtosMedidasDo) Group(cols ...field.Expr) IProdutosMedidasDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p produtosMedidasDo) Having(conds ...gen.Condition) IProdutosMedidasDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p produtosMedidasDo) Limit(limit int) IProdutosMedidasDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p produtosMedidasDo) Offset(offset int) IProdutosMedidasDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p produtosMedidasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProdutosMedidasDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p produtosMedidasDo) Unscoped() IProdutosMedidasDo {
	return p.withDO(p.DO.Unscoped())
}

func (p produtosMedidasDo) Create(values ...*model.ProdutosMedidas) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p produtosMedidasDo) CreateInBatches(values []*model.ProdutosMedidas, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p produtosMedidasDo) Save(values ...*model.ProdutosMedidas) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p produtosMedidasDo) First() (*model.ProdutosMedidas, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosMedidas), nil
	}
}

func (p produtosMedidasDo) Take() (*model.ProdutosMedidas, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosMedidas), nil
	}
}

func (p produtosMedidasDo) Last() (*model.ProdutosMedidas, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosMedidas), nil
	}
}

func (p produtosMedidasDo) Find() ([]*model.ProdutosMedidas, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProdutosMedidas), err
}

func (p produtosMedidasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProdutosMedidas, err error) {
	buf := make([]*model.ProdutosMedidas, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p produtosMedidasDo) FindInBatches(result *[]*model.ProdutosMedidas, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p produtosMedidasDo) Attrs(attrs ...field.AssignExpr) IProdutosMedidasDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p produtosMedidasDo) Assign(attrs ...field.AssignExpr) IProdutosMedidasDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p produtosMedidasDo) Joins(fields ...field.RelationField) IProdutosMedidasDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p produtosMedidasDo) Preload(fields ...field.RelationField) IProdutosMedidasDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p produtosMedidasDo) FirstOrInit() (*model.ProdutosMedidas, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosMedidas), nil
	}
}

func (p produtosMedidasDo) FirstOrCreate() (*model.ProdutosMedidas, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosMedidas), nil
	}
}

func (p produtosMedidasDo) FindByPage(offset int, limit int) (result []*model.ProdutosMedidas, count int64, err error) {
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

func (p produtosMedidasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p produtosMedidasDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p produtosMedidasDo) Delete(models ...*model.ProdutosMedidas) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *produtosMedidasDo) withDO(do gen.Dao) *produtosMedidasDo {
	p.DO = *do.(*gen.DO)
	return p
}
