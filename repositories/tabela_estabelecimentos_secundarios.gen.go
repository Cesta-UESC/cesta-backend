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

func newEstabelecimentosSecundarios(db *gorm.DB, opts ...gen.DOOption) estabelecimentosSecundarios {
	_estabelecimentosSecundarios := estabelecimentosSecundarios{}

	_estabelecimentosSecundarios.estabelecimentosSecundariosDo.UseDB(db, opts...)
	_estabelecimentosSecundarios.estabelecimentosSecundariosDo.UseModel(&models.EstabelecimentosSecundarios{})

	tableName := _estabelecimentosSecundarios.estabelecimentosSecundariosDo.TableName()
	_estabelecimentosSecundarios.ALL = field.NewAsterisk(tableName)
	_estabelecimentosSecundarios.EstabelecimentoSecID = field.NewInt32(tableName, "estabelecimento_sec_id")
	_estabelecimentosSecundarios.EstabelecimentoSecNome = field.NewString(tableName, "estabelecimento_sec_nome")

	_estabelecimentosSecundarios.fillFieldMap()

	return _estabelecimentosSecundarios
}

type estabelecimentosSecundarios struct {
	estabelecimentosSecundariosDo

	ALL                    field.Asterisk
	EstabelecimentoSecID   field.Int32
	EstabelecimentoSecNome field.String

	fieldMap map[string]field.Expr
}

func (e estabelecimentosSecundarios) Table(newTableName string) *estabelecimentosSecundarios {
	e.estabelecimentosSecundariosDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e estabelecimentosSecundarios) As(alias string) *estabelecimentosSecundarios {
	e.estabelecimentosSecundariosDo.DO = *(e.estabelecimentosSecundariosDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *estabelecimentosSecundarios) updateTableName(table string) *estabelecimentosSecundarios {
	e.ALL = field.NewAsterisk(table)
	e.EstabelecimentoSecID = field.NewInt32(table, "estabelecimento_sec_id")
	e.EstabelecimentoSecNome = field.NewString(table, "estabelecimento_sec_nome")

	e.fillFieldMap()

	return e
}

func (e *estabelecimentosSecundarios) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *estabelecimentosSecundarios) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 2)
	e.fieldMap["estabelecimento_sec_id"] = e.EstabelecimentoSecID
	e.fieldMap["estabelecimento_sec_nome"] = e.EstabelecimentoSecNome
}

func (e estabelecimentosSecundarios) clone(db *gorm.DB) estabelecimentosSecundarios {
	e.estabelecimentosSecundariosDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e estabelecimentosSecundarios) replaceDB(db *gorm.DB) estabelecimentosSecundarios {
	e.estabelecimentosSecundariosDo.ReplaceDB(db)
	return e
}

type estabelecimentosSecundariosDo struct{ gen.DO }

type IEstabelecimentosSecundariosDo interface {
	gen.SubQuery
	Debug() IEstabelecimentosSecundariosDo
	WithContext(ctx context.Context) IEstabelecimentosSecundariosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEstabelecimentosSecundariosDo
	WriteDB() IEstabelecimentosSecundariosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEstabelecimentosSecundariosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEstabelecimentosSecundariosDo
	Not(conds ...gen.Condition) IEstabelecimentosSecundariosDo
	Or(conds ...gen.Condition) IEstabelecimentosSecundariosDo
	Select(conds ...field.Expr) IEstabelecimentosSecundariosDo
	Where(conds ...gen.Condition) IEstabelecimentosSecundariosDo
	Order(conds ...field.Expr) IEstabelecimentosSecundariosDo
	Distinct(cols ...field.Expr) IEstabelecimentosSecundariosDo
	Omit(cols ...field.Expr) IEstabelecimentosSecundariosDo
	Join(table schema.Tabler, on ...field.Expr) IEstabelecimentosSecundariosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosSecundariosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosSecundariosDo
	Group(cols ...field.Expr) IEstabelecimentosSecundariosDo
	Having(conds ...gen.Condition) IEstabelecimentosSecundariosDo
	Limit(limit int) IEstabelecimentosSecundariosDo
	Offset(offset int) IEstabelecimentosSecundariosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEstabelecimentosSecundariosDo
	Unscoped() IEstabelecimentosSecundariosDo
	Create(values ...*models.EstabelecimentosSecundarios) error
	CreateInBatches(values []*models.EstabelecimentosSecundarios, batchSize int) error
	Save(values ...*models.EstabelecimentosSecundarios) error
	First() (*models.EstabelecimentosSecundarios, error)
	Take() (*models.EstabelecimentosSecundarios, error)
	Last() (*models.EstabelecimentosSecundarios, error)
	Find() ([]*models.EstabelecimentosSecundarios, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.EstabelecimentosSecundarios, err error)
	FindInBatches(result *[]*models.EstabelecimentosSecundarios, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.EstabelecimentosSecundarios) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEstabelecimentosSecundariosDo
	Assign(attrs ...field.AssignExpr) IEstabelecimentosSecundariosDo
	Joins(fields ...field.RelationField) IEstabelecimentosSecundariosDo
	Preload(fields ...field.RelationField) IEstabelecimentosSecundariosDo
	FirstOrInit() (*models.EstabelecimentosSecundarios, error)
	FirstOrCreate() (*models.EstabelecimentosSecundarios, error)
	FindByPage(offset int, limit int) (result []*models.EstabelecimentosSecundarios, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEstabelecimentosSecundariosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e estabelecimentosSecundariosDo) Debug() IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Debug())
}

func (e estabelecimentosSecundariosDo) WithContext(ctx context.Context) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e estabelecimentosSecundariosDo) ReadDB() IEstabelecimentosSecundariosDo {
	return e.Clauses(dbresolver.Read)
}

func (e estabelecimentosSecundariosDo) WriteDB() IEstabelecimentosSecundariosDo {
	return e.Clauses(dbresolver.Write)
}

func (e estabelecimentosSecundariosDo) Session(config *gorm.Session) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Session(config))
}

func (e estabelecimentosSecundariosDo) Clauses(conds ...clause.Expression) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e estabelecimentosSecundariosDo) Returning(value interface{}, columns ...string) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e estabelecimentosSecundariosDo) Not(conds ...gen.Condition) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e estabelecimentosSecundariosDo) Or(conds ...gen.Condition) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e estabelecimentosSecundariosDo) Select(conds ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e estabelecimentosSecundariosDo) Where(conds ...gen.Condition) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e estabelecimentosSecundariosDo) Order(conds ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e estabelecimentosSecundariosDo) Distinct(cols ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e estabelecimentosSecundariosDo) Omit(cols ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e estabelecimentosSecundariosDo) Join(table schema.Tabler, on ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e estabelecimentosSecundariosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e estabelecimentosSecundariosDo) RightJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e estabelecimentosSecundariosDo) Group(cols ...field.Expr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e estabelecimentosSecundariosDo) Having(conds ...gen.Condition) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e estabelecimentosSecundariosDo) Limit(limit int) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e estabelecimentosSecundariosDo) Offset(offset int) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e estabelecimentosSecundariosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e estabelecimentosSecundariosDo) Unscoped() IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Unscoped())
}

func (e estabelecimentosSecundariosDo) Create(values ...*models.EstabelecimentosSecundarios) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e estabelecimentosSecundariosDo) CreateInBatches(values []*models.EstabelecimentosSecundarios, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e estabelecimentosSecundariosDo) Save(values ...*models.EstabelecimentosSecundarios) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e estabelecimentosSecundariosDo) First() (*models.EstabelecimentosSecundarios, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.EstabelecimentosSecundarios), nil
	}
}

func (e estabelecimentosSecundariosDo) Take() (*models.EstabelecimentosSecundarios, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.EstabelecimentosSecundarios), nil
	}
}

func (e estabelecimentosSecundariosDo) Last() (*models.EstabelecimentosSecundarios, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.EstabelecimentosSecundarios), nil
	}
}

func (e estabelecimentosSecundariosDo) Find() ([]*models.EstabelecimentosSecundarios, error) {
	result, err := e.DO.Find()
	return result.([]*models.EstabelecimentosSecundarios), err
}

func (e estabelecimentosSecundariosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.EstabelecimentosSecundarios, err error) {
	buf := make([]*models.EstabelecimentosSecundarios, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e estabelecimentosSecundariosDo) FindInBatches(result *[]*models.EstabelecimentosSecundarios, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e estabelecimentosSecundariosDo) Attrs(attrs ...field.AssignExpr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e estabelecimentosSecundariosDo) Assign(attrs ...field.AssignExpr) IEstabelecimentosSecundariosDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e estabelecimentosSecundariosDo) Joins(fields ...field.RelationField) IEstabelecimentosSecundariosDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e estabelecimentosSecundariosDo) Preload(fields ...field.RelationField) IEstabelecimentosSecundariosDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e estabelecimentosSecundariosDo) FirstOrInit() (*models.EstabelecimentosSecundarios, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.EstabelecimentosSecundarios), nil
	}
}

func (e estabelecimentosSecundariosDo) FirstOrCreate() (*models.EstabelecimentosSecundarios, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.EstabelecimentosSecundarios), nil
	}
}

func (e estabelecimentosSecundariosDo) FindByPage(offset int, limit int) (result []*models.EstabelecimentosSecundarios, count int64, err error) {
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

func (e estabelecimentosSecundariosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e estabelecimentosSecundariosDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e estabelecimentosSecundariosDo) Delete(models ...*models.EstabelecimentosSecundarios) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *estabelecimentosSecundariosDo) withDO(do gen.Dao) *estabelecimentosSecundariosDo {
	e.DO = *do.(*gen.DO)
	return e
}
