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

func newTiposProdutos(db *gorm.DB, opts ...gen.DOOption) tiposProdutos {
	_tiposProdutos := tiposProdutos{}

	_tiposProdutos.tiposProdutosDo.UseDB(db, opts...)
	_tiposProdutos.tiposProdutosDo.UseModel(&models.TiposProdutos{})

	tableName := _tiposProdutos.tiposProdutosDo.TableName()
	_tiposProdutos.ALL = field.NewAsterisk(tableName)
	_tiposProdutos.TipoID = field.NewInt32(tableName, "tipo_id")
	_tiposProdutos.TipoNome = field.NewString(tableName, "tipo_nome")

	_tiposProdutos.fillFieldMap()

	return _tiposProdutos
}

type tiposProdutos struct {
	tiposProdutosDo

	ALL      field.Asterisk
	TipoID   field.Int32
	TipoNome field.String

	fieldMap map[string]field.Expr
}

func (t tiposProdutos) Table(newTableName string) *tiposProdutos {
	t.tiposProdutosDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tiposProdutos) As(alias string) *tiposProdutos {
	t.tiposProdutosDo.DO = *(t.tiposProdutosDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tiposProdutos) updateTableName(table string) *tiposProdutos {
	t.ALL = field.NewAsterisk(table)
	t.TipoID = field.NewInt32(table, "tipo_id")
	t.TipoNome = field.NewString(table, "tipo_nome")

	t.fillFieldMap()

	return t
}

func (t *tiposProdutos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tiposProdutos) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["tipo_id"] = t.TipoID
	t.fieldMap["tipo_nome"] = t.TipoNome
}

func (t tiposProdutos) clone(db *gorm.DB) tiposProdutos {
	t.tiposProdutosDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tiposProdutos) replaceDB(db *gorm.DB) tiposProdutos {
	t.tiposProdutosDo.ReplaceDB(db)
	return t
}

type tiposProdutosDo struct{ gen.DO }

type ITiposProdutosDo interface {
	gen.SubQuery
	Debug() ITiposProdutosDo
	WithContext(ctx context.Context) ITiposProdutosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITiposProdutosDo
	WriteDB() ITiposProdutosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITiposProdutosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITiposProdutosDo
	Not(conds ...gen.Condition) ITiposProdutosDo
	Or(conds ...gen.Condition) ITiposProdutosDo
	Select(conds ...field.Expr) ITiposProdutosDo
	Where(conds ...gen.Condition) ITiposProdutosDo
	Order(conds ...field.Expr) ITiposProdutosDo
	Distinct(cols ...field.Expr) ITiposProdutosDo
	Omit(cols ...field.Expr) ITiposProdutosDo
	Join(table schema.Tabler, on ...field.Expr) ITiposProdutosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITiposProdutosDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITiposProdutosDo
	Group(cols ...field.Expr) ITiposProdutosDo
	Having(conds ...gen.Condition) ITiposProdutosDo
	Limit(limit int) ITiposProdutosDo
	Offset(offset int) ITiposProdutosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITiposProdutosDo
	Unscoped() ITiposProdutosDo
	Create(values ...*models.TiposProdutos) error
	CreateInBatches(values []*models.TiposProdutos, batchSize int) error
	Save(values ...*models.TiposProdutos) error
	First() (*models.TiposProdutos, error)
	Take() (*models.TiposProdutos, error)
	Last() (*models.TiposProdutos, error)
	Find() ([]*models.TiposProdutos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.TiposProdutos, err error)
	FindInBatches(result *[]*models.TiposProdutos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.TiposProdutos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITiposProdutosDo
	Assign(attrs ...field.AssignExpr) ITiposProdutosDo
	Joins(fields ...field.RelationField) ITiposProdutosDo
	Preload(fields ...field.RelationField) ITiposProdutosDo
	FirstOrInit() (*models.TiposProdutos, error)
	FirstOrCreate() (*models.TiposProdutos, error)
	FindByPage(offset int, limit int) (result []*models.TiposProdutos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITiposProdutosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tiposProdutosDo) Debug() ITiposProdutosDo {
	return t.withDO(t.DO.Debug())
}

func (t tiposProdutosDo) WithContext(ctx context.Context) ITiposProdutosDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tiposProdutosDo) ReadDB() ITiposProdutosDo {
	return t.Clauses(dbresolver.Read)
}

func (t tiposProdutosDo) WriteDB() ITiposProdutosDo {
	return t.Clauses(dbresolver.Write)
}

func (t tiposProdutosDo) Session(config *gorm.Session) ITiposProdutosDo {
	return t.withDO(t.DO.Session(config))
}

func (t tiposProdutosDo) Clauses(conds ...clause.Expression) ITiposProdutosDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tiposProdutosDo) Returning(value interface{}, columns ...string) ITiposProdutosDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tiposProdutosDo) Not(conds ...gen.Condition) ITiposProdutosDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tiposProdutosDo) Or(conds ...gen.Condition) ITiposProdutosDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tiposProdutosDo) Select(conds ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tiposProdutosDo) Where(conds ...gen.Condition) ITiposProdutosDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tiposProdutosDo) Order(conds ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tiposProdutosDo) Distinct(cols ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tiposProdutosDo) Omit(cols ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tiposProdutosDo) Join(table schema.Tabler, on ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tiposProdutosDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tiposProdutosDo) RightJoin(table schema.Tabler, on ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tiposProdutosDo) Group(cols ...field.Expr) ITiposProdutosDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tiposProdutosDo) Having(conds ...gen.Condition) ITiposProdutosDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tiposProdutosDo) Limit(limit int) ITiposProdutosDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tiposProdutosDo) Offset(offset int) ITiposProdutosDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tiposProdutosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITiposProdutosDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tiposProdutosDo) Unscoped() ITiposProdutosDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tiposProdutosDo) Create(values ...*models.TiposProdutos) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tiposProdutosDo) CreateInBatches(values []*models.TiposProdutos, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tiposProdutosDo) Save(values ...*models.TiposProdutos) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tiposProdutosDo) First() (*models.TiposProdutos, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.TiposProdutos), nil
	}
}

func (t tiposProdutosDo) Take() (*models.TiposProdutos, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.TiposProdutos), nil
	}
}

func (t tiposProdutosDo) Last() (*models.TiposProdutos, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.TiposProdutos), nil
	}
}

func (t tiposProdutosDo) Find() ([]*models.TiposProdutos, error) {
	result, err := t.DO.Find()
	return result.([]*models.TiposProdutos), err
}

func (t tiposProdutosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.TiposProdutos, err error) {
	buf := make([]*models.TiposProdutos, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tiposProdutosDo) FindInBatches(result *[]*models.TiposProdutos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tiposProdutosDo) Attrs(attrs ...field.AssignExpr) ITiposProdutosDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tiposProdutosDo) Assign(attrs ...field.AssignExpr) ITiposProdutosDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tiposProdutosDo) Joins(fields ...field.RelationField) ITiposProdutosDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tiposProdutosDo) Preload(fields ...field.RelationField) ITiposProdutosDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tiposProdutosDo) FirstOrInit() (*models.TiposProdutos, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.TiposProdutos), nil
	}
}

func (t tiposProdutosDo) FirstOrCreate() (*models.TiposProdutos, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.TiposProdutos), nil
	}
}

func (t tiposProdutosDo) FindByPage(offset int, limit int) (result []*models.TiposProdutos, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tiposProdutosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tiposProdutosDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tiposProdutosDo) Delete(models ...*models.TiposProdutos) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tiposProdutosDo) withDO(do gen.Dao) *tiposProdutosDo {
	t.DO = *do.(*gen.DO)
	return t
}
