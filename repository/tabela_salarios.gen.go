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

func newSalarios(db *gorm.DB, opts ...gen.DOOption) salarios {
	_salarios := salarios{}

	_salarios.salariosDo.UseDB(db, opts...)
	_salarios.salariosDo.UseModel(&model.Salarios{})

	tableName := _salarios.salariosDo.TableName()
	_salarios.ALL = field.NewAsterisk(tableName)
	_salarios.SalarioID = field.NewInt32(tableName, "salario_id")
	_salarios.SalarioValorBruto = field.NewFloat64(tableName, "salario_valor_bruto")
	_salarios.SalarioValorLiquido = field.NewFloat64(tableName, "salario_valor_liquido")
	_salarios.SalarioEmUso = field.NewInt32(tableName, "salario_em_uso")
	_salarios.SalarioDataRegistro = field.NewTime(tableName, "salario_data_registro")
	_salarios.SalarioNome = field.NewString(tableName, "salario_nome")
	_salarios.SalarioSimbolo = field.NewString(tableName, "salario_simbolo")

	_salarios.fillFieldMap()

	return _salarios
}

type salarios struct {
	salariosDo

	ALL                 field.Asterisk
	SalarioID           field.Int32
	SalarioValorBruto   field.Float64
	SalarioValorLiquido field.Float64
	SalarioEmUso        field.Int32
	SalarioDataRegistro field.Time
	SalarioNome         field.String
	SalarioSimbolo      field.String

	fieldMap map[string]field.Expr
}

func (s salarios) Table(newTableName string) *salarios {
	s.salariosDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s salarios) As(alias string) *salarios {
	s.salariosDo.DO = *(s.salariosDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *salarios) updateTableName(table string) *salarios {
	s.ALL = field.NewAsterisk(table)
	s.SalarioID = field.NewInt32(table, "salario_id")
	s.SalarioValorBruto = field.NewFloat64(table, "salario_valor_bruto")
	s.SalarioValorLiquido = field.NewFloat64(table, "salario_valor_liquido")
	s.SalarioEmUso = field.NewInt32(table, "salario_em_uso")
	s.SalarioDataRegistro = field.NewTime(table, "salario_data_registro")
	s.SalarioNome = field.NewString(table, "salario_nome")
	s.SalarioSimbolo = field.NewString(table, "salario_simbolo")

	s.fillFieldMap()

	return s
}

func (s *salarios) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *salarios) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["salario_id"] = s.SalarioID
	s.fieldMap["salario_valor_bruto"] = s.SalarioValorBruto
	s.fieldMap["salario_valor_liquido"] = s.SalarioValorLiquido
	s.fieldMap["salario_em_uso"] = s.SalarioEmUso
	s.fieldMap["salario_data_registro"] = s.SalarioDataRegistro
	s.fieldMap["salario_nome"] = s.SalarioNome
	s.fieldMap["salario_simbolo"] = s.SalarioSimbolo
}

func (s salarios) clone(db *gorm.DB) salarios {
	s.salariosDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s salarios) replaceDB(db *gorm.DB) salarios {
	s.salariosDo.ReplaceDB(db)
	return s
}

type salariosDo struct{ gen.DO }

type ISalariosDo interface {
	gen.SubQuery
	Debug() ISalariosDo
	WithContext(ctx context.Context) ISalariosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISalariosDo
	WriteDB() ISalariosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISalariosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISalariosDo
	Not(conds ...gen.Condition) ISalariosDo
	Or(conds ...gen.Condition) ISalariosDo
	Select(conds ...field.Expr) ISalariosDo
	Where(conds ...gen.Condition) ISalariosDo
	Order(conds ...field.Expr) ISalariosDo
	Distinct(cols ...field.Expr) ISalariosDo
	Omit(cols ...field.Expr) ISalariosDo
	Join(table schema.Tabler, on ...field.Expr) ISalariosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISalariosDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISalariosDo
	Group(cols ...field.Expr) ISalariosDo
	Having(conds ...gen.Condition) ISalariosDo
	Limit(limit int) ISalariosDo
	Offset(offset int) ISalariosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISalariosDo
	Unscoped() ISalariosDo
	Create(values ...*model.Salarios) error
	CreateInBatches(values []*model.Salarios, batchSize int) error
	Save(values ...*model.Salarios) error
	First() (*model.Salarios, error)
	Take() (*model.Salarios, error)
	Last() (*model.Salarios, error)
	Find() ([]*model.Salarios, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Salarios, err error)
	FindInBatches(result *[]*model.Salarios, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Salarios) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISalariosDo
	Assign(attrs ...field.AssignExpr) ISalariosDo
	Joins(fields ...field.RelationField) ISalariosDo
	Preload(fields ...field.RelationField) ISalariosDo
	FirstOrInit() (*model.Salarios, error)
	FirstOrCreate() (*model.Salarios, error)
	FindByPage(offset int, limit int) (result []*model.Salarios, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISalariosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s salariosDo) Debug() ISalariosDo {
	return s.withDO(s.DO.Debug())
}

func (s salariosDo) WithContext(ctx context.Context) ISalariosDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s salariosDo) ReadDB() ISalariosDo {
	return s.Clauses(dbresolver.Read)
}

func (s salariosDo) WriteDB() ISalariosDo {
	return s.Clauses(dbresolver.Write)
}

func (s salariosDo) Session(config *gorm.Session) ISalariosDo {
	return s.withDO(s.DO.Session(config))
}

func (s salariosDo) Clauses(conds ...clause.Expression) ISalariosDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s salariosDo) Returning(value interface{}, columns ...string) ISalariosDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s salariosDo) Not(conds ...gen.Condition) ISalariosDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s salariosDo) Or(conds ...gen.Condition) ISalariosDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s salariosDo) Select(conds ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s salariosDo) Where(conds ...gen.Condition) ISalariosDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s salariosDo) Order(conds ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s salariosDo) Distinct(cols ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s salariosDo) Omit(cols ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s salariosDo) Join(table schema.Tabler, on ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s salariosDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s salariosDo) RightJoin(table schema.Tabler, on ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s salariosDo) Group(cols ...field.Expr) ISalariosDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s salariosDo) Having(conds ...gen.Condition) ISalariosDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s salariosDo) Limit(limit int) ISalariosDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s salariosDo) Offset(offset int) ISalariosDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s salariosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISalariosDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s salariosDo) Unscoped() ISalariosDo {
	return s.withDO(s.DO.Unscoped())
}

func (s salariosDo) Create(values ...*model.Salarios) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s salariosDo) CreateInBatches(values []*model.Salarios, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s salariosDo) Save(values ...*model.Salarios) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s salariosDo) First() (*model.Salarios, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salarios), nil
	}
}

func (s salariosDo) Take() (*model.Salarios, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salarios), nil
	}
}

func (s salariosDo) Last() (*model.Salarios, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salarios), nil
	}
}

func (s salariosDo) Find() ([]*model.Salarios, error) {
	result, err := s.DO.Find()
	return result.([]*model.Salarios), err
}

func (s salariosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Salarios, err error) {
	buf := make([]*model.Salarios, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s salariosDo) FindInBatches(result *[]*model.Salarios, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s salariosDo) Attrs(attrs ...field.AssignExpr) ISalariosDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s salariosDo) Assign(attrs ...field.AssignExpr) ISalariosDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s salariosDo) Joins(fields ...field.RelationField) ISalariosDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s salariosDo) Preload(fields ...field.RelationField) ISalariosDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s salariosDo) FirstOrInit() (*model.Salarios, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salarios), nil
	}
}

func (s salariosDo) FirstOrCreate() (*model.Salarios, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salarios), nil
	}
}

func (s salariosDo) FindByPage(offset int, limit int) (result []*model.Salarios, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s salariosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s salariosDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s salariosDo) Delete(models ...*model.Salarios) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *salariosDo) withDO(do gen.Dao) *salariosDo {
	s.DO = *do.(*gen.DO)
	return s
}
