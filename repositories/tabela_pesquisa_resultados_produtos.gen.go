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

func newPesquisaResultadosProdutos(db *gorm.DB, opts ...gen.DOOption) pesquisaResultadosProdutos {
	_pesquisaResultadosProdutos := pesquisaResultadosProdutos{}

	_pesquisaResultadosProdutos.pesquisaResultadosProdutosDo.UseDB(db, opts...)
	_pesquisaResultadosProdutos.pesquisaResultadosProdutosDo.UseModel(&models.PesquisaResultadosProdutos{})

	tableName := _pesquisaResultadosProdutos.pesquisaResultadosProdutosDo.TableName()
	_pesquisaResultadosProdutos.ALL = field.NewAsterisk(tableName)
	_pesquisaResultadosProdutos.ProdutoID = field.NewInt32(tableName, "produto_id")
	_pesquisaResultadosProdutos.CidadeID = field.NewInt32(tableName, "cidade_id")
	_pesquisaResultadosProdutos.PesquisaID = field.NewInt32(tableName, "pesquisa_id")
	_pesquisaResultadosProdutos.ProdutoVariacaoMensal = field.NewFloat64(tableName, "produto_variacao_mensal")
	_pesquisaResultadosProdutos.ProdutoVariacaoSemestral = field.NewFloat64(tableName, "produto_variacao_semestral")
	_pesquisaResultadosProdutos.ProdutoVariacaoAnual = field.NewFloat64(tableName, "produto_variacao_anual")
	_pesquisaResultadosProdutos.ProdutoPrecoMedio = field.NewFloat64(tableName, "produto_preco_medio")
	_pesquisaResultadosProdutos.ProdutoPrecoTotal = field.NewFloat64(tableName, "produto_preco_total")
	_pesquisaResultadosProdutos.ProdutoTempoTrabalho = field.NewFloat64(tableName, "produto_tempo_trabalho")

	_pesquisaResultadosProdutos.fillFieldMap()

	return _pesquisaResultadosProdutos
}

type pesquisaResultadosProdutos struct {
	pesquisaResultadosProdutosDo

	ALL                      field.Asterisk
	ProdutoID                field.Int32
	CidadeID                 field.Int32
	PesquisaID               field.Int32
	ProdutoVariacaoMensal    field.Float64
	ProdutoVariacaoSemestral field.Float64
	ProdutoVariacaoAnual     field.Float64
	ProdutoPrecoMedio        field.Float64
	ProdutoPrecoTotal        field.Float64
	ProdutoTempoTrabalho     field.Float64

	fieldMap map[string]field.Expr
}

func (p pesquisaResultadosProdutos) Table(newTableName string) *pesquisaResultadosProdutos {
	p.pesquisaResultadosProdutosDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pesquisaResultadosProdutos) As(alias string) *pesquisaResultadosProdutos {
	p.pesquisaResultadosProdutosDo.DO = *(p.pesquisaResultadosProdutosDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pesquisaResultadosProdutos) updateTableName(table string) *pesquisaResultadosProdutos {
	p.ALL = field.NewAsterisk(table)
	p.ProdutoID = field.NewInt32(table, "produto_id")
	p.CidadeID = field.NewInt32(table, "cidade_id")
	p.PesquisaID = field.NewInt32(table, "pesquisa_id")
	p.ProdutoVariacaoMensal = field.NewFloat64(table, "produto_variacao_mensal")
	p.ProdutoVariacaoSemestral = field.NewFloat64(table, "produto_variacao_semestral")
	p.ProdutoVariacaoAnual = field.NewFloat64(table, "produto_variacao_anual")
	p.ProdutoPrecoMedio = field.NewFloat64(table, "produto_preco_medio")
	p.ProdutoPrecoTotal = field.NewFloat64(table, "produto_preco_total")
	p.ProdutoTempoTrabalho = field.NewFloat64(table, "produto_tempo_trabalho")

	p.fillFieldMap()

	return p
}

func (p *pesquisaResultadosProdutos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pesquisaResultadosProdutos) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["produto_id"] = p.ProdutoID
	p.fieldMap["cidade_id"] = p.CidadeID
	p.fieldMap["pesquisa_id"] = p.PesquisaID
	p.fieldMap["produto_variacao_mensal"] = p.ProdutoVariacaoMensal
	p.fieldMap["produto_variacao_semestral"] = p.ProdutoVariacaoSemestral
	p.fieldMap["produto_variacao_anual"] = p.ProdutoVariacaoAnual
	p.fieldMap["produto_preco_medio"] = p.ProdutoPrecoMedio
	p.fieldMap["produto_preco_total"] = p.ProdutoPrecoTotal
	p.fieldMap["produto_tempo_trabalho"] = p.ProdutoTempoTrabalho
}

func (p pesquisaResultadosProdutos) clone(db *gorm.DB) pesquisaResultadosProdutos {
	p.pesquisaResultadosProdutosDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pesquisaResultadosProdutos) replaceDB(db *gorm.DB) pesquisaResultadosProdutos {
	p.pesquisaResultadosProdutosDo.ReplaceDB(db)
	return p
}

type pesquisaResultadosProdutosDo struct{ gen.DO }

type IPesquisaResultadosProdutosDo interface {
	gen.SubQuery
	Debug() IPesquisaResultadosProdutosDo
	WithContext(ctx context.Context) IPesquisaResultadosProdutosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPesquisaResultadosProdutosDo
	WriteDB() IPesquisaResultadosProdutosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPesquisaResultadosProdutosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPesquisaResultadosProdutosDo
	Not(conds ...gen.Condition) IPesquisaResultadosProdutosDo
	Or(conds ...gen.Condition) IPesquisaResultadosProdutosDo
	Select(conds ...field.Expr) IPesquisaResultadosProdutosDo
	Where(conds ...gen.Condition) IPesquisaResultadosProdutosDo
	Order(conds ...field.Expr) IPesquisaResultadosProdutosDo
	Distinct(cols ...field.Expr) IPesquisaResultadosProdutosDo
	Omit(cols ...field.Expr) IPesquisaResultadosProdutosDo
	Join(table schema.Tabler, on ...field.Expr) IPesquisaResultadosProdutosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPesquisaResultadosProdutosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPesquisaResultadosProdutosDo
	Group(cols ...field.Expr) IPesquisaResultadosProdutosDo
	Having(conds ...gen.Condition) IPesquisaResultadosProdutosDo
	Limit(limit int) IPesquisaResultadosProdutosDo
	Offset(offset int) IPesquisaResultadosProdutosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPesquisaResultadosProdutosDo
	Unscoped() IPesquisaResultadosProdutosDo
	Create(values ...*models.PesquisaResultadosProdutos) error
	CreateInBatches(values []*models.PesquisaResultadosProdutos, batchSize int) error
	Save(values ...*models.PesquisaResultadosProdutos) error
	First() (*models.PesquisaResultadosProdutos, error)
	Take() (*models.PesquisaResultadosProdutos, error)
	Last() (*models.PesquisaResultadosProdutos, error)
	Find() ([]*models.PesquisaResultadosProdutos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.PesquisaResultadosProdutos, err error)
	FindInBatches(result *[]*models.PesquisaResultadosProdutos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.PesquisaResultadosProdutos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPesquisaResultadosProdutosDo
	Assign(attrs ...field.AssignExpr) IPesquisaResultadosProdutosDo
	Joins(fields ...field.RelationField) IPesquisaResultadosProdutosDo
	Preload(fields ...field.RelationField) IPesquisaResultadosProdutosDo
	FirstOrInit() (*models.PesquisaResultadosProdutos, error)
	FirstOrCreate() (*models.PesquisaResultadosProdutos, error)
	FindByPage(offset int, limit int) (result []*models.PesquisaResultadosProdutos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPesquisaResultadosProdutosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pesquisaResultadosProdutosDo) Debug() IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Debug())
}

func (p pesquisaResultadosProdutosDo) WithContext(ctx context.Context) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pesquisaResultadosProdutosDo) ReadDB() IPesquisaResultadosProdutosDo {
	return p.Clauses(dbresolver.Read)
}

func (p pesquisaResultadosProdutosDo) WriteDB() IPesquisaResultadosProdutosDo {
	return p.Clauses(dbresolver.Write)
}

func (p pesquisaResultadosProdutosDo) Session(config *gorm.Session) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Session(config))
}

func (p pesquisaResultadosProdutosDo) Clauses(conds ...clause.Expression) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pesquisaResultadosProdutosDo) Returning(value interface{}, columns ...string) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pesquisaResultadosProdutosDo) Not(conds ...gen.Condition) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pesquisaResultadosProdutosDo) Or(conds ...gen.Condition) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pesquisaResultadosProdutosDo) Select(conds ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pesquisaResultadosProdutosDo) Where(conds ...gen.Condition) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pesquisaResultadosProdutosDo) Order(conds ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pesquisaResultadosProdutosDo) Distinct(cols ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pesquisaResultadosProdutosDo) Omit(cols ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pesquisaResultadosProdutosDo) Join(table schema.Tabler, on ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pesquisaResultadosProdutosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pesquisaResultadosProdutosDo) RightJoin(table schema.Tabler, on ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pesquisaResultadosProdutosDo) Group(cols ...field.Expr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pesquisaResultadosProdutosDo) Having(conds ...gen.Condition) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pesquisaResultadosProdutosDo) Limit(limit int) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pesquisaResultadosProdutosDo) Offset(offset int) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pesquisaResultadosProdutosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pesquisaResultadosProdutosDo) Unscoped() IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pesquisaResultadosProdutosDo) Create(values ...*models.PesquisaResultadosProdutos) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pesquisaResultadosProdutosDo) CreateInBatches(values []*models.PesquisaResultadosProdutos, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pesquisaResultadosProdutosDo) Save(values ...*models.PesquisaResultadosProdutos) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pesquisaResultadosProdutosDo) First() (*models.PesquisaResultadosProdutos, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisaResultadosProdutos), nil
	}
}

func (p pesquisaResultadosProdutosDo) Take() (*models.PesquisaResultadosProdutos, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisaResultadosProdutos), nil
	}
}

func (p pesquisaResultadosProdutosDo) Last() (*models.PesquisaResultadosProdutos, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisaResultadosProdutos), nil
	}
}

func (p pesquisaResultadosProdutosDo) Find() ([]*models.PesquisaResultadosProdutos, error) {
	result, err := p.DO.Find()
	return result.([]*models.PesquisaResultadosProdutos), err
}

func (p pesquisaResultadosProdutosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.PesquisaResultadosProdutos, err error) {
	buf := make([]*models.PesquisaResultadosProdutos, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pesquisaResultadosProdutosDo) FindInBatches(result *[]*models.PesquisaResultadosProdutos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pesquisaResultadosProdutosDo) Attrs(attrs ...field.AssignExpr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pesquisaResultadosProdutosDo) Assign(attrs ...field.AssignExpr) IPesquisaResultadosProdutosDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pesquisaResultadosProdutosDo) Joins(fields ...field.RelationField) IPesquisaResultadosProdutosDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pesquisaResultadosProdutosDo) Preload(fields ...field.RelationField) IPesquisaResultadosProdutosDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pesquisaResultadosProdutosDo) FirstOrInit() (*models.PesquisaResultadosProdutos, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisaResultadosProdutos), nil
	}
}

func (p pesquisaResultadosProdutosDo) FirstOrCreate() (*models.PesquisaResultadosProdutos, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisaResultadosProdutos), nil
	}
}

func (p pesquisaResultadosProdutosDo) FindByPage(offset int, limit int) (result []*models.PesquisaResultadosProdutos, count int64, err error) {
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

func (p pesquisaResultadosProdutosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pesquisaResultadosProdutosDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pesquisaResultadosProdutosDo) Delete(models ...*models.PesquisaResultadosProdutos) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pesquisaResultadosProdutosDo) withDO(do gen.Dao) *pesquisaResultadosProdutosDo {
	p.DO = *do.(*gen.DO)
	return p
}
