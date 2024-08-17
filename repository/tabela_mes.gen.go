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

func newMes(db *gorm.DB, opts ...gen.DOOption) mes {
	_mes := mes{}

	_mes.mesDo.UseDB(db, opts...)
	_mes.mesDo.UseModel(&model.Mes{})

	tableName := _mes.mesDo.TableName()
	_mes.ALL = field.NewAsterisk(tableName)
	_mes.MesID = field.NewInt32(tableName, "mes_id")
	_mes.MesNome = field.NewString(tableName, "mes_nome")

	_mes.fillFieldMap()

	return _mes
}

type mes struct {
	mesDo

	ALL     field.Asterisk
	MesID   field.Int32
	MesNome field.String

	fieldMap map[string]field.Expr
}

func (m mes) Table(newTableName string) *mes {
	m.mesDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mes) As(alias string) *mes {
	m.mesDo.DO = *(m.mesDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mes) updateTableName(table string) *mes {
	m.ALL = field.NewAsterisk(table)
	m.MesID = field.NewInt32(table, "mes_id")
	m.MesNome = field.NewString(table, "mes_nome")

	m.fillFieldMap()

	return m
}

func (m *mes) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mes) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 2)
	m.fieldMap["mes_id"] = m.MesID
	m.fieldMap["mes_nome"] = m.MesNome
}

func (m mes) clone(db *gorm.DB) mes {
	m.mesDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mes) replaceDB(db *gorm.DB) mes {
	m.mesDo.ReplaceDB(db)
	return m
}

type mesDo struct{ gen.DO }

type IMesDo interface {
	gen.SubQuery
	Debug() IMesDo
	WithContext(ctx context.Context) IMesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMesDo
	WriteDB() IMesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMesDo
	Not(conds ...gen.Condition) IMesDo
	Or(conds ...gen.Condition) IMesDo
	Select(conds ...field.Expr) IMesDo
	Where(conds ...gen.Condition) IMesDo
	Order(conds ...field.Expr) IMesDo
	Distinct(cols ...field.Expr) IMesDo
	Omit(cols ...field.Expr) IMesDo
	Join(table schema.Tabler, on ...field.Expr) IMesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMesDo
	Group(cols ...field.Expr) IMesDo
	Having(conds ...gen.Condition) IMesDo
	Limit(limit int) IMesDo
	Offset(offset int) IMesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMesDo
	Unscoped() IMesDo
	Create(values ...*model.Mes) error
	CreateInBatches(values []*model.Mes, batchSize int) error
	Save(values ...*model.Mes) error
	First() (*model.Mes, error)
	Take() (*model.Mes, error)
	Last() (*model.Mes, error)
	Find() ([]*model.Mes, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Mes, err error)
	FindInBatches(result *[]*model.Mes, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Mes) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMesDo
	Assign(attrs ...field.AssignExpr) IMesDo
	Joins(fields ...field.RelationField) IMesDo
	Preload(fields ...field.RelationField) IMesDo
	FirstOrInit() (*model.Mes, error)
	FirstOrCreate() (*model.Mes, error)
	FindByPage(offset int, limit int) (result []*model.Mes, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mesDo) Debug() IMesDo {
	return m.withDO(m.DO.Debug())
}

func (m mesDo) WithContext(ctx context.Context) IMesDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mesDo) ReadDB() IMesDo {
	return m.Clauses(dbresolver.Read)
}

func (m mesDo) WriteDB() IMesDo {
	return m.Clauses(dbresolver.Write)
}

func (m mesDo) Session(config *gorm.Session) IMesDo {
	return m.withDO(m.DO.Session(config))
}

func (m mesDo) Clauses(conds ...clause.Expression) IMesDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mesDo) Returning(value interface{}, columns ...string) IMesDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mesDo) Not(conds ...gen.Condition) IMesDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mesDo) Or(conds ...gen.Condition) IMesDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mesDo) Select(conds ...field.Expr) IMesDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mesDo) Where(conds ...gen.Condition) IMesDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mesDo) Order(conds ...field.Expr) IMesDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mesDo) Distinct(cols ...field.Expr) IMesDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mesDo) Omit(cols ...field.Expr) IMesDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mesDo) Join(table schema.Tabler, on ...field.Expr) IMesDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMesDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mesDo) RightJoin(table schema.Tabler, on ...field.Expr) IMesDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mesDo) Group(cols ...field.Expr) IMesDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mesDo) Having(conds ...gen.Condition) IMesDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mesDo) Limit(limit int) IMesDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mesDo) Offset(offset int) IMesDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMesDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mesDo) Unscoped() IMesDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mesDo) Create(values ...*model.Mes) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mesDo) CreateInBatches(values []*model.Mes, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mesDo) Save(values ...*model.Mes) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mesDo) First() (*model.Mes, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mes), nil
	}
}

func (m mesDo) Take() (*model.Mes, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mes), nil
	}
}

func (m mesDo) Last() (*model.Mes, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mes), nil
	}
}

func (m mesDo) Find() ([]*model.Mes, error) {
	result, err := m.DO.Find()
	return result.([]*model.Mes), err
}

func (m mesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Mes, err error) {
	buf := make([]*model.Mes, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mesDo) FindInBatches(result *[]*model.Mes, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mesDo) Attrs(attrs ...field.AssignExpr) IMesDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mesDo) Assign(attrs ...field.AssignExpr) IMesDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mesDo) Joins(fields ...field.RelationField) IMesDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mesDo) Preload(fields ...field.RelationField) IMesDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mesDo) FirstOrInit() (*model.Mes, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mes), nil
	}
}

func (m mesDo) FirstOrCreate() (*model.Mes, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mes), nil
	}
}

func (m mesDo) FindByPage(offset int, limit int) (result []*model.Mes, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m mesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mesDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mesDo) Delete(models ...*model.Mes) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mesDo) withDO(do gen.Dao) *mesDo {
	m.DO = *do.(*gen.DO)
	return m
}
