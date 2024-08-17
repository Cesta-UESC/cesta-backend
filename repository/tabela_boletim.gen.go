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

func newBoletim(db *gorm.DB, opts ...gen.DOOption) boletim {
	_boletim := boletim{}

	_boletim.boletimDo.UseDB(db, opts...)
	_boletim.boletimDo.UseModel(&model.Boletim{})

	tableName := _boletim.boletimDo.TableName()
	_boletim.ALL = field.NewAsterisk(tableName)
	_boletim.BoletimID = field.NewInt32(tableName, "boletim_id")
	_boletim.BoletimNome = field.NewString(tableName, "boletim_nome")

	_boletim.fillFieldMap()

	return _boletim
}

type boletim struct {
	boletimDo

	ALL         field.Asterisk
	BoletimID   field.Int32
	BoletimNome field.String

	fieldMap map[string]field.Expr
}

func (b boletim) Table(newTableName string) *boletim {
	b.boletimDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b boletim) As(alias string) *boletim {
	b.boletimDo.DO = *(b.boletimDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *boletim) updateTableName(table string) *boletim {
	b.ALL = field.NewAsterisk(table)
	b.BoletimID = field.NewInt32(table, "boletim_id")
	b.BoletimNome = field.NewString(table, "boletim_nome")

	b.fillFieldMap()

	return b
}

func (b *boletim) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *boletim) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 2)
	b.fieldMap["boletim_id"] = b.BoletimID
	b.fieldMap["boletim_nome"] = b.BoletimNome
}

func (b boletim) clone(db *gorm.DB) boletim {
	b.boletimDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b boletim) replaceDB(db *gorm.DB) boletim {
	b.boletimDo.ReplaceDB(db)
	return b
}

type boletimDo struct{ gen.DO }

type IBoletimDo interface {
	gen.SubQuery
	Debug() IBoletimDo
	WithContext(ctx context.Context) IBoletimDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBoletimDo
	WriteDB() IBoletimDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBoletimDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBoletimDo
	Not(conds ...gen.Condition) IBoletimDo
	Or(conds ...gen.Condition) IBoletimDo
	Select(conds ...field.Expr) IBoletimDo
	Where(conds ...gen.Condition) IBoletimDo
	Order(conds ...field.Expr) IBoletimDo
	Distinct(cols ...field.Expr) IBoletimDo
	Omit(cols ...field.Expr) IBoletimDo
	Join(table schema.Tabler, on ...field.Expr) IBoletimDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBoletimDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBoletimDo
	Group(cols ...field.Expr) IBoletimDo
	Having(conds ...gen.Condition) IBoletimDo
	Limit(limit int) IBoletimDo
	Offset(offset int) IBoletimDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBoletimDo
	Unscoped() IBoletimDo
	Create(values ...*model.Boletim) error
	CreateInBatches(values []*model.Boletim, batchSize int) error
	Save(values ...*model.Boletim) error
	First() (*model.Boletim, error)
	Take() (*model.Boletim, error)
	Last() (*model.Boletim, error)
	Find() ([]*model.Boletim, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Boletim, err error)
	FindInBatches(result *[]*model.Boletim, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Boletim) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBoletimDo
	Assign(attrs ...field.AssignExpr) IBoletimDo
	Joins(fields ...field.RelationField) IBoletimDo
	Preload(fields ...field.RelationField) IBoletimDo
	FirstOrInit() (*model.Boletim, error)
	FirstOrCreate() (*model.Boletim, error)
	FindByPage(offset int, limit int) (result []*model.Boletim, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBoletimDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b boletimDo) Debug() IBoletimDo {
	return b.withDO(b.DO.Debug())
}

func (b boletimDo) WithContext(ctx context.Context) IBoletimDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b boletimDo) ReadDB() IBoletimDo {
	return b.Clauses(dbresolver.Read)
}

func (b boletimDo) WriteDB() IBoletimDo {
	return b.Clauses(dbresolver.Write)
}

func (b boletimDo) Session(config *gorm.Session) IBoletimDo {
	return b.withDO(b.DO.Session(config))
}

func (b boletimDo) Clauses(conds ...clause.Expression) IBoletimDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b boletimDo) Returning(value interface{}, columns ...string) IBoletimDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b boletimDo) Not(conds ...gen.Condition) IBoletimDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b boletimDo) Or(conds ...gen.Condition) IBoletimDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b boletimDo) Select(conds ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b boletimDo) Where(conds ...gen.Condition) IBoletimDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b boletimDo) Order(conds ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b boletimDo) Distinct(cols ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b boletimDo) Omit(cols ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b boletimDo) Join(table schema.Tabler, on ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b boletimDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b boletimDo) RightJoin(table schema.Tabler, on ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b boletimDo) Group(cols ...field.Expr) IBoletimDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b boletimDo) Having(conds ...gen.Condition) IBoletimDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b boletimDo) Limit(limit int) IBoletimDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b boletimDo) Offset(offset int) IBoletimDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b boletimDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBoletimDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b boletimDo) Unscoped() IBoletimDo {
	return b.withDO(b.DO.Unscoped())
}

func (b boletimDo) Create(values ...*model.Boletim) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b boletimDo) CreateInBatches(values []*model.Boletim, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b boletimDo) Save(values ...*model.Boletim) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b boletimDo) First() (*model.Boletim, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Boletim), nil
	}
}

func (b boletimDo) Take() (*model.Boletim, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Boletim), nil
	}
}

func (b boletimDo) Last() (*model.Boletim, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Boletim), nil
	}
}

func (b boletimDo) Find() ([]*model.Boletim, error) {
	result, err := b.DO.Find()
	return result.([]*model.Boletim), err
}

func (b boletimDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Boletim, err error) {
	buf := make([]*model.Boletim, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b boletimDo) FindInBatches(result *[]*model.Boletim, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b boletimDo) Attrs(attrs ...field.AssignExpr) IBoletimDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b boletimDo) Assign(attrs ...field.AssignExpr) IBoletimDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b boletimDo) Joins(fields ...field.RelationField) IBoletimDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b boletimDo) Preload(fields ...field.RelationField) IBoletimDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b boletimDo) FirstOrInit() (*model.Boletim, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Boletim), nil
	}
}

func (b boletimDo) FirstOrCreate() (*model.Boletim, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Boletim), nil
	}
}

func (b boletimDo) FindByPage(offset int, limit int) (result []*model.Boletim, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b boletimDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b boletimDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b boletimDo) Delete(models ...*model.Boletim) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *boletimDo) withDO(do gen.Dao) *boletimDo {
	b.DO = *do.(*gen.DO)
	return b
}
