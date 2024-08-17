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

func newSessao(db *gorm.DB, opts ...gen.DOOption) sessao {
	_sessao := sessao{}

	_sessao.sessaoDo.UseDB(db, opts...)
	_sessao.sessaoDo.UseModel(&model.Sessao{})

	tableName := _sessao.sessaoDo.TableName()
	_sessao.ALL = field.NewAsterisk(tableName)
	_sessao.SessaoID = field.NewString(tableName, "sessao_id")
	_sessao.SessaoUsuario = field.NewInt32(tableName, "sessao_usuario")
	_sessao.SessaoIP = field.NewString(tableName, "sessao_ip")
	_sessao.SessaoTempo = field.NewInt64(tableName, "sessao_tempo")

	_sessao.fillFieldMap()

	return _sessao
}

type sessao struct {
	sessaoDo

	ALL           field.Asterisk
	SessaoID      field.String
	SessaoUsuario field.Int32
	SessaoIP      field.String
	SessaoTempo   field.Int64

	fieldMap map[string]field.Expr
}

func (s sessao) Table(newTableName string) *sessao {
	s.sessaoDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sessao) As(alias string) *sessao {
	s.sessaoDo.DO = *(s.sessaoDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sessao) updateTableName(table string) *sessao {
	s.ALL = field.NewAsterisk(table)
	s.SessaoID = field.NewString(table, "sessao_id")
	s.SessaoUsuario = field.NewInt32(table, "sessao_usuario")
	s.SessaoIP = field.NewString(table, "sessao_ip")
	s.SessaoTempo = field.NewInt64(table, "sessao_tempo")

	s.fillFieldMap()

	return s
}

func (s *sessao) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sessao) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["sessao_id"] = s.SessaoID
	s.fieldMap["sessao_usuario"] = s.SessaoUsuario
	s.fieldMap["sessao_ip"] = s.SessaoIP
	s.fieldMap["sessao_tempo"] = s.SessaoTempo
}

func (s sessao) clone(db *gorm.DB) sessao {
	s.sessaoDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sessao) replaceDB(db *gorm.DB) sessao {
	s.sessaoDo.ReplaceDB(db)
	return s
}

type sessaoDo struct{ gen.DO }

type ISessaoDo interface {
	gen.SubQuery
	Debug() ISessaoDo
	WithContext(ctx context.Context) ISessaoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISessaoDo
	WriteDB() ISessaoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISessaoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISessaoDo
	Not(conds ...gen.Condition) ISessaoDo
	Or(conds ...gen.Condition) ISessaoDo
	Select(conds ...field.Expr) ISessaoDo
	Where(conds ...gen.Condition) ISessaoDo
	Order(conds ...field.Expr) ISessaoDo
	Distinct(cols ...field.Expr) ISessaoDo
	Omit(cols ...field.Expr) ISessaoDo
	Join(table schema.Tabler, on ...field.Expr) ISessaoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISessaoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISessaoDo
	Group(cols ...field.Expr) ISessaoDo
	Having(conds ...gen.Condition) ISessaoDo
	Limit(limit int) ISessaoDo
	Offset(offset int) ISessaoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISessaoDo
	Unscoped() ISessaoDo
	Create(values ...*model.Sessao) error
	CreateInBatches(values []*model.Sessao, batchSize int) error
	Save(values ...*model.Sessao) error
	First() (*model.Sessao, error)
	Take() (*model.Sessao, error)
	Last() (*model.Sessao, error)
	Find() ([]*model.Sessao, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sessao, err error)
	FindInBatches(result *[]*model.Sessao, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Sessao) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISessaoDo
	Assign(attrs ...field.AssignExpr) ISessaoDo
	Joins(fields ...field.RelationField) ISessaoDo
	Preload(fields ...field.RelationField) ISessaoDo
	FirstOrInit() (*model.Sessao, error)
	FirstOrCreate() (*model.Sessao, error)
	FindByPage(offset int, limit int) (result []*model.Sessao, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISessaoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sessaoDo) Debug() ISessaoDo {
	return s.withDO(s.DO.Debug())
}

func (s sessaoDo) WithContext(ctx context.Context) ISessaoDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sessaoDo) ReadDB() ISessaoDo {
	return s.Clauses(dbresolver.Read)
}

func (s sessaoDo) WriteDB() ISessaoDo {
	return s.Clauses(dbresolver.Write)
}

func (s sessaoDo) Session(config *gorm.Session) ISessaoDo {
	return s.withDO(s.DO.Session(config))
}

func (s sessaoDo) Clauses(conds ...clause.Expression) ISessaoDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sessaoDo) Returning(value interface{}, columns ...string) ISessaoDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sessaoDo) Not(conds ...gen.Condition) ISessaoDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sessaoDo) Or(conds ...gen.Condition) ISessaoDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sessaoDo) Select(conds ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sessaoDo) Where(conds ...gen.Condition) ISessaoDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sessaoDo) Order(conds ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sessaoDo) Distinct(cols ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sessaoDo) Omit(cols ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sessaoDo) Join(table schema.Tabler, on ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sessaoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sessaoDo) RightJoin(table schema.Tabler, on ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sessaoDo) Group(cols ...field.Expr) ISessaoDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sessaoDo) Having(conds ...gen.Condition) ISessaoDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sessaoDo) Limit(limit int) ISessaoDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sessaoDo) Offset(offset int) ISessaoDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sessaoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISessaoDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sessaoDo) Unscoped() ISessaoDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sessaoDo) Create(values ...*model.Sessao) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sessaoDo) CreateInBatches(values []*model.Sessao, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sessaoDo) Save(values ...*model.Sessao) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sessaoDo) First() (*model.Sessao, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sessao), nil
	}
}

func (s sessaoDo) Take() (*model.Sessao, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sessao), nil
	}
}

func (s sessaoDo) Last() (*model.Sessao, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sessao), nil
	}
}

func (s sessaoDo) Find() ([]*model.Sessao, error) {
	result, err := s.DO.Find()
	return result.([]*model.Sessao), err
}

func (s sessaoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sessao, err error) {
	buf := make([]*model.Sessao, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sessaoDo) FindInBatches(result *[]*model.Sessao, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sessaoDo) Attrs(attrs ...field.AssignExpr) ISessaoDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sessaoDo) Assign(attrs ...field.AssignExpr) ISessaoDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sessaoDo) Joins(fields ...field.RelationField) ISessaoDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sessaoDo) Preload(fields ...field.RelationField) ISessaoDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sessaoDo) FirstOrInit() (*model.Sessao, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sessao), nil
	}
}

func (s sessaoDo) FirstOrCreate() (*model.Sessao, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sessao), nil
	}
}

func (s sessaoDo) FindByPage(offset int, limit int) (result []*model.Sessao, count int64, err error) {
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

func (s sessaoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sessaoDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sessaoDo) Delete(models ...*model.Sessao) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sessaoDo) withDO(do gen.Dao) *sessaoDo {
	s.DO = *do.(*gen.DO)
	return s
}
