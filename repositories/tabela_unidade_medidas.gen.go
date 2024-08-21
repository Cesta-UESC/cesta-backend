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

func newUnidadeMedidas(db *gorm.DB, opts ...gen.DOOption) unidadeMedidas {
	_unidadeMedidas := unidadeMedidas{}

	_unidadeMedidas.unidadeMedidasDo.UseDB(db, opts...)
	_unidadeMedidas.unidadeMedidasDo.UseModel(&models.UnidadeMedidas{})

	tableName := _unidadeMedidas.unidadeMedidasDo.TableName()
	_unidadeMedidas.ALL = field.NewAsterisk(tableName)
	_unidadeMedidas.MedidaID = field.NewInt32(tableName, "medida_id")
	_unidadeMedidas.MedidaDescricao = field.NewString(tableName, "medida_descricao")
	_unidadeMedidas.MedidaSimbolo = field.NewString(tableName, "medida_simbolo")

	_unidadeMedidas.fillFieldMap()

	return _unidadeMedidas
}

type unidadeMedidas struct {
	unidadeMedidasDo

	ALL             field.Asterisk
	MedidaID        field.Int32
	MedidaDescricao field.String
	MedidaSimbolo   field.String

	fieldMap map[string]field.Expr
}

func (u unidadeMedidas) Table(newTableName string) *unidadeMedidas {
	u.unidadeMedidasDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u unidadeMedidas) As(alias string) *unidadeMedidas {
	u.unidadeMedidasDo.DO = *(u.unidadeMedidasDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *unidadeMedidas) updateTableName(table string) *unidadeMedidas {
	u.ALL = field.NewAsterisk(table)
	u.MedidaID = field.NewInt32(table, "medida_id")
	u.MedidaDescricao = field.NewString(table, "medida_descricao")
	u.MedidaSimbolo = field.NewString(table, "medida_simbolo")

	u.fillFieldMap()

	return u
}

func (u *unidadeMedidas) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *unidadeMedidas) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["medida_id"] = u.MedidaID
	u.fieldMap["medida_descricao"] = u.MedidaDescricao
	u.fieldMap["medida_simbolo"] = u.MedidaSimbolo
}

func (u unidadeMedidas) clone(db *gorm.DB) unidadeMedidas {
	u.unidadeMedidasDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u unidadeMedidas) replaceDB(db *gorm.DB) unidadeMedidas {
	u.unidadeMedidasDo.ReplaceDB(db)
	return u
}

type unidadeMedidasDo struct{ gen.DO }

type IUnidadeMedidasDo interface {
	gen.SubQuery
	Debug() IUnidadeMedidasDo
	WithContext(ctx context.Context) IUnidadeMedidasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUnidadeMedidasDo
	WriteDB() IUnidadeMedidasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUnidadeMedidasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUnidadeMedidasDo
	Not(conds ...gen.Condition) IUnidadeMedidasDo
	Or(conds ...gen.Condition) IUnidadeMedidasDo
	Select(conds ...field.Expr) IUnidadeMedidasDo
	Where(conds ...gen.Condition) IUnidadeMedidasDo
	Order(conds ...field.Expr) IUnidadeMedidasDo
	Distinct(cols ...field.Expr) IUnidadeMedidasDo
	Omit(cols ...field.Expr) IUnidadeMedidasDo
	Join(table schema.Tabler, on ...field.Expr) IUnidadeMedidasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUnidadeMedidasDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUnidadeMedidasDo
	Group(cols ...field.Expr) IUnidadeMedidasDo
	Having(conds ...gen.Condition) IUnidadeMedidasDo
	Limit(limit int) IUnidadeMedidasDo
	Offset(offset int) IUnidadeMedidasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUnidadeMedidasDo
	Unscoped() IUnidadeMedidasDo
	Create(values ...*models.UnidadeMedidas) error
	CreateInBatches(values []*models.UnidadeMedidas, batchSize int) error
	Save(values ...*models.UnidadeMedidas) error
	First() (*models.UnidadeMedidas, error)
	Take() (*models.UnidadeMedidas, error)
	Last() (*models.UnidadeMedidas, error)
	Find() ([]*models.UnidadeMedidas, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UnidadeMedidas, err error)
	FindInBatches(result *[]*models.UnidadeMedidas, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.UnidadeMedidas) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUnidadeMedidasDo
	Assign(attrs ...field.AssignExpr) IUnidadeMedidasDo
	Joins(fields ...field.RelationField) IUnidadeMedidasDo
	Preload(fields ...field.RelationField) IUnidadeMedidasDo
	FirstOrInit() (*models.UnidadeMedidas, error)
	FirstOrCreate() (*models.UnidadeMedidas, error)
	FindByPage(offset int, limit int) (result []*models.UnidadeMedidas, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUnidadeMedidasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u unidadeMedidasDo) Debug() IUnidadeMedidasDo {
	return u.withDO(u.DO.Debug())
}

func (u unidadeMedidasDo) WithContext(ctx context.Context) IUnidadeMedidasDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u unidadeMedidasDo) ReadDB() IUnidadeMedidasDo {
	return u.Clauses(dbresolver.Read)
}

func (u unidadeMedidasDo) WriteDB() IUnidadeMedidasDo {
	return u.Clauses(dbresolver.Write)
}

func (u unidadeMedidasDo) Session(config *gorm.Session) IUnidadeMedidasDo {
	return u.withDO(u.DO.Session(config))
}

func (u unidadeMedidasDo) Clauses(conds ...clause.Expression) IUnidadeMedidasDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u unidadeMedidasDo) Returning(value interface{}, columns ...string) IUnidadeMedidasDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u unidadeMedidasDo) Not(conds ...gen.Condition) IUnidadeMedidasDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u unidadeMedidasDo) Or(conds ...gen.Condition) IUnidadeMedidasDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u unidadeMedidasDo) Select(conds ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u unidadeMedidasDo) Where(conds ...gen.Condition) IUnidadeMedidasDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u unidadeMedidasDo) Order(conds ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u unidadeMedidasDo) Distinct(cols ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u unidadeMedidasDo) Omit(cols ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u unidadeMedidasDo) Join(table schema.Tabler, on ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u unidadeMedidasDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u unidadeMedidasDo) RightJoin(table schema.Tabler, on ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u unidadeMedidasDo) Group(cols ...field.Expr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u unidadeMedidasDo) Having(conds ...gen.Condition) IUnidadeMedidasDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u unidadeMedidasDo) Limit(limit int) IUnidadeMedidasDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u unidadeMedidasDo) Offset(offset int) IUnidadeMedidasDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u unidadeMedidasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUnidadeMedidasDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u unidadeMedidasDo) Unscoped() IUnidadeMedidasDo {
	return u.withDO(u.DO.Unscoped())
}

func (u unidadeMedidasDo) Create(values ...*models.UnidadeMedidas) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u unidadeMedidasDo) CreateInBatches(values []*models.UnidadeMedidas, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u unidadeMedidasDo) Save(values ...*models.UnidadeMedidas) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u unidadeMedidasDo) First() (*models.UnidadeMedidas, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.UnidadeMedidas), nil
	}
}

func (u unidadeMedidasDo) Take() (*models.UnidadeMedidas, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.UnidadeMedidas), nil
	}
}

func (u unidadeMedidasDo) Last() (*models.UnidadeMedidas, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.UnidadeMedidas), nil
	}
}

func (u unidadeMedidasDo) Find() ([]*models.UnidadeMedidas, error) {
	result, err := u.DO.Find()
	return result.([]*models.UnidadeMedidas), err
}

func (u unidadeMedidasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UnidadeMedidas, err error) {
	buf := make([]*models.UnidadeMedidas, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u unidadeMedidasDo) FindInBatches(result *[]*models.UnidadeMedidas, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u unidadeMedidasDo) Attrs(attrs ...field.AssignExpr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u unidadeMedidasDo) Assign(attrs ...field.AssignExpr) IUnidadeMedidasDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u unidadeMedidasDo) Joins(fields ...field.RelationField) IUnidadeMedidasDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u unidadeMedidasDo) Preload(fields ...field.RelationField) IUnidadeMedidasDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u unidadeMedidasDo) FirstOrInit() (*models.UnidadeMedidas, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.UnidadeMedidas), nil
	}
}

func (u unidadeMedidasDo) FirstOrCreate() (*models.UnidadeMedidas, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.UnidadeMedidas), nil
	}
}

func (u unidadeMedidasDo) FindByPage(offset int, limit int) (result []*models.UnidadeMedidas, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u unidadeMedidasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u unidadeMedidasDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u unidadeMedidasDo) Delete(models ...*models.UnidadeMedidas) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *unidadeMedidasDo) withDO(do gen.Dao) *unidadeMedidasDo {
	u.DO = *do.(*gen.DO)
	return u
}
