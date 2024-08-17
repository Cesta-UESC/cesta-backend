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

func newEstabelecimentos(db *gorm.DB, opts ...gen.DOOption) estabelecimentos {
	_estabelecimentos := estabelecimentos{}

	_estabelecimentos.estabelecimentosDo.UseDB(db, opts...)
	_estabelecimentos.estabelecimentosDo.UseModel(&model.Estabelecimentos{})

	tableName := _estabelecimentos.estabelecimentosDo.TableName()
	_estabelecimentos.ALL = field.NewAsterisk(tableName)
	_estabelecimentos.EstabelecimentoID = field.NewInt32(tableName, "estabelecimento_id")
	_estabelecimentos.BairroID = field.NewInt32(tableName, "bairro_id")
	_estabelecimentos.EstabelecimentoNome = field.NewString(tableName, "estabelecimento_nome")
	_estabelecimentos.EstabelecimentoData = field.NewTime(tableName, "estabelecimento_data")
	_estabelecimentos.EstabelecimentoEndereco = field.NewString(tableName, "estabelecimento_endereco")
	_estabelecimentos.EstabelecimentoContato = field.NewString(tableName, "estabelecimento_contato")
	_estabelecimentos.EstabelecimentoTelefone = field.NewString(tableName, "estabelecimento_telefone")
	_estabelecimentos.EstabelecimentoReferencial = field.NewString(tableName, "estabelecimento_referencial")
	_estabelecimentos.EstabelecimentoAtivo = field.NewBool(tableName, "estabelecimento_ativo")

	_estabelecimentos.fillFieldMap()

	return _estabelecimentos
}

type estabelecimentos struct {
	estabelecimentosDo

	ALL                        field.Asterisk
	EstabelecimentoID          field.Int32
	BairroID                   field.Int32
	EstabelecimentoNome        field.String
	EstabelecimentoData        field.Time
	EstabelecimentoEndereco    field.String
	EstabelecimentoContato     field.String
	EstabelecimentoTelefone    field.String
	EstabelecimentoReferencial field.String
	EstabelecimentoAtivo       field.Bool

	fieldMap map[string]field.Expr
}

func (e estabelecimentos) Table(newTableName string) *estabelecimentos {
	e.estabelecimentosDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e estabelecimentos) As(alias string) *estabelecimentos {
	e.estabelecimentosDo.DO = *(e.estabelecimentosDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *estabelecimentos) updateTableName(table string) *estabelecimentos {
	e.ALL = field.NewAsterisk(table)
	e.EstabelecimentoID = field.NewInt32(table, "estabelecimento_id")
	e.BairroID = field.NewInt32(table, "bairro_id")
	e.EstabelecimentoNome = field.NewString(table, "estabelecimento_nome")
	e.EstabelecimentoData = field.NewTime(table, "estabelecimento_data")
	e.EstabelecimentoEndereco = field.NewString(table, "estabelecimento_endereco")
	e.EstabelecimentoContato = field.NewString(table, "estabelecimento_contato")
	e.EstabelecimentoTelefone = field.NewString(table, "estabelecimento_telefone")
	e.EstabelecimentoReferencial = field.NewString(table, "estabelecimento_referencial")
	e.EstabelecimentoAtivo = field.NewBool(table, "estabelecimento_ativo")

	e.fillFieldMap()

	return e
}

func (e *estabelecimentos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *estabelecimentos) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 9)
	e.fieldMap["estabelecimento_id"] = e.EstabelecimentoID
	e.fieldMap["bairro_id"] = e.BairroID
	e.fieldMap["estabelecimento_nome"] = e.EstabelecimentoNome
	e.fieldMap["estabelecimento_data"] = e.EstabelecimentoData
	e.fieldMap["estabelecimento_endereco"] = e.EstabelecimentoEndereco
	e.fieldMap["estabelecimento_contato"] = e.EstabelecimentoContato
	e.fieldMap["estabelecimento_telefone"] = e.EstabelecimentoTelefone
	e.fieldMap["estabelecimento_referencial"] = e.EstabelecimentoReferencial
	e.fieldMap["estabelecimento_ativo"] = e.EstabelecimentoAtivo
}

func (e estabelecimentos) clone(db *gorm.DB) estabelecimentos {
	e.estabelecimentosDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e estabelecimentos) replaceDB(db *gorm.DB) estabelecimentos {
	e.estabelecimentosDo.ReplaceDB(db)
	return e
}

type estabelecimentosDo struct{ gen.DO }

type IEstabelecimentosDo interface {
	gen.SubQuery
	Debug() IEstabelecimentosDo
	WithContext(ctx context.Context) IEstabelecimentosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEstabelecimentosDo
	WriteDB() IEstabelecimentosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEstabelecimentosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEstabelecimentosDo
	Not(conds ...gen.Condition) IEstabelecimentosDo
	Or(conds ...gen.Condition) IEstabelecimentosDo
	Select(conds ...field.Expr) IEstabelecimentosDo
	Where(conds ...gen.Condition) IEstabelecimentosDo
	Order(conds ...field.Expr) IEstabelecimentosDo
	Distinct(cols ...field.Expr) IEstabelecimentosDo
	Omit(cols ...field.Expr) IEstabelecimentosDo
	Join(table schema.Tabler, on ...field.Expr) IEstabelecimentosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosDo
	Group(cols ...field.Expr) IEstabelecimentosDo
	Having(conds ...gen.Condition) IEstabelecimentosDo
	Limit(limit int) IEstabelecimentosDo
	Offset(offset int) IEstabelecimentosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEstabelecimentosDo
	Unscoped() IEstabelecimentosDo
	Create(values ...*model.Estabelecimentos) error
	CreateInBatches(values []*model.Estabelecimentos, batchSize int) error
	Save(values ...*model.Estabelecimentos) error
	First() (*model.Estabelecimentos, error)
	Take() (*model.Estabelecimentos, error)
	Last() (*model.Estabelecimentos, error)
	Find() ([]*model.Estabelecimentos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Estabelecimentos, err error)
	FindInBatches(result *[]*model.Estabelecimentos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Estabelecimentos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEstabelecimentosDo
	Assign(attrs ...field.AssignExpr) IEstabelecimentosDo
	Joins(fields ...field.RelationField) IEstabelecimentosDo
	Preload(fields ...field.RelationField) IEstabelecimentosDo
	FirstOrInit() (*model.Estabelecimentos, error)
	FirstOrCreate() (*model.Estabelecimentos, error)
	FindByPage(offset int, limit int) (result []*model.Estabelecimentos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEstabelecimentosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e estabelecimentosDo) Debug() IEstabelecimentosDo {
	return e.withDO(e.DO.Debug())
}

func (e estabelecimentosDo) WithContext(ctx context.Context) IEstabelecimentosDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e estabelecimentosDo) ReadDB() IEstabelecimentosDo {
	return e.Clauses(dbresolver.Read)
}

func (e estabelecimentosDo) WriteDB() IEstabelecimentosDo {
	return e.Clauses(dbresolver.Write)
}

func (e estabelecimentosDo) Session(config *gorm.Session) IEstabelecimentosDo {
	return e.withDO(e.DO.Session(config))
}

func (e estabelecimentosDo) Clauses(conds ...clause.Expression) IEstabelecimentosDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e estabelecimentosDo) Returning(value interface{}, columns ...string) IEstabelecimentosDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e estabelecimentosDo) Not(conds ...gen.Condition) IEstabelecimentosDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e estabelecimentosDo) Or(conds ...gen.Condition) IEstabelecimentosDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e estabelecimentosDo) Select(conds ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e estabelecimentosDo) Where(conds ...gen.Condition) IEstabelecimentosDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e estabelecimentosDo) Order(conds ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e estabelecimentosDo) Distinct(cols ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e estabelecimentosDo) Omit(cols ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e estabelecimentosDo) Join(table schema.Tabler, on ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e estabelecimentosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e estabelecimentosDo) RightJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e estabelecimentosDo) Group(cols ...field.Expr) IEstabelecimentosDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e estabelecimentosDo) Having(conds ...gen.Condition) IEstabelecimentosDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e estabelecimentosDo) Limit(limit int) IEstabelecimentosDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e estabelecimentosDo) Offset(offset int) IEstabelecimentosDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e estabelecimentosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEstabelecimentosDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e estabelecimentosDo) Unscoped() IEstabelecimentosDo {
	return e.withDO(e.DO.Unscoped())
}

func (e estabelecimentosDo) Create(values ...*model.Estabelecimentos) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e estabelecimentosDo) CreateInBatches(values []*model.Estabelecimentos, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e estabelecimentosDo) Save(values ...*model.Estabelecimentos) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e estabelecimentosDo) First() (*model.Estabelecimentos, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Estabelecimentos), nil
	}
}

func (e estabelecimentosDo) Take() (*model.Estabelecimentos, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Estabelecimentos), nil
	}
}

func (e estabelecimentosDo) Last() (*model.Estabelecimentos, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Estabelecimentos), nil
	}
}

func (e estabelecimentosDo) Find() ([]*model.Estabelecimentos, error) {
	result, err := e.DO.Find()
	return result.([]*model.Estabelecimentos), err
}

func (e estabelecimentosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Estabelecimentos, err error) {
	buf := make([]*model.Estabelecimentos, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e estabelecimentosDo) FindInBatches(result *[]*model.Estabelecimentos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e estabelecimentosDo) Attrs(attrs ...field.AssignExpr) IEstabelecimentosDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e estabelecimentosDo) Assign(attrs ...field.AssignExpr) IEstabelecimentosDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e estabelecimentosDo) Joins(fields ...field.RelationField) IEstabelecimentosDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e estabelecimentosDo) Preload(fields ...field.RelationField) IEstabelecimentosDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e estabelecimentosDo) FirstOrInit() (*model.Estabelecimentos, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Estabelecimentos), nil
	}
}

func (e estabelecimentosDo) FirstOrCreate() (*model.Estabelecimentos, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Estabelecimentos), nil
	}
}

func (e estabelecimentosDo) FindByPage(offset int, limit int) (result []*model.Estabelecimentos, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e estabelecimentosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e estabelecimentosDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e estabelecimentosDo) Delete(models ...*model.Estabelecimentos) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *estabelecimentosDo) withDO(do gen.Dao) *estabelecimentosDo {
	e.DO = *do.(*gen.DO)
	return e
}
