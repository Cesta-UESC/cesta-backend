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

func newRacaoMinima(db *gorm.DB, opts ...gen.DOOption) racaoMinima {
	_racaoMinima := racaoMinima{}

	_racaoMinima.racaoMinimaDo.UseDB(db, opts...)
	_racaoMinima.racaoMinimaDo.UseModel(&model.RacaoMinima{})

	tableName := _racaoMinima.racaoMinimaDo.TableName()
	_racaoMinima.ALL = field.NewAsterisk(tableName)
	_racaoMinima.DelimitadorID = field.NewInt32(tableName, "delimitador_id")
	_racaoMinima.ProdutoID = field.NewInt32(tableName, "produto_id")
	_racaoMinima.MedidaID = field.NewInt32(tableName, "medida_id")
	_racaoMinima.RacaoMinimaQuantidade = field.NewFloat64(tableName, "racao_minima_quantidade")
	_racaoMinima.RacaoMinimaTransformador = field.NewFloat64(tableName, "racao_minima_transformador")
	_racaoMinima.RacaoMinimaMedida = field.NewInt32(tableName, "racao_minima_medida")

	_racaoMinima.fillFieldMap()

	return _racaoMinima
}

type racaoMinima struct {
	racaoMinimaDo

	ALL                      field.Asterisk
	DelimitadorID            field.Int32
	ProdutoID                field.Int32
	MedidaID                 field.Int32
	RacaoMinimaQuantidade    field.Float64
	RacaoMinimaTransformador field.Float64
	RacaoMinimaMedida        field.Int32

	fieldMap map[string]field.Expr
}

func (r racaoMinima) Table(newTableName string) *racaoMinima {
	r.racaoMinimaDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r racaoMinima) As(alias string) *racaoMinima {
	r.racaoMinimaDo.DO = *(r.racaoMinimaDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *racaoMinima) updateTableName(table string) *racaoMinima {
	r.ALL = field.NewAsterisk(table)
	r.DelimitadorID = field.NewInt32(table, "delimitador_id")
	r.ProdutoID = field.NewInt32(table, "produto_id")
	r.MedidaID = field.NewInt32(table, "medida_id")
	r.RacaoMinimaQuantidade = field.NewFloat64(table, "racao_minima_quantidade")
	r.RacaoMinimaTransformador = field.NewFloat64(table, "racao_minima_transformador")
	r.RacaoMinimaMedida = field.NewInt32(table, "racao_minima_medida")

	r.fillFieldMap()

	return r
}

func (r *racaoMinima) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *racaoMinima) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["delimitador_id"] = r.DelimitadorID
	r.fieldMap["produto_id"] = r.ProdutoID
	r.fieldMap["medida_id"] = r.MedidaID
	r.fieldMap["racao_minima_quantidade"] = r.RacaoMinimaQuantidade
	r.fieldMap["racao_minima_transformador"] = r.RacaoMinimaTransformador
	r.fieldMap["racao_minima_medida"] = r.RacaoMinimaMedida
}

func (r racaoMinima) clone(db *gorm.DB) racaoMinima {
	r.racaoMinimaDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r racaoMinima) replaceDB(db *gorm.DB) racaoMinima {
	r.racaoMinimaDo.ReplaceDB(db)
	return r
}

type racaoMinimaDo struct{ gen.DO }

type IRacaoMinimaDo interface {
	gen.SubQuery
	Debug() IRacaoMinimaDo
	WithContext(ctx context.Context) IRacaoMinimaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRacaoMinimaDo
	WriteDB() IRacaoMinimaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRacaoMinimaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRacaoMinimaDo
	Not(conds ...gen.Condition) IRacaoMinimaDo
	Or(conds ...gen.Condition) IRacaoMinimaDo
	Select(conds ...field.Expr) IRacaoMinimaDo
	Where(conds ...gen.Condition) IRacaoMinimaDo
	Order(conds ...field.Expr) IRacaoMinimaDo
	Distinct(cols ...field.Expr) IRacaoMinimaDo
	Omit(cols ...field.Expr) IRacaoMinimaDo
	Join(table schema.Tabler, on ...field.Expr) IRacaoMinimaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRacaoMinimaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRacaoMinimaDo
	Group(cols ...field.Expr) IRacaoMinimaDo
	Having(conds ...gen.Condition) IRacaoMinimaDo
	Limit(limit int) IRacaoMinimaDo
	Offset(offset int) IRacaoMinimaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRacaoMinimaDo
	Unscoped() IRacaoMinimaDo
	Create(values ...*model.RacaoMinima) error
	CreateInBatches(values []*model.RacaoMinima, batchSize int) error
	Save(values ...*model.RacaoMinima) error
	First() (*model.RacaoMinima, error)
	Take() (*model.RacaoMinima, error)
	Last() (*model.RacaoMinima, error)
	Find() ([]*model.RacaoMinima, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RacaoMinima, err error)
	FindInBatches(result *[]*model.RacaoMinima, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RacaoMinima) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRacaoMinimaDo
	Assign(attrs ...field.AssignExpr) IRacaoMinimaDo
	Joins(fields ...field.RelationField) IRacaoMinimaDo
	Preload(fields ...field.RelationField) IRacaoMinimaDo
	FirstOrInit() (*model.RacaoMinima, error)
	FirstOrCreate() (*model.RacaoMinima, error)
	FindByPage(offset int, limit int) (result []*model.RacaoMinima, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRacaoMinimaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r racaoMinimaDo) Debug() IRacaoMinimaDo {
	return r.withDO(r.DO.Debug())
}

func (r racaoMinimaDo) WithContext(ctx context.Context) IRacaoMinimaDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r racaoMinimaDo) ReadDB() IRacaoMinimaDo {
	return r.Clauses(dbresolver.Read)
}

func (r racaoMinimaDo) WriteDB() IRacaoMinimaDo {
	return r.Clauses(dbresolver.Write)
}

func (r racaoMinimaDo) Session(config *gorm.Session) IRacaoMinimaDo {
	return r.withDO(r.DO.Session(config))
}

func (r racaoMinimaDo) Clauses(conds ...clause.Expression) IRacaoMinimaDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r racaoMinimaDo) Returning(value interface{}, columns ...string) IRacaoMinimaDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r racaoMinimaDo) Not(conds ...gen.Condition) IRacaoMinimaDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r racaoMinimaDo) Or(conds ...gen.Condition) IRacaoMinimaDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r racaoMinimaDo) Select(conds ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r racaoMinimaDo) Where(conds ...gen.Condition) IRacaoMinimaDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r racaoMinimaDo) Order(conds ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r racaoMinimaDo) Distinct(cols ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r racaoMinimaDo) Omit(cols ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r racaoMinimaDo) Join(table schema.Tabler, on ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r racaoMinimaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r racaoMinimaDo) RightJoin(table schema.Tabler, on ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r racaoMinimaDo) Group(cols ...field.Expr) IRacaoMinimaDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r racaoMinimaDo) Having(conds ...gen.Condition) IRacaoMinimaDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r racaoMinimaDo) Limit(limit int) IRacaoMinimaDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r racaoMinimaDo) Offset(offset int) IRacaoMinimaDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r racaoMinimaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRacaoMinimaDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r racaoMinimaDo) Unscoped() IRacaoMinimaDo {
	return r.withDO(r.DO.Unscoped())
}

func (r racaoMinimaDo) Create(values ...*model.RacaoMinima) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r racaoMinimaDo) CreateInBatches(values []*model.RacaoMinima, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r racaoMinimaDo) Save(values ...*model.RacaoMinima) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r racaoMinimaDo) First() (*model.RacaoMinima, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RacaoMinima), nil
	}
}

func (r racaoMinimaDo) Take() (*model.RacaoMinima, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RacaoMinima), nil
	}
}

func (r racaoMinimaDo) Last() (*model.RacaoMinima, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RacaoMinima), nil
	}
}

func (r racaoMinimaDo) Find() ([]*model.RacaoMinima, error) {
	result, err := r.DO.Find()
	return result.([]*model.RacaoMinima), err
}

func (r racaoMinimaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RacaoMinima, err error) {
	buf := make([]*model.RacaoMinima, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r racaoMinimaDo) FindInBatches(result *[]*model.RacaoMinima, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r racaoMinimaDo) Attrs(attrs ...field.AssignExpr) IRacaoMinimaDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r racaoMinimaDo) Assign(attrs ...field.AssignExpr) IRacaoMinimaDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r racaoMinimaDo) Joins(fields ...field.RelationField) IRacaoMinimaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r racaoMinimaDo) Preload(fields ...field.RelationField) IRacaoMinimaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r racaoMinimaDo) FirstOrInit() (*model.RacaoMinima, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RacaoMinima), nil
	}
}

func (r racaoMinimaDo) FirstOrCreate() (*model.RacaoMinima, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RacaoMinima), nil
	}
}

func (r racaoMinimaDo) FindByPage(offset int, limit int) (result []*model.RacaoMinima, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r racaoMinimaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r racaoMinimaDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r racaoMinimaDo) Delete(models ...*model.RacaoMinima) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *racaoMinimaDo) withDO(do gen.Dao) *racaoMinimaDo {
	r.DO = *do.(*gen.DO)
	return r
}
