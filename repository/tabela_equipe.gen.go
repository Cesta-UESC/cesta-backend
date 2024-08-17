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

func newEquipe(db *gorm.DB, opts ...gen.DOOption) equipe {
	_equipe := equipe{}

	_equipe.equipeDo.UseDB(db, opts...)
	_equipe.equipeDo.UseModel(&model.Equipe{})

	tableName := _equipe.equipeDo.TableName()
	_equipe.ALL = field.NewAsterisk(tableName)
	_equipe.ID = field.NewInt32(tableName, "id")
	_equipe.NomeCompleto = field.NewString(tableName, "nome_completo")
	_equipe.Email = field.NewString(tableName, "email")
	_equipe.FuncaoID = field.NewInt32(tableName, "funcao_id")
	_equipe.MostrarHome = field.NewBool(tableName, "mostrar_home")
	_equipe.MostrarContatos = field.NewBool(tableName, "mostrar_contatos")

	_equipe.fillFieldMap()

	return _equipe
}

type equipe struct {
	equipeDo

	ALL             field.Asterisk
	ID              field.Int32
	NomeCompleto    field.String
	Email           field.String
	FuncaoID        field.Int32
	MostrarHome     field.Bool
	MostrarContatos field.Bool

	fieldMap map[string]field.Expr
}

func (e equipe) Table(newTableName string) *equipe {
	e.equipeDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipe) As(alias string) *equipe {
	e.equipeDo.DO = *(e.equipeDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipe) updateTableName(table string) *equipe {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.NomeCompleto = field.NewString(table, "nome_completo")
	e.Email = field.NewString(table, "email")
	e.FuncaoID = field.NewInt32(table, "funcao_id")
	e.MostrarHome = field.NewBool(table, "mostrar_home")
	e.MostrarContatos = field.NewBool(table, "mostrar_contatos")

	e.fillFieldMap()

	return e
}

func (e *equipe) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipe) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["id"] = e.ID
	e.fieldMap["nome_completo"] = e.NomeCompleto
	e.fieldMap["email"] = e.Email
	e.fieldMap["funcao_id"] = e.FuncaoID
	e.fieldMap["mostrar_home"] = e.MostrarHome
	e.fieldMap["mostrar_contatos"] = e.MostrarContatos
}

func (e equipe) clone(db *gorm.DB) equipe {
	e.equipeDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipe) replaceDB(db *gorm.DB) equipe {
	e.equipeDo.ReplaceDB(db)
	return e
}

type equipeDo struct{ gen.DO }

type IEquipeDo interface {
	gen.SubQuery
	Debug() IEquipeDo
	WithContext(ctx context.Context) IEquipeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEquipeDo
	WriteDB() IEquipeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEquipeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEquipeDo
	Not(conds ...gen.Condition) IEquipeDo
	Or(conds ...gen.Condition) IEquipeDo
	Select(conds ...field.Expr) IEquipeDo
	Where(conds ...gen.Condition) IEquipeDo
	Order(conds ...field.Expr) IEquipeDo
	Distinct(cols ...field.Expr) IEquipeDo
	Omit(cols ...field.Expr) IEquipeDo
	Join(table schema.Tabler, on ...field.Expr) IEquipeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEquipeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEquipeDo
	Group(cols ...field.Expr) IEquipeDo
	Having(conds ...gen.Condition) IEquipeDo
	Limit(limit int) IEquipeDo
	Offset(offset int) IEquipeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEquipeDo
	Unscoped() IEquipeDo
	Create(values ...*model.Equipe) error
	CreateInBatches(values []*model.Equipe, batchSize int) error
	Save(values ...*model.Equipe) error
	First() (*model.Equipe, error)
	Take() (*model.Equipe, error)
	Last() (*model.Equipe, error)
	Find() ([]*model.Equipe, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Equipe, err error)
	FindInBatches(result *[]*model.Equipe, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Equipe) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEquipeDo
	Assign(attrs ...field.AssignExpr) IEquipeDo
	Joins(fields ...field.RelationField) IEquipeDo
	Preload(fields ...field.RelationField) IEquipeDo
	FirstOrInit() (*model.Equipe, error)
	FirstOrCreate() (*model.Equipe, error)
	FindByPage(offset int, limit int) (result []*model.Equipe, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEquipeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e equipeDo) Debug() IEquipeDo {
	return e.withDO(e.DO.Debug())
}

func (e equipeDo) WithContext(ctx context.Context) IEquipeDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipeDo) ReadDB() IEquipeDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipeDo) WriteDB() IEquipeDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipeDo) Session(config *gorm.Session) IEquipeDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipeDo) Clauses(conds ...clause.Expression) IEquipeDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipeDo) Returning(value interface{}, columns ...string) IEquipeDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipeDo) Not(conds ...gen.Condition) IEquipeDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipeDo) Or(conds ...gen.Condition) IEquipeDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipeDo) Select(conds ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipeDo) Where(conds ...gen.Condition) IEquipeDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipeDo) Order(conds ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipeDo) Distinct(cols ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipeDo) Omit(cols ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipeDo) Join(table schema.Tabler, on ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipeDo) RightJoin(table schema.Tabler, on ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipeDo) Group(cols ...field.Expr) IEquipeDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipeDo) Having(conds ...gen.Condition) IEquipeDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipeDo) Limit(limit int) IEquipeDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipeDo) Offset(offset int) IEquipeDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEquipeDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipeDo) Unscoped() IEquipeDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipeDo) Create(values ...*model.Equipe) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipeDo) CreateInBatches(values []*model.Equipe, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipeDo) Save(values ...*model.Equipe) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipeDo) First() (*model.Equipe, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Equipe), nil
	}
}

func (e equipeDo) Take() (*model.Equipe, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Equipe), nil
	}
}

func (e equipeDo) Last() (*model.Equipe, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Equipe), nil
	}
}

func (e equipeDo) Find() ([]*model.Equipe, error) {
	result, err := e.DO.Find()
	return result.([]*model.Equipe), err
}

func (e equipeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Equipe, err error) {
	buf := make([]*model.Equipe, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipeDo) FindInBatches(result *[]*model.Equipe, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipeDo) Attrs(attrs ...field.AssignExpr) IEquipeDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipeDo) Assign(attrs ...field.AssignExpr) IEquipeDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipeDo) Joins(fields ...field.RelationField) IEquipeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipeDo) Preload(fields ...field.RelationField) IEquipeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipeDo) FirstOrInit() (*model.Equipe, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Equipe), nil
	}
}

func (e equipeDo) FirstOrCreate() (*model.Equipe, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Equipe), nil
	}
}

func (e equipeDo) FindByPage(offset int, limit int) (result []*model.Equipe, count int64, err error) {
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

func (e equipeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipeDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipeDo) Delete(models ...*model.Equipe) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipeDo) withDO(do gen.Dao) *equipeDo {
	e.DO = *do.(*gen.DO)
	return e
}
