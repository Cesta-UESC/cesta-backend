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

func newPesquisas(db *gorm.DB, opts ...gen.DOOption) pesquisas {
	_pesquisas := pesquisas{}

	_pesquisas.pesquisasDo.UseDB(db, opts...)
	_pesquisas.pesquisasDo.UseModel(&model.Pesquisas{})

	tableName := _pesquisas.pesquisasDo.TableName()
	_pesquisas.ALL = field.NewAsterisk(tableName)
	_pesquisas.PesquisaID = field.NewInt32(tableName, "pesquisa_id")
	_pesquisas.DelimitadorID = field.NewInt32(tableName, "delimitador_id")
	_pesquisas.SalarioID = field.NewInt32(tableName, "salario_id")
	_pesquisas.PesquisaFechada = field.NewInt32(tableName, "pesquisa_fechada")
	_pesquisas.PesquisaData = field.NewTime(tableName, "pesquisa_data")
	_pesquisas.PesquisaDetalhada = field.NewInt32(tableName, "pesquisa_detalhada")

	_pesquisas.fillFieldMap()

	return _pesquisas
}

type pesquisas struct {
	pesquisasDo

	ALL               field.Asterisk
	PesquisaID        field.Int32
	DelimitadorID     field.Int32
	SalarioID         field.Int32
	PesquisaFechada   field.Int32
	PesquisaData      field.Time
	PesquisaDetalhada field.Int32

	fieldMap map[string]field.Expr
}

func (p pesquisas) Table(newTableName string) *pesquisas {
	p.pesquisasDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pesquisas) As(alias string) *pesquisas {
	p.pesquisasDo.DO = *(p.pesquisasDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pesquisas) updateTableName(table string) *pesquisas {
	p.ALL = field.NewAsterisk(table)
	p.PesquisaID = field.NewInt32(table, "pesquisa_id")
	p.DelimitadorID = field.NewInt32(table, "delimitador_id")
	p.SalarioID = field.NewInt32(table, "salario_id")
	p.PesquisaFechada = field.NewInt32(table, "pesquisa_fechada")
	p.PesquisaData = field.NewTime(table, "pesquisa_data")
	p.PesquisaDetalhada = field.NewInt32(table, "pesquisa_detalhada")

	p.fillFieldMap()

	return p
}

func (p *pesquisas) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pesquisas) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["pesquisa_id"] = p.PesquisaID
	p.fieldMap["delimitador_id"] = p.DelimitadorID
	p.fieldMap["salario_id"] = p.SalarioID
	p.fieldMap["pesquisa_fechada"] = p.PesquisaFechada
	p.fieldMap["pesquisa_data"] = p.PesquisaData
	p.fieldMap["pesquisa_detalhada"] = p.PesquisaDetalhada
}

func (p pesquisas) clone(db *gorm.DB) pesquisas {
	p.pesquisasDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pesquisas) replaceDB(db *gorm.DB) pesquisas {
	p.pesquisasDo.ReplaceDB(db)
	return p
}

type pesquisasDo struct{ gen.DO }

type IPesquisasDo interface {
	gen.SubQuery
	Debug() IPesquisasDo
	WithContext(ctx context.Context) IPesquisasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPesquisasDo
	WriteDB() IPesquisasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPesquisasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPesquisasDo
	Not(conds ...gen.Condition) IPesquisasDo
	Or(conds ...gen.Condition) IPesquisasDo
	Select(conds ...field.Expr) IPesquisasDo
	Where(conds ...gen.Condition) IPesquisasDo
	Order(conds ...field.Expr) IPesquisasDo
	Distinct(cols ...field.Expr) IPesquisasDo
	Omit(cols ...field.Expr) IPesquisasDo
	Join(table schema.Tabler, on ...field.Expr) IPesquisasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPesquisasDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPesquisasDo
	Group(cols ...field.Expr) IPesquisasDo
	Having(conds ...gen.Condition) IPesquisasDo
	Limit(limit int) IPesquisasDo
	Offset(offset int) IPesquisasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPesquisasDo
	Unscoped() IPesquisasDo
	Create(values ...*model.Pesquisas) error
	CreateInBatches(values []*model.Pesquisas, batchSize int) error
	Save(values ...*model.Pesquisas) error
	First() (*model.Pesquisas, error)
	Take() (*model.Pesquisas, error)
	Last() (*model.Pesquisas, error)
	Find() ([]*model.Pesquisas, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Pesquisas, err error)
	FindInBatches(result *[]*model.Pesquisas, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Pesquisas) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPesquisasDo
	Assign(attrs ...field.AssignExpr) IPesquisasDo
	Joins(fields ...field.RelationField) IPesquisasDo
	Preload(fields ...field.RelationField) IPesquisasDo
	FirstOrInit() (*model.Pesquisas, error)
	FirstOrCreate() (*model.Pesquisas, error)
	FindByPage(offset int, limit int) (result []*model.Pesquisas, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPesquisasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pesquisasDo) Debug() IPesquisasDo {
	return p.withDO(p.DO.Debug())
}

func (p pesquisasDo) WithContext(ctx context.Context) IPesquisasDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pesquisasDo) ReadDB() IPesquisasDo {
	return p.Clauses(dbresolver.Read)
}

func (p pesquisasDo) WriteDB() IPesquisasDo {
	return p.Clauses(dbresolver.Write)
}

func (p pesquisasDo) Session(config *gorm.Session) IPesquisasDo {
	return p.withDO(p.DO.Session(config))
}

func (p pesquisasDo) Clauses(conds ...clause.Expression) IPesquisasDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pesquisasDo) Returning(value interface{}, columns ...string) IPesquisasDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pesquisasDo) Not(conds ...gen.Condition) IPesquisasDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pesquisasDo) Or(conds ...gen.Condition) IPesquisasDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pesquisasDo) Select(conds ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pesquisasDo) Where(conds ...gen.Condition) IPesquisasDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pesquisasDo) Order(conds ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pesquisasDo) Distinct(cols ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pesquisasDo) Omit(cols ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pesquisasDo) Join(table schema.Tabler, on ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pesquisasDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pesquisasDo) RightJoin(table schema.Tabler, on ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pesquisasDo) Group(cols ...field.Expr) IPesquisasDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pesquisasDo) Having(conds ...gen.Condition) IPesquisasDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pesquisasDo) Limit(limit int) IPesquisasDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pesquisasDo) Offset(offset int) IPesquisasDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pesquisasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPesquisasDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pesquisasDo) Unscoped() IPesquisasDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pesquisasDo) Create(values ...*model.Pesquisas) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pesquisasDo) CreateInBatches(values []*model.Pesquisas, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pesquisasDo) Save(values ...*model.Pesquisas) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pesquisasDo) First() (*model.Pesquisas, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pesquisas), nil
	}
}

func (p pesquisasDo) Take() (*model.Pesquisas, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pesquisas), nil
	}
}

func (p pesquisasDo) Last() (*model.Pesquisas, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pesquisas), nil
	}
}

func (p pesquisasDo) Find() ([]*model.Pesquisas, error) {
	result, err := p.DO.Find()
	return result.([]*model.Pesquisas), err
}

func (p pesquisasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Pesquisas, err error) {
	buf := make([]*model.Pesquisas, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pesquisasDo) FindInBatches(result *[]*model.Pesquisas, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pesquisasDo) Attrs(attrs ...field.AssignExpr) IPesquisasDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pesquisasDo) Assign(attrs ...field.AssignExpr) IPesquisasDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pesquisasDo) Joins(fields ...field.RelationField) IPesquisasDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pesquisasDo) Preload(fields ...field.RelationField) IPesquisasDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pesquisasDo) FirstOrInit() (*model.Pesquisas, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pesquisas), nil
	}
}

func (p pesquisasDo) FirstOrCreate() (*model.Pesquisas, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pesquisas), nil
	}
}

func (p pesquisasDo) FindByPage(offset int, limit int) (result []*model.Pesquisas, count int64, err error) {
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

func (p pesquisasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pesquisasDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pesquisasDo) Delete(models ...*model.Pesquisas) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pesquisasDo) withDO(do gen.Dao) *pesquisasDo {
	p.DO = *do.(*gen.DO)
	return p
}
