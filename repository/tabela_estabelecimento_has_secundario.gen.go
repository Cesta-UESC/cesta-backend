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

func newEstabelecimentoHasSecundario(db *gorm.DB, opts ...gen.DOOption) estabelecimentoHasSecundario {
	_estabelecimentoHasSecundario := estabelecimentoHasSecundario{}

	_estabelecimentoHasSecundario.estabelecimentoHasSecundarioDo.UseDB(db, opts...)
	_estabelecimentoHasSecundario.estabelecimentoHasSecundarioDo.UseModel(&model.EstabelecimentoHasSecundario{})

	tableName := _estabelecimentoHasSecundario.estabelecimentoHasSecundarioDo.TableName()
	_estabelecimentoHasSecundario.ALL = field.NewAsterisk(tableName)
	_estabelecimentoHasSecundario.EstHasSecID = field.NewInt32(tableName, "est_has_sec_id")
	_estabelecimentoHasSecundario.EstabelecimentoID = field.NewInt32(tableName, "estabelecimento_id")
	_estabelecimentoHasSecundario.EstabelecimentoSecID = field.NewInt32(tableName, "estabelecimento_sec_id")

	_estabelecimentoHasSecundario.fillFieldMap()

	return _estabelecimentoHasSecundario
}

type estabelecimentoHasSecundario struct {
	estabelecimentoHasSecundarioDo

	ALL                  field.Asterisk
	EstHasSecID          field.Int32
	EstabelecimentoID    field.Int32
	EstabelecimentoSecID field.Int32

	fieldMap map[string]field.Expr
}

func (e estabelecimentoHasSecundario) Table(newTableName string) *estabelecimentoHasSecundario {
	e.estabelecimentoHasSecundarioDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e estabelecimentoHasSecundario) As(alias string) *estabelecimentoHasSecundario {
	e.estabelecimentoHasSecundarioDo.DO = *(e.estabelecimentoHasSecundarioDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *estabelecimentoHasSecundario) updateTableName(table string) *estabelecimentoHasSecundario {
	e.ALL = field.NewAsterisk(table)
	e.EstHasSecID = field.NewInt32(table, "est_has_sec_id")
	e.EstabelecimentoID = field.NewInt32(table, "estabelecimento_id")
	e.EstabelecimentoSecID = field.NewInt32(table, "estabelecimento_sec_id")

	e.fillFieldMap()

	return e
}

func (e *estabelecimentoHasSecundario) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *estabelecimentoHasSecundario) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 3)
	e.fieldMap["est_has_sec_id"] = e.EstHasSecID
	e.fieldMap["estabelecimento_id"] = e.EstabelecimentoID
	e.fieldMap["estabelecimento_sec_id"] = e.EstabelecimentoSecID
}

func (e estabelecimentoHasSecundario) clone(db *gorm.DB) estabelecimentoHasSecundario {
	e.estabelecimentoHasSecundarioDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e estabelecimentoHasSecundario) replaceDB(db *gorm.DB) estabelecimentoHasSecundario {
	e.estabelecimentoHasSecundarioDo.ReplaceDB(db)
	return e
}

type estabelecimentoHasSecundarioDo struct{ gen.DO }

type IEstabelecimentoHasSecundarioDo interface {
	gen.SubQuery
	Debug() IEstabelecimentoHasSecundarioDo
	WithContext(ctx context.Context) IEstabelecimentoHasSecundarioDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEstabelecimentoHasSecundarioDo
	WriteDB() IEstabelecimentoHasSecundarioDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEstabelecimentoHasSecundarioDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEstabelecimentoHasSecundarioDo
	Not(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo
	Or(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo
	Select(conds ...field.Expr) IEstabelecimentoHasSecundarioDo
	Where(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo
	Order(conds ...field.Expr) IEstabelecimentoHasSecundarioDo
	Distinct(cols ...field.Expr) IEstabelecimentoHasSecundarioDo
	Omit(cols ...field.Expr) IEstabelecimentoHasSecundarioDo
	Join(table schema.Tabler, on ...field.Expr) IEstabelecimentoHasSecundarioDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentoHasSecundarioDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentoHasSecundarioDo
	Group(cols ...field.Expr) IEstabelecimentoHasSecundarioDo
	Having(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo
	Limit(limit int) IEstabelecimentoHasSecundarioDo
	Offset(offset int) IEstabelecimentoHasSecundarioDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEstabelecimentoHasSecundarioDo
	Unscoped() IEstabelecimentoHasSecundarioDo
	Create(values ...*model.EstabelecimentoHasSecundario) error
	CreateInBatches(values []*model.EstabelecimentoHasSecundario, batchSize int) error
	Save(values ...*model.EstabelecimentoHasSecundario) error
	First() (*model.EstabelecimentoHasSecundario, error)
	Take() (*model.EstabelecimentoHasSecundario, error)
	Last() (*model.EstabelecimentoHasSecundario, error)
	Find() ([]*model.EstabelecimentoHasSecundario, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EstabelecimentoHasSecundario, err error)
	FindInBatches(result *[]*model.EstabelecimentoHasSecundario, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EstabelecimentoHasSecundario) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEstabelecimentoHasSecundarioDo
	Assign(attrs ...field.AssignExpr) IEstabelecimentoHasSecundarioDo
	Joins(fields ...field.RelationField) IEstabelecimentoHasSecundarioDo
	Preload(fields ...field.RelationField) IEstabelecimentoHasSecundarioDo
	FirstOrInit() (*model.EstabelecimentoHasSecundario, error)
	FirstOrCreate() (*model.EstabelecimentoHasSecundario, error)
	FindByPage(offset int, limit int) (result []*model.EstabelecimentoHasSecundario, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEstabelecimentoHasSecundarioDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e estabelecimentoHasSecundarioDo) Debug() IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Debug())
}

func (e estabelecimentoHasSecundarioDo) WithContext(ctx context.Context) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e estabelecimentoHasSecundarioDo) ReadDB() IEstabelecimentoHasSecundarioDo {
	return e.Clauses(dbresolver.Read)
}

func (e estabelecimentoHasSecundarioDo) WriteDB() IEstabelecimentoHasSecundarioDo {
	return e.Clauses(dbresolver.Write)
}

func (e estabelecimentoHasSecundarioDo) Session(config *gorm.Session) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Session(config))
}

func (e estabelecimentoHasSecundarioDo) Clauses(conds ...clause.Expression) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e estabelecimentoHasSecundarioDo) Returning(value interface{}, columns ...string) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e estabelecimentoHasSecundarioDo) Not(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e estabelecimentoHasSecundarioDo) Or(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e estabelecimentoHasSecundarioDo) Select(conds ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e estabelecimentoHasSecundarioDo) Where(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e estabelecimentoHasSecundarioDo) Order(conds ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e estabelecimentoHasSecundarioDo) Distinct(cols ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e estabelecimentoHasSecundarioDo) Omit(cols ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e estabelecimentoHasSecundarioDo) Join(table schema.Tabler, on ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e estabelecimentoHasSecundarioDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e estabelecimentoHasSecundarioDo) RightJoin(table schema.Tabler, on ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e estabelecimentoHasSecundarioDo) Group(cols ...field.Expr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e estabelecimentoHasSecundarioDo) Having(conds ...gen.Condition) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e estabelecimentoHasSecundarioDo) Limit(limit int) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e estabelecimentoHasSecundarioDo) Offset(offset int) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e estabelecimentoHasSecundarioDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e estabelecimentoHasSecundarioDo) Unscoped() IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Unscoped())
}

func (e estabelecimentoHasSecundarioDo) Create(values ...*model.EstabelecimentoHasSecundario) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e estabelecimentoHasSecundarioDo) CreateInBatches(values []*model.EstabelecimentoHasSecundario, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e estabelecimentoHasSecundarioDo) Save(values ...*model.EstabelecimentoHasSecundario) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e estabelecimentoHasSecundarioDo) First() (*model.EstabelecimentoHasSecundario, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstabelecimentoHasSecundario), nil
	}
}

func (e estabelecimentoHasSecundarioDo) Take() (*model.EstabelecimentoHasSecundario, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstabelecimentoHasSecundario), nil
	}
}

func (e estabelecimentoHasSecundarioDo) Last() (*model.EstabelecimentoHasSecundario, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstabelecimentoHasSecundario), nil
	}
}

func (e estabelecimentoHasSecundarioDo) Find() ([]*model.EstabelecimentoHasSecundario, error) {
	result, err := e.DO.Find()
	return result.([]*model.EstabelecimentoHasSecundario), err
}

func (e estabelecimentoHasSecundarioDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EstabelecimentoHasSecundario, err error) {
	buf := make([]*model.EstabelecimentoHasSecundario, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e estabelecimentoHasSecundarioDo) FindInBatches(result *[]*model.EstabelecimentoHasSecundario, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e estabelecimentoHasSecundarioDo) Attrs(attrs ...field.AssignExpr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e estabelecimentoHasSecundarioDo) Assign(attrs ...field.AssignExpr) IEstabelecimentoHasSecundarioDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e estabelecimentoHasSecundarioDo) Joins(fields ...field.RelationField) IEstabelecimentoHasSecundarioDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e estabelecimentoHasSecundarioDo) Preload(fields ...field.RelationField) IEstabelecimentoHasSecundarioDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e estabelecimentoHasSecundarioDo) FirstOrInit() (*model.EstabelecimentoHasSecundario, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstabelecimentoHasSecundario), nil
	}
}

func (e estabelecimentoHasSecundarioDo) FirstOrCreate() (*model.EstabelecimentoHasSecundario, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstabelecimentoHasSecundario), nil
	}
}

func (e estabelecimentoHasSecundarioDo) FindByPage(offset int, limit int) (result []*model.EstabelecimentoHasSecundario, count int64, err error) {
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

func (e estabelecimentoHasSecundarioDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e estabelecimentoHasSecundarioDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e estabelecimentoHasSecundarioDo) Delete(models ...*model.EstabelecimentoHasSecundario) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *estabelecimentoHasSecundarioDo) withDO(do gen.Dao) *estabelecimentoHasSecundarioDo {
	e.DO = *do.(*gen.DO)
	return e
}
