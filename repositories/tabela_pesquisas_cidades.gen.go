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

func newPesquisasCidades(db *gorm.DB, opts ...gen.DOOption) pesquisasCidades {
	_pesquisasCidades := pesquisasCidades{}

	_pesquisasCidades.pesquisasCidadesDo.UseDB(db, opts...)
	_pesquisasCidades.pesquisasCidadesDo.UseModel(&models.PesquisasCidades{})

	tableName := _pesquisasCidades.pesquisasCidadesDo.TableName()
	_pesquisasCidades.ALL = field.NewAsterisk(tableName)
	_pesquisasCidades.PesquisaID = field.NewInt32(tableName, "pesquisa_id")
	_pesquisasCidades.CidadeID = field.NewInt32(tableName, "cidade_id")
	_pesquisasCidades.VariacaoMensal = field.NewFloat64(tableName, "variacao_mensal")
	_pesquisasCidades.VariacaoSemestral = field.NewFloat64(tableName, "variacao_semestral")
	_pesquisasCidades.VariacaoAnual = field.NewFloat64(tableName, "variacao_anual")
	_pesquisasCidades.GastoMensalCesta = field.NewFloat64(tableName, "gasto_mensal_cesta")
	_pesquisasCidades.TempoTrabalho = field.NewFloat64(tableName, "tempo_trabalho")

	_pesquisasCidades.fillFieldMap()

	return _pesquisasCidades
}

type pesquisasCidades struct {
	pesquisasCidadesDo

	ALL               field.Asterisk
	PesquisaID        field.Int32
	CidadeID          field.Int32
	VariacaoMensal    field.Float64
	VariacaoSemestral field.Float64
	VariacaoAnual     field.Float64
	GastoMensalCesta  field.Float64
	TempoTrabalho     field.Float64

	fieldMap map[string]field.Expr
}

func (p pesquisasCidades) Table(newTableName string) *pesquisasCidades {
	p.pesquisasCidadesDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pesquisasCidades) As(alias string) *pesquisasCidades {
	p.pesquisasCidadesDo.DO = *(p.pesquisasCidadesDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pesquisasCidades) updateTableName(table string) *pesquisasCidades {
	p.ALL = field.NewAsterisk(table)
	p.PesquisaID = field.NewInt32(table, "pesquisa_id")
	p.CidadeID = field.NewInt32(table, "cidade_id")
	p.VariacaoMensal = field.NewFloat64(table, "variacao_mensal")
	p.VariacaoSemestral = field.NewFloat64(table, "variacao_semestral")
	p.VariacaoAnual = field.NewFloat64(table, "variacao_anual")
	p.GastoMensalCesta = field.NewFloat64(table, "gasto_mensal_cesta")
	p.TempoTrabalho = field.NewFloat64(table, "tempo_trabalho")

	p.fillFieldMap()

	return p
}

func (p *pesquisasCidades) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pesquisasCidades) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["pesquisa_id"] = p.PesquisaID
	p.fieldMap["cidade_id"] = p.CidadeID
	p.fieldMap["variacao_mensal"] = p.VariacaoMensal
	p.fieldMap["variacao_semestral"] = p.VariacaoSemestral
	p.fieldMap["variacao_anual"] = p.VariacaoAnual
	p.fieldMap["gasto_mensal_cesta"] = p.GastoMensalCesta
	p.fieldMap["tempo_trabalho"] = p.TempoTrabalho
}

func (p pesquisasCidades) clone(db *gorm.DB) pesquisasCidades {
	p.pesquisasCidadesDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pesquisasCidades) replaceDB(db *gorm.DB) pesquisasCidades {
	p.pesquisasCidadesDo.ReplaceDB(db)
	return p
}

type pesquisasCidadesDo struct{ gen.DO }

type IPesquisasCidadesDo interface {
	gen.SubQuery
	Debug() IPesquisasCidadesDo
	WithContext(ctx context.Context) IPesquisasCidadesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPesquisasCidadesDo
	WriteDB() IPesquisasCidadesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPesquisasCidadesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPesquisasCidadesDo
	Not(conds ...gen.Condition) IPesquisasCidadesDo
	Or(conds ...gen.Condition) IPesquisasCidadesDo
	Select(conds ...field.Expr) IPesquisasCidadesDo
	Where(conds ...gen.Condition) IPesquisasCidadesDo
	Order(conds ...field.Expr) IPesquisasCidadesDo
	Distinct(cols ...field.Expr) IPesquisasCidadesDo
	Omit(cols ...field.Expr) IPesquisasCidadesDo
	Join(table schema.Tabler, on ...field.Expr) IPesquisasCidadesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPesquisasCidadesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPesquisasCidadesDo
	Group(cols ...field.Expr) IPesquisasCidadesDo
	Having(conds ...gen.Condition) IPesquisasCidadesDo
	Limit(limit int) IPesquisasCidadesDo
	Offset(offset int) IPesquisasCidadesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPesquisasCidadesDo
	Unscoped() IPesquisasCidadesDo
	Create(values ...*models.PesquisasCidades) error
	CreateInBatches(values []*models.PesquisasCidades, batchSize int) error
	Save(values ...*models.PesquisasCidades) error
	First() (*models.PesquisasCidades, error)
	Take() (*models.PesquisasCidades, error)
	Last() (*models.PesquisasCidades, error)
	Find() ([]*models.PesquisasCidades, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.PesquisasCidades, err error)
	FindInBatches(result *[]*models.PesquisasCidades, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.PesquisasCidades) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPesquisasCidadesDo
	Assign(attrs ...field.AssignExpr) IPesquisasCidadesDo
	Joins(fields ...field.RelationField) IPesquisasCidadesDo
	Preload(fields ...field.RelationField) IPesquisasCidadesDo
	FirstOrInit() (*models.PesquisasCidades, error)
	FirstOrCreate() (*models.PesquisasCidades, error)
	FindByPage(offset int, limit int) (result []*models.PesquisasCidades, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPesquisasCidadesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pesquisasCidadesDo) Debug() IPesquisasCidadesDo {
	return p.withDO(p.DO.Debug())
}

func (p pesquisasCidadesDo) WithContext(ctx context.Context) IPesquisasCidadesDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pesquisasCidadesDo) ReadDB() IPesquisasCidadesDo {
	return p.Clauses(dbresolver.Read)
}

func (p pesquisasCidadesDo) WriteDB() IPesquisasCidadesDo {
	return p.Clauses(dbresolver.Write)
}

func (p pesquisasCidadesDo) Session(config *gorm.Session) IPesquisasCidadesDo {
	return p.withDO(p.DO.Session(config))
}

func (p pesquisasCidadesDo) Clauses(conds ...clause.Expression) IPesquisasCidadesDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pesquisasCidadesDo) Returning(value interface{}, columns ...string) IPesquisasCidadesDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pesquisasCidadesDo) Not(conds ...gen.Condition) IPesquisasCidadesDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pesquisasCidadesDo) Or(conds ...gen.Condition) IPesquisasCidadesDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pesquisasCidadesDo) Select(conds ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pesquisasCidadesDo) Where(conds ...gen.Condition) IPesquisasCidadesDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pesquisasCidadesDo) Order(conds ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pesquisasCidadesDo) Distinct(cols ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pesquisasCidadesDo) Omit(cols ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pesquisasCidadesDo) Join(table schema.Tabler, on ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pesquisasCidadesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pesquisasCidadesDo) RightJoin(table schema.Tabler, on ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pesquisasCidadesDo) Group(cols ...field.Expr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pesquisasCidadesDo) Having(conds ...gen.Condition) IPesquisasCidadesDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pesquisasCidadesDo) Limit(limit int) IPesquisasCidadesDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pesquisasCidadesDo) Offset(offset int) IPesquisasCidadesDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pesquisasCidadesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPesquisasCidadesDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pesquisasCidadesDo) Unscoped() IPesquisasCidadesDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pesquisasCidadesDo) Create(values ...*models.PesquisasCidades) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pesquisasCidadesDo) CreateInBatches(values []*models.PesquisasCidades, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pesquisasCidadesDo) Save(values ...*models.PesquisasCidades) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pesquisasCidadesDo) First() (*models.PesquisasCidades, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisasCidades), nil
	}
}

func (p pesquisasCidadesDo) Take() (*models.PesquisasCidades, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisasCidades), nil
	}
}

func (p pesquisasCidadesDo) Last() (*models.PesquisasCidades, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisasCidades), nil
	}
}

func (p pesquisasCidadesDo) Find() ([]*models.PesquisasCidades, error) {
	result, err := p.DO.Find()
	return result.([]*models.PesquisasCidades), err
}

func (p pesquisasCidadesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.PesquisasCidades, err error) {
	buf := make([]*models.PesquisasCidades, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pesquisasCidadesDo) FindInBatches(result *[]*models.PesquisasCidades, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pesquisasCidadesDo) Attrs(attrs ...field.AssignExpr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pesquisasCidadesDo) Assign(attrs ...field.AssignExpr) IPesquisasCidadesDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pesquisasCidadesDo) Joins(fields ...field.RelationField) IPesquisasCidadesDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pesquisasCidadesDo) Preload(fields ...field.RelationField) IPesquisasCidadesDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pesquisasCidadesDo) FirstOrInit() (*models.PesquisasCidades, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisasCidades), nil
	}
}

func (p pesquisasCidadesDo) FirstOrCreate() (*models.PesquisasCidades, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.PesquisasCidades), nil
	}
}

func (p pesquisasCidadesDo) FindByPage(offset int, limit int) (result []*models.PesquisasCidades, count int64, err error) {
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

func (p pesquisasCidadesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pesquisasCidadesDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pesquisasCidadesDo) Delete(models ...*models.PesquisasCidades) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pesquisasCidadesDo) withDO(do gen.Dao) *pesquisasCidadesDo {
	p.DO = *do.(*gen.DO)
	return p
}
