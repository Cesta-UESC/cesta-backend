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

func newAuxiliarCronograma(db *gorm.DB, opts ...gen.DOOption) auxiliarCronograma {
	_auxiliarCronograma := auxiliarCronograma{}

	_auxiliarCronograma.auxiliarCronogramaDo.UseDB(db, opts...)
	_auxiliarCronograma.auxiliarCronogramaDo.UseModel(&model.AuxiliarCronograma{})

	tableName := _auxiliarCronograma.auxiliarCronogramaDo.TableName()
	_auxiliarCronograma.ALL = field.NewAsterisk(tableName)
	_auxiliarCronograma.CronogramaID = field.NewInt32(tableName, "cronograma_id")
	_auxiliarCronograma.MesID = field.NewInt32(tableName, "mes_id")
	_auxiliarCronograma.InicioColeta = field.NewTime(tableName, "inicio_coleta")
	_auxiliarCronograma.FimColeta = field.NewTime(tableName, "fim_coleta")

	_auxiliarCronograma.fillFieldMap()

	return _auxiliarCronograma
}

type auxiliarCronograma struct {
	auxiliarCronogramaDo

	ALL          field.Asterisk
	CronogramaID field.Int32
	MesID        field.Int32
	InicioColeta field.Time
	FimColeta    field.Time

	fieldMap map[string]field.Expr
}

func (a auxiliarCronograma) Table(newTableName string) *auxiliarCronograma {
	a.auxiliarCronogramaDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a auxiliarCronograma) As(alias string) *auxiliarCronograma {
	a.auxiliarCronogramaDo.DO = *(a.auxiliarCronogramaDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *auxiliarCronograma) updateTableName(table string) *auxiliarCronograma {
	a.ALL = field.NewAsterisk(table)
	a.CronogramaID = field.NewInt32(table, "cronograma_id")
	a.MesID = field.NewInt32(table, "mes_id")
	a.InicioColeta = field.NewTime(table, "inicio_coleta")
	a.FimColeta = field.NewTime(table, "fim_coleta")

	a.fillFieldMap()

	return a
}

func (a *auxiliarCronograma) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *auxiliarCronograma) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["cronograma_id"] = a.CronogramaID
	a.fieldMap["mes_id"] = a.MesID
	a.fieldMap["inicio_coleta"] = a.InicioColeta
	a.fieldMap["fim_coleta"] = a.FimColeta
}

func (a auxiliarCronograma) clone(db *gorm.DB) auxiliarCronograma {
	a.auxiliarCronogramaDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a auxiliarCronograma) replaceDB(db *gorm.DB) auxiliarCronograma {
	a.auxiliarCronogramaDo.ReplaceDB(db)
	return a
}

type auxiliarCronogramaDo struct{ gen.DO }

type IAuxiliarCronogramaDo interface {
	gen.SubQuery
	Debug() IAuxiliarCronogramaDo
	WithContext(ctx context.Context) IAuxiliarCronogramaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuxiliarCronogramaDo
	WriteDB() IAuxiliarCronogramaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuxiliarCronogramaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuxiliarCronogramaDo
	Not(conds ...gen.Condition) IAuxiliarCronogramaDo
	Or(conds ...gen.Condition) IAuxiliarCronogramaDo
	Select(conds ...field.Expr) IAuxiliarCronogramaDo
	Where(conds ...gen.Condition) IAuxiliarCronogramaDo
	Order(conds ...field.Expr) IAuxiliarCronogramaDo
	Distinct(cols ...field.Expr) IAuxiliarCronogramaDo
	Omit(cols ...field.Expr) IAuxiliarCronogramaDo
	Join(table schema.Tabler, on ...field.Expr) IAuxiliarCronogramaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuxiliarCronogramaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuxiliarCronogramaDo
	Group(cols ...field.Expr) IAuxiliarCronogramaDo
	Having(conds ...gen.Condition) IAuxiliarCronogramaDo
	Limit(limit int) IAuxiliarCronogramaDo
	Offset(offset int) IAuxiliarCronogramaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuxiliarCronogramaDo
	Unscoped() IAuxiliarCronogramaDo
	Create(values ...*model.AuxiliarCronograma) error
	CreateInBatches(values []*model.AuxiliarCronograma, batchSize int) error
	Save(values ...*model.AuxiliarCronograma) error
	First() (*model.AuxiliarCronograma, error)
	Take() (*model.AuxiliarCronograma, error)
	Last() (*model.AuxiliarCronograma, error)
	Find() ([]*model.AuxiliarCronograma, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuxiliarCronograma, err error)
	FindInBatches(result *[]*model.AuxiliarCronograma, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AuxiliarCronograma) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuxiliarCronogramaDo
	Assign(attrs ...field.AssignExpr) IAuxiliarCronogramaDo
	Joins(fields ...field.RelationField) IAuxiliarCronogramaDo
	Preload(fields ...field.RelationField) IAuxiliarCronogramaDo
	FirstOrInit() (*model.AuxiliarCronograma, error)
	FirstOrCreate() (*model.AuxiliarCronograma, error)
	FindByPage(offset int, limit int) (result []*model.AuxiliarCronograma, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuxiliarCronogramaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a auxiliarCronogramaDo) Debug() IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Debug())
}

func (a auxiliarCronogramaDo) WithContext(ctx context.Context) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a auxiliarCronogramaDo) ReadDB() IAuxiliarCronogramaDo {
	return a.Clauses(dbresolver.Read)
}

func (a auxiliarCronogramaDo) WriteDB() IAuxiliarCronogramaDo {
	return a.Clauses(dbresolver.Write)
}

func (a auxiliarCronogramaDo) Session(config *gorm.Session) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Session(config))
}

func (a auxiliarCronogramaDo) Clauses(conds ...clause.Expression) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a auxiliarCronogramaDo) Returning(value interface{}, columns ...string) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a auxiliarCronogramaDo) Not(conds ...gen.Condition) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a auxiliarCronogramaDo) Or(conds ...gen.Condition) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a auxiliarCronogramaDo) Select(conds ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a auxiliarCronogramaDo) Where(conds ...gen.Condition) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a auxiliarCronogramaDo) Order(conds ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a auxiliarCronogramaDo) Distinct(cols ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a auxiliarCronogramaDo) Omit(cols ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a auxiliarCronogramaDo) Join(table schema.Tabler, on ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a auxiliarCronogramaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a auxiliarCronogramaDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a auxiliarCronogramaDo) Group(cols ...field.Expr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a auxiliarCronogramaDo) Having(conds ...gen.Condition) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a auxiliarCronogramaDo) Limit(limit int) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a auxiliarCronogramaDo) Offset(offset int) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a auxiliarCronogramaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a auxiliarCronogramaDo) Unscoped() IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Unscoped())
}

func (a auxiliarCronogramaDo) Create(values ...*model.AuxiliarCronograma) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a auxiliarCronogramaDo) CreateInBatches(values []*model.AuxiliarCronograma, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a auxiliarCronogramaDo) Save(values ...*model.AuxiliarCronograma) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a auxiliarCronogramaDo) First() (*model.AuxiliarCronograma, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuxiliarCronograma), nil
	}
}

func (a auxiliarCronogramaDo) Take() (*model.AuxiliarCronograma, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuxiliarCronograma), nil
	}
}

func (a auxiliarCronogramaDo) Last() (*model.AuxiliarCronograma, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuxiliarCronograma), nil
	}
}

func (a auxiliarCronogramaDo) Find() ([]*model.AuxiliarCronograma, error) {
	result, err := a.DO.Find()
	return result.([]*model.AuxiliarCronograma), err
}

func (a auxiliarCronogramaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuxiliarCronograma, err error) {
	buf := make([]*model.AuxiliarCronograma, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a auxiliarCronogramaDo) FindInBatches(result *[]*model.AuxiliarCronograma, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a auxiliarCronogramaDo) Attrs(attrs ...field.AssignExpr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a auxiliarCronogramaDo) Assign(attrs ...field.AssignExpr) IAuxiliarCronogramaDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a auxiliarCronogramaDo) Joins(fields ...field.RelationField) IAuxiliarCronogramaDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a auxiliarCronogramaDo) Preload(fields ...field.RelationField) IAuxiliarCronogramaDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a auxiliarCronogramaDo) FirstOrInit() (*model.AuxiliarCronograma, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuxiliarCronograma), nil
	}
}

func (a auxiliarCronogramaDo) FirstOrCreate() (*model.AuxiliarCronograma, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuxiliarCronograma), nil
	}
}

func (a auxiliarCronogramaDo) FindByPage(offset int, limit int) (result []*model.AuxiliarCronograma, count int64, err error) {
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

func (a auxiliarCronogramaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a auxiliarCronogramaDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a auxiliarCronogramaDo) Delete(models ...*model.AuxiliarCronograma) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *auxiliarCronogramaDo) withDO(do gen.Dao) *auxiliarCronogramaDo {
	a.DO = *do.(*gen.DO)
	return a
}
