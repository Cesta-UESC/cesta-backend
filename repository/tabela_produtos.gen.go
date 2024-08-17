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

func newProdutos(db *gorm.DB, opts ...gen.DOOption) produtos {
	_produtos := produtos{}

	_produtos.produtosDo.UseDB(db, opts...)
	_produtos.produtosDo.UseModel(&model.Produtos{})

	tableName := _produtos.produtosDo.TableName()
	_produtos.ALL = field.NewAsterisk(tableName)
	_produtos.ProdutoID = field.NewInt32(tableName, "produto_id")
	_produtos.ProdutoNome = field.NewString(tableName, "produto_nome")
	_produtos.ProdutoCesta = field.NewInt32(tableName, "produto_cesta")
	_produtos.ProdutoNomeVisualizacao = field.NewString(tableName, "produto_nome_visualizacao")
	_produtos.ProdutoTipo = field.NewInt32(tableName, "produto_tipo")

	_produtos.fillFieldMap()

	return _produtos
}

type produtos struct {
	produtosDo

	ALL                     field.Asterisk
	ProdutoID               field.Int32
	ProdutoNome             field.String
	ProdutoCesta            field.Int32
	ProdutoNomeVisualizacao field.String
	ProdutoTipo             field.Int32

	fieldMap map[string]field.Expr
}

func (p produtos) Table(newTableName string) *produtos {
	p.produtosDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p produtos) As(alias string) *produtos {
	p.produtosDo.DO = *(p.produtosDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *produtos) updateTableName(table string) *produtos {
	p.ALL = field.NewAsterisk(table)
	p.ProdutoID = field.NewInt32(table, "produto_id")
	p.ProdutoNome = field.NewString(table, "produto_nome")
	p.ProdutoCesta = field.NewInt32(table, "produto_cesta")
	p.ProdutoNomeVisualizacao = field.NewString(table, "produto_nome_visualizacao")
	p.ProdutoTipo = field.NewInt32(table, "produto_tipo")

	p.fillFieldMap()

	return p
}

func (p *produtos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *produtos) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["produto_id"] = p.ProdutoID
	p.fieldMap["produto_nome"] = p.ProdutoNome
	p.fieldMap["produto_cesta"] = p.ProdutoCesta
	p.fieldMap["produto_nome_visualizacao"] = p.ProdutoNomeVisualizacao
	p.fieldMap["produto_tipo"] = p.ProdutoTipo
}

func (p produtos) clone(db *gorm.DB) produtos {
	p.produtosDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p produtos) replaceDB(db *gorm.DB) produtos {
	p.produtosDo.ReplaceDB(db)
	return p
}

type produtosDo struct{ gen.DO }

type IProdutosDo interface {
	gen.SubQuery
	Debug() IProdutosDo
	WithContext(ctx context.Context) IProdutosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProdutosDo
	WriteDB() IProdutosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProdutosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProdutosDo
	Not(conds ...gen.Condition) IProdutosDo
	Or(conds ...gen.Condition) IProdutosDo
	Select(conds ...field.Expr) IProdutosDo
	Where(conds ...gen.Condition) IProdutosDo
	Order(conds ...field.Expr) IProdutosDo
	Distinct(cols ...field.Expr) IProdutosDo
	Omit(cols ...field.Expr) IProdutosDo
	Join(table schema.Tabler, on ...field.Expr) IProdutosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProdutosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProdutosDo
	Group(cols ...field.Expr) IProdutosDo
	Having(conds ...gen.Condition) IProdutosDo
	Limit(limit int) IProdutosDo
	Offset(offset int) IProdutosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProdutosDo
	Unscoped() IProdutosDo
	Create(values ...*model.Produtos) error
	CreateInBatches(values []*model.Produtos, batchSize int) error
	Save(values ...*model.Produtos) error
	First() (*model.Produtos, error)
	Take() (*model.Produtos, error)
	Last() (*model.Produtos, error)
	Find() ([]*model.Produtos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Produtos, err error)
	FindInBatches(result *[]*model.Produtos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Produtos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProdutosDo
	Assign(attrs ...field.AssignExpr) IProdutosDo
	Joins(fields ...field.RelationField) IProdutosDo
	Preload(fields ...field.RelationField) IProdutosDo
	FirstOrInit() (*model.Produtos, error)
	FirstOrCreate() (*model.Produtos, error)
	FindByPage(offset int, limit int) (result []*model.Produtos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProdutosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p produtosDo) Debug() IProdutosDo {
	return p.withDO(p.DO.Debug())
}

func (p produtosDo) WithContext(ctx context.Context) IProdutosDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p produtosDo) ReadDB() IProdutosDo {
	return p.Clauses(dbresolver.Read)
}

func (p produtosDo) WriteDB() IProdutosDo {
	return p.Clauses(dbresolver.Write)
}

func (p produtosDo) Session(config *gorm.Session) IProdutosDo {
	return p.withDO(p.DO.Session(config))
}

func (p produtosDo) Clauses(conds ...clause.Expression) IProdutosDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p produtosDo) Returning(value interface{}, columns ...string) IProdutosDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p produtosDo) Not(conds ...gen.Condition) IProdutosDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p produtosDo) Or(conds ...gen.Condition) IProdutosDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p produtosDo) Select(conds ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p produtosDo) Where(conds ...gen.Condition) IProdutosDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p produtosDo) Order(conds ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p produtosDo) Distinct(cols ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p produtosDo) Omit(cols ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p produtosDo) Join(table schema.Tabler, on ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p produtosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p produtosDo) RightJoin(table schema.Tabler, on ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p produtosDo) Group(cols ...field.Expr) IProdutosDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p produtosDo) Having(conds ...gen.Condition) IProdutosDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p produtosDo) Limit(limit int) IProdutosDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p produtosDo) Offset(offset int) IProdutosDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p produtosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProdutosDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p produtosDo) Unscoped() IProdutosDo {
	return p.withDO(p.DO.Unscoped())
}

func (p produtosDo) Create(values ...*model.Produtos) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p produtosDo) CreateInBatches(values []*model.Produtos, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p produtosDo) Save(values ...*model.Produtos) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p produtosDo) First() (*model.Produtos, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Produtos), nil
	}
}

func (p produtosDo) Take() (*model.Produtos, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Produtos), nil
	}
}

func (p produtosDo) Last() (*model.Produtos, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Produtos), nil
	}
}

func (p produtosDo) Find() ([]*model.Produtos, error) {
	result, err := p.DO.Find()
	return result.([]*model.Produtos), err
}

func (p produtosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Produtos, err error) {
	buf := make([]*model.Produtos, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p produtosDo) FindInBatches(result *[]*model.Produtos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p produtosDo) Attrs(attrs ...field.AssignExpr) IProdutosDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p produtosDo) Assign(attrs ...field.AssignExpr) IProdutosDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p produtosDo) Joins(fields ...field.RelationField) IProdutosDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p produtosDo) Preload(fields ...field.RelationField) IProdutosDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p produtosDo) FirstOrInit() (*model.Produtos, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Produtos), nil
	}
}

func (p produtosDo) FirstOrCreate() (*model.Produtos, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Produtos), nil
	}
}

func (p produtosDo) FindByPage(offset int, limit int) (result []*model.Produtos, count int64, err error) {
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

func (p produtosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p produtosDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p produtosDo) Delete(models ...*model.Produtos) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *produtosDo) withDO(do gen.Dao) *produtosDo {
	p.DO = *do.(*gen.DO)
	return p
}
