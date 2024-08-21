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

func newAuxiliarPrecos(db *gorm.DB, opts ...gen.DOOption) auxiliarPrecos {
	_auxiliarPrecos := auxiliarPrecos{}

	_auxiliarPrecos.auxiliarPrecosDo.UseDB(db, opts...)
	_auxiliarPrecos.auxiliarPrecosDo.UseModel(&models.AuxiliarPrecos{})

	tableName := _auxiliarPrecos.auxiliarPrecosDo.TableName()
	_auxiliarPrecos.ALL = field.NewAsterisk(tableName)
	_auxiliarPrecos.PrecosID = field.NewInt32(tableName, "precos_id")
	_auxiliarPrecos.PrecoProduto = field.NewFloat64(tableName, "preco_produto")

	_auxiliarPrecos.fillFieldMap()

	return _auxiliarPrecos
}

type auxiliarPrecos struct {
	auxiliarPrecosDo

	ALL          field.Asterisk
	PrecosID     field.Int32
	PrecoProduto field.Float64

	fieldMap map[string]field.Expr
}

func (a auxiliarPrecos) Table(newTableName string) *auxiliarPrecos {
	a.auxiliarPrecosDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a auxiliarPrecos) As(alias string) *auxiliarPrecos {
	a.auxiliarPrecosDo.DO = *(a.auxiliarPrecosDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *auxiliarPrecos) updateTableName(table string) *auxiliarPrecos {
	a.ALL = field.NewAsterisk(table)
	a.PrecosID = field.NewInt32(table, "precos_id")
	a.PrecoProduto = field.NewFloat64(table, "preco_produto")

	a.fillFieldMap()

	return a
}

func (a *auxiliarPrecos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *auxiliarPrecos) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["precos_id"] = a.PrecosID
	a.fieldMap["preco_produto"] = a.PrecoProduto
}

func (a auxiliarPrecos) clone(db *gorm.DB) auxiliarPrecos {
	a.auxiliarPrecosDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a auxiliarPrecos) replaceDB(db *gorm.DB) auxiliarPrecos {
	a.auxiliarPrecosDo.ReplaceDB(db)
	return a
}

type auxiliarPrecosDo struct{ gen.DO }

type IAuxiliarPrecosDo interface {
	gen.SubQuery
	Debug() IAuxiliarPrecosDo
	WithContext(ctx context.Context) IAuxiliarPrecosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuxiliarPrecosDo
	WriteDB() IAuxiliarPrecosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuxiliarPrecosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuxiliarPrecosDo
	Not(conds ...gen.Condition) IAuxiliarPrecosDo
	Or(conds ...gen.Condition) IAuxiliarPrecosDo
	Select(conds ...field.Expr) IAuxiliarPrecosDo
	Where(conds ...gen.Condition) IAuxiliarPrecosDo
	Order(conds ...field.Expr) IAuxiliarPrecosDo
	Distinct(cols ...field.Expr) IAuxiliarPrecosDo
	Omit(cols ...field.Expr) IAuxiliarPrecosDo
	Join(table schema.Tabler, on ...field.Expr) IAuxiliarPrecosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuxiliarPrecosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuxiliarPrecosDo
	Group(cols ...field.Expr) IAuxiliarPrecosDo
	Having(conds ...gen.Condition) IAuxiliarPrecosDo
	Limit(limit int) IAuxiliarPrecosDo
	Offset(offset int) IAuxiliarPrecosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuxiliarPrecosDo
	Unscoped() IAuxiliarPrecosDo
	Create(values ...*models.AuxiliarPrecos) error
	CreateInBatches(values []*models.AuxiliarPrecos, batchSize int) error
	Save(values ...*models.AuxiliarPrecos) error
	First() (*models.AuxiliarPrecos, error)
	Take() (*models.AuxiliarPrecos, error)
	Last() (*models.AuxiliarPrecos, error)
	Find() ([]*models.AuxiliarPrecos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuxiliarPrecos, err error)
	FindInBatches(result *[]*models.AuxiliarPrecos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.AuxiliarPrecos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuxiliarPrecosDo
	Assign(attrs ...field.AssignExpr) IAuxiliarPrecosDo
	Joins(fields ...field.RelationField) IAuxiliarPrecosDo
	Preload(fields ...field.RelationField) IAuxiliarPrecosDo
	FirstOrInit() (*models.AuxiliarPrecos, error)
	FirstOrCreate() (*models.AuxiliarPrecos, error)
	FindByPage(offset int, limit int) (result []*models.AuxiliarPrecos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuxiliarPrecosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a auxiliarPrecosDo) Debug() IAuxiliarPrecosDo {
	return a.withDO(a.DO.Debug())
}

func (a auxiliarPrecosDo) WithContext(ctx context.Context) IAuxiliarPrecosDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a auxiliarPrecosDo) ReadDB() IAuxiliarPrecosDo {
	return a.Clauses(dbresolver.Read)
}

func (a auxiliarPrecosDo) WriteDB() IAuxiliarPrecosDo {
	return a.Clauses(dbresolver.Write)
}

func (a auxiliarPrecosDo) Session(config *gorm.Session) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Session(config))
}

func (a auxiliarPrecosDo) Clauses(conds ...clause.Expression) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a auxiliarPrecosDo) Returning(value interface{}, columns ...string) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a auxiliarPrecosDo) Not(conds ...gen.Condition) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a auxiliarPrecosDo) Or(conds ...gen.Condition) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a auxiliarPrecosDo) Select(conds ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a auxiliarPrecosDo) Where(conds ...gen.Condition) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a auxiliarPrecosDo) Order(conds ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a auxiliarPrecosDo) Distinct(cols ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a auxiliarPrecosDo) Omit(cols ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a auxiliarPrecosDo) Join(table schema.Tabler, on ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a auxiliarPrecosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a auxiliarPrecosDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a auxiliarPrecosDo) Group(cols ...field.Expr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a auxiliarPrecosDo) Having(conds ...gen.Condition) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a auxiliarPrecosDo) Limit(limit int) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a auxiliarPrecosDo) Offset(offset int) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a auxiliarPrecosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a auxiliarPrecosDo) Unscoped() IAuxiliarPrecosDo {
	return a.withDO(a.DO.Unscoped())
}

func (a auxiliarPrecosDo) Create(values ...*models.AuxiliarPrecos) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a auxiliarPrecosDo) CreateInBatches(values []*models.AuxiliarPrecos, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a auxiliarPrecosDo) Save(values ...*models.AuxiliarPrecos) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a auxiliarPrecosDo) First() (*models.AuxiliarPrecos, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuxiliarPrecos), nil
	}
}

func (a auxiliarPrecosDo) Take() (*models.AuxiliarPrecos, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuxiliarPrecos), nil
	}
}

func (a auxiliarPrecosDo) Last() (*models.AuxiliarPrecos, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuxiliarPrecos), nil
	}
}

func (a auxiliarPrecosDo) Find() ([]*models.AuxiliarPrecos, error) {
	result, err := a.DO.Find()
	return result.([]*models.AuxiliarPrecos), err
}

func (a auxiliarPrecosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuxiliarPrecos, err error) {
	buf := make([]*models.AuxiliarPrecos, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a auxiliarPrecosDo) FindInBatches(result *[]*models.AuxiliarPrecos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a auxiliarPrecosDo) Attrs(attrs ...field.AssignExpr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a auxiliarPrecosDo) Assign(attrs ...field.AssignExpr) IAuxiliarPrecosDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a auxiliarPrecosDo) Joins(fields ...field.RelationField) IAuxiliarPrecosDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a auxiliarPrecosDo) Preload(fields ...field.RelationField) IAuxiliarPrecosDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a auxiliarPrecosDo) FirstOrInit() (*models.AuxiliarPrecos, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuxiliarPrecos), nil
	}
}

func (a auxiliarPrecosDo) FirstOrCreate() (*models.AuxiliarPrecos, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuxiliarPrecos), nil
	}
}

func (a auxiliarPrecosDo) FindByPage(offset int, limit int) (result []*models.AuxiliarPrecos, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a auxiliarPrecosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a auxiliarPrecosDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a auxiliarPrecosDo) Delete(models ...*models.AuxiliarPrecos) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *auxiliarPrecosDo) withDO(do gen.Dao) *auxiliarPrecosDo {
	a.DO = *do.(*gen.DO)
	return a
}
