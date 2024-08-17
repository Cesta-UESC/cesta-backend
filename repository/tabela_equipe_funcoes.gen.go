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

func newEquipeFuncoes(db *gorm.DB, opts ...gen.DOOption) equipeFuncoes {
	_equipeFuncoes := equipeFuncoes{}

	_equipeFuncoes.equipeFuncoesDo.UseDB(db, opts...)
	_equipeFuncoes.equipeFuncoesDo.UseModel(&model.EquipeFuncoes{})

	tableName := _equipeFuncoes.equipeFuncoesDo.TableName()
	_equipeFuncoes.ALL = field.NewAsterisk(tableName)
	_equipeFuncoes.ID = field.NewInt32(tableName, "id")
	_equipeFuncoes.Funcao = field.NewString(tableName, "funcao")

	_equipeFuncoes.fillFieldMap()

	return _equipeFuncoes
}

type equipeFuncoes struct {
	equipeFuncoesDo

	ALL    field.Asterisk
	ID     field.Int32
	Funcao field.String

	fieldMap map[string]field.Expr
}

func (e equipeFuncoes) Table(newTableName string) *equipeFuncoes {
	e.equipeFuncoesDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipeFuncoes) As(alias string) *equipeFuncoes {
	e.equipeFuncoesDo.DO = *(e.equipeFuncoesDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipeFuncoes) updateTableName(table string) *equipeFuncoes {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Funcao = field.NewString(table, "funcao")

	e.fillFieldMap()

	return e
}

func (e *equipeFuncoes) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipeFuncoes) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 2)
	e.fieldMap["id"] = e.ID
	e.fieldMap["funcao"] = e.Funcao
}

func (e equipeFuncoes) clone(db *gorm.DB) equipeFuncoes {
	e.equipeFuncoesDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipeFuncoes) replaceDB(db *gorm.DB) equipeFuncoes {
	e.equipeFuncoesDo.ReplaceDB(db)
	return e
}

type equipeFuncoesDo struct{ gen.DO }

type IEquipeFuncoesDo interface {
	gen.SubQuery
	Debug() IEquipeFuncoesDo
	WithContext(ctx context.Context) IEquipeFuncoesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEquipeFuncoesDo
	WriteDB() IEquipeFuncoesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEquipeFuncoesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEquipeFuncoesDo
	Not(conds ...gen.Condition) IEquipeFuncoesDo
	Or(conds ...gen.Condition) IEquipeFuncoesDo
	Select(conds ...field.Expr) IEquipeFuncoesDo
	Where(conds ...gen.Condition) IEquipeFuncoesDo
	Order(conds ...field.Expr) IEquipeFuncoesDo
	Distinct(cols ...field.Expr) IEquipeFuncoesDo
	Omit(cols ...field.Expr) IEquipeFuncoesDo
	Join(table schema.Tabler, on ...field.Expr) IEquipeFuncoesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEquipeFuncoesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEquipeFuncoesDo
	Group(cols ...field.Expr) IEquipeFuncoesDo
	Having(conds ...gen.Condition) IEquipeFuncoesDo
	Limit(limit int) IEquipeFuncoesDo
	Offset(offset int) IEquipeFuncoesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEquipeFuncoesDo
	Unscoped() IEquipeFuncoesDo
	Create(values ...*model.EquipeFuncoes) error
	CreateInBatches(values []*model.EquipeFuncoes, batchSize int) error
	Save(values ...*model.EquipeFuncoes) error
	First() (*model.EquipeFuncoes, error)
	Take() (*model.EquipeFuncoes, error)
	Last() (*model.EquipeFuncoes, error)
	Find() ([]*model.EquipeFuncoes, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EquipeFuncoes, err error)
	FindInBatches(result *[]*model.EquipeFuncoes, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EquipeFuncoes) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEquipeFuncoesDo
	Assign(attrs ...field.AssignExpr) IEquipeFuncoesDo
	Joins(fields ...field.RelationField) IEquipeFuncoesDo
	Preload(fields ...field.RelationField) IEquipeFuncoesDo
	FirstOrInit() (*model.EquipeFuncoes, error)
	FirstOrCreate() (*model.EquipeFuncoes, error)
	FindByPage(offset int, limit int) (result []*model.EquipeFuncoes, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEquipeFuncoesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e equipeFuncoesDo) Debug() IEquipeFuncoesDo {
	return e.withDO(e.DO.Debug())
}

func (e equipeFuncoesDo) WithContext(ctx context.Context) IEquipeFuncoesDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipeFuncoesDo) ReadDB() IEquipeFuncoesDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipeFuncoesDo) WriteDB() IEquipeFuncoesDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipeFuncoesDo) Session(config *gorm.Session) IEquipeFuncoesDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipeFuncoesDo) Clauses(conds ...clause.Expression) IEquipeFuncoesDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipeFuncoesDo) Returning(value interface{}, columns ...string) IEquipeFuncoesDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipeFuncoesDo) Not(conds ...gen.Condition) IEquipeFuncoesDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipeFuncoesDo) Or(conds ...gen.Condition) IEquipeFuncoesDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipeFuncoesDo) Select(conds ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipeFuncoesDo) Where(conds ...gen.Condition) IEquipeFuncoesDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipeFuncoesDo) Order(conds ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipeFuncoesDo) Distinct(cols ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipeFuncoesDo) Omit(cols ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipeFuncoesDo) Join(table schema.Tabler, on ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipeFuncoesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipeFuncoesDo) RightJoin(table schema.Tabler, on ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipeFuncoesDo) Group(cols ...field.Expr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipeFuncoesDo) Having(conds ...gen.Condition) IEquipeFuncoesDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipeFuncoesDo) Limit(limit int) IEquipeFuncoesDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipeFuncoesDo) Offset(offset int) IEquipeFuncoesDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipeFuncoesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEquipeFuncoesDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipeFuncoesDo) Unscoped() IEquipeFuncoesDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipeFuncoesDo) Create(values ...*model.EquipeFuncoes) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipeFuncoesDo) CreateInBatches(values []*model.EquipeFuncoes, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipeFuncoesDo) Save(values ...*model.EquipeFuncoes) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipeFuncoesDo) First() (*model.EquipeFuncoes, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EquipeFuncoes), nil
	}
}

func (e equipeFuncoesDo) Take() (*model.EquipeFuncoes, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EquipeFuncoes), nil
	}
}

func (e equipeFuncoesDo) Last() (*model.EquipeFuncoes, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EquipeFuncoes), nil
	}
}

func (e equipeFuncoesDo) Find() ([]*model.EquipeFuncoes, error) {
	result, err := e.DO.Find()
	return result.([]*model.EquipeFuncoes), err
}

func (e equipeFuncoesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EquipeFuncoes, err error) {
	buf := make([]*model.EquipeFuncoes, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipeFuncoesDo) FindInBatches(result *[]*model.EquipeFuncoes, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipeFuncoesDo) Attrs(attrs ...field.AssignExpr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipeFuncoesDo) Assign(attrs ...field.AssignExpr) IEquipeFuncoesDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipeFuncoesDo) Joins(fields ...field.RelationField) IEquipeFuncoesDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipeFuncoesDo) Preload(fields ...field.RelationField) IEquipeFuncoesDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipeFuncoesDo) FirstOrInit() (*model.EquipeFuncoes, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EquipeFuncoes), nil
	}
}

func (e equipeFuncoesDo) FirstOrCreate() (*model.EquipeFuncoes, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EquipeFuncoes), nil
	}
}

func (e equipeFuncoesDo) FindByPage(offset int, limit int) (result []*model.EquipeFuncoes, count int64, err error) {
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

func (e equipeFuncoesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipeFuncoesDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipeFuncoesDo) Delete(models ...*model.EquipeFuncoes) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipeFuncoesDo) withDO(do gen.Dao) *equipeFuncoesDo {
	e.DO = *do.(*gen.DO)
	return e
}
