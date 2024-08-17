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

func newUsuarios(db *gorm.DB, opts ...gen.DOOption) usuarios {
	_usuarios := usuarios{}

	_usuarios.usuariosDo.UseDB(db, opts...)
	_usuarios.usuariosDo.UseModel(&model.Usuarios{})

	tableName := _usuarios.usuariosDo.TableName()
	_usuarios.ALL = field.NewAsterisk(tableName)
	_usuarios.UsuarioID = field.NewInt32(tableName, "usuario_id")
	_usuarios.UsuarioNome = field.NewString(tableName, "usuario_nome")
	_usuarios.UsuarioSenha = field.NewString(tableName, "usuario_senha")
	_usuarios.UsuarioEmail = field.NewString(tableName, "usuario_email")

	_usuarios.fillFieldMap()

	return _usuarios
}

type usuarios struct {
	usuariosDo

	ALL          field.Asterisk
	UsuarioID    field.Int32
	UsuarioNome  field.String
	UsuarioSenha field.String
	UsuarioEmail field.String

	fieldMap map[string]field.Expr
}

func (u usuarios) Table(newTableName string) *usuarios {
	u.usuariosDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u usuarios) As(alias string) *usuarios {
	u.usuariosDo.DO = *(u.usuariosDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *usuarios) updateTableName(table string) *usuarios {
	u.ALL = field.NewAsterisk(table)
	u.UsuarioID = field.NewInt32(table, "usuario_id")
	u.UsuarioNome = field.NewString(table, "usuario_nome")
	u.UsuarioSenha = field.NewString(table, "usuario_senha")
	u.UsuarioEmail = field.NewString(table, "usuario_email")

	u.fillFieldMap()

	return u
}

func (u *usuarios) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *usuarios) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["usuario_id"] = u.UsuarioID
	u.fieldMap["usuario_nome"] = u.UsuarioNome
	u.fieldMap["usuario_senha"] = u.UsuarioSenha
	u.fieldMap["usuario_email"] = u.UsuarioEmail
}

func (u usuarios) clone(db *gorm.DB) usuarios {
	u.usuariosDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u usuarios) replaceDB(db *gorm.DB) usuarios {
	u.usuariosDo.ReplaceDB(db)
	return u
}

type usuariosDo struct{ gen.DO }

type IUsuariosDo interface {
	gen.SubQuery
	Debug() IUsuariosDo
	WithContext(ctx context.Context) IUsuariosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUsuariosDo
	WriteDB() IUsuariosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUsuariosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUsuariosDo
	Not(conds ...gen.Condition) IUsuariosDo
	Or(conds ...gen.Condition) IUsuariosDo
	Select(conds ...field.Expr) IUsuariosDo
	Where(conds ...gen.Condition) IUsuariosDo
	Order(conds ...field.Expr) IUsuariosDo
	Distinct(cols ...field.Expr) IUsuariosDo
	Omit(cols ...field.Expr) IUsuariosDo
	Join(table schema.Tabler, on ...field.Expr) IUsuariosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUsuariosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUsuariosDo
	Group(cols ...field.Expr) IUsuariosDo
	Having(conds ...gen.Condition) IUsuariosDo
	Limit(limit int) IUsuariosDo
	Offset(offset int) IUsuariosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUsuariosDo
	Unscoped() IUsuariosDo
	Create(values ...*model.Usuarios) error
	CreateInBatches(values []*model.Usuarios, batchSize int) error
	Save(values ...*model.Usuarios) error
	First() (*model.Usuarios, error)
	Take() (*model.Usuarios, error)
	Last() (*model.Usuarios, error)
	Find() ([]*model.Usuarios, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Usuarios, err error)
	FindInBatches(result *[]*model.Usuarios, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Usuarios) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUsuariosDo
	Assign(attrs ...field.AssignExpr) IUsuariosDo
	Joins(fields ...field.RelationField) IUsuariosDo
	Preload(fields ...field.RelationField) IUsuariosDo
	FirstOrInit() (*model.Usuarios, error)
	FirstOrCreate() (*model.Usuarios, error)
	FindByPage(offset int, limit int) (result []*model.Usuarios, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUsuariosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u usuariosDo) Debug() IUsuariosDo {
	return u.withDO(u.DO.Debug())
}

func (u usuariosDo) WithContext(ctx context.Context) IUsuariosDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u usuariosDo) ReadDB() IUsuariosDo {
	return u.Clauses(dbresolver.Read)
}

func (u usuariosDo) WriteDB() IUsuariosDo {
	return u.Clauses(dbresolver.Write)
}

func (u usuariosDo) Session(config *gorm.Session) IUsuariosDo {
	return u.withDO(u.DO.Session(config))
}

func (u usuariosDo) Clauses(conds ...clause.Expression) IUsuariosDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u usuariosDo) Returning(value interface{}, columns ...string) IUsuariosDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u usuariosDo) Not(conds ...gen.Condition) IUsuariosDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u usuariosDo) Or(conds ...gen.Condition) IUsuariosDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u usuariosDo) Select(conds ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u usuariosDo) Where(conds ...gen.Condition) IUsuariosDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u usuariosDo) Order(conds ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u usuariosDo) Distinct(cols ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u usuariosDo) Omit(cols ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u usuariosDo) Join(table schema.Tabler, on ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u usuariosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u usuariosDo) RightJoin(table schema.Tabler, on ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u usuariosDo) Group(cols ...field.Expr) IUsuariosDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u usuariosDo) Having(conds ...gen.Condition) IUsuariosDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u usuariosDo) Limit(limit int) IUsuariosDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u usuariosDo) Offset(offset int) IUsuariosDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u usuariosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUsuariosDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u usuariosDo) Unscoped() IUsuariosDo {
	return u.withDO(u.DO.Unscoped())
}

func (u usuariosDo) Create(values ...*model.Usuarios) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u usuariosDo) CreateInBatches(values []*model.Usuarios, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u usuariosDo) Save(values ...*model.Usuarios) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u usuariosDo) First() (*model.Usuarios, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Usuarios), nil
	}
}

func (u usuariosDo) Take() (*model.Usuarios, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Usuarios), nil
	}
}

func (u usuariosDo) Last() (*model.Usuarios, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Usuarios), nil
	}
}

func (u usuariosDo) Find() ([]*model.Usuarios, error) {
	result, err := u.DO.Find()
	return result.([]*model.Usuarios), err
}

func (u usuariosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Usuarios, err error) {
	buf := make([]*model.Usuarios, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u usuariosDo) FindInBatches(result *[]*model.Usuarios, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u usuariosDo) Attrs(attrs ...field.AssignExpr) IUsuariosDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u usuariosDo) Assign(attrs ...field.AssignExpr) IUsuariosDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u usuariosDo) Joins(fields ...field.RelationField) IUsuariosDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u usuariosDo) Preload(fields ...field.RelationField) IUsuariosDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u usuariosDo) FirstOrInit() (*model.Usuarios, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Usuarios), nil
	}
}

func (u usuariosDo) FirstOrCreate() (*model.Usuarios, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Usuarios), nil
	}
}

func (u usuariosDo) FindByPage(offset int, limit int) (result []*model.Usuarios, count int64, err error) {
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

func (u usuariosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u usuariosDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u usuariosDo) Delete(models ...*model.Usuarios) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *usuariosDo) withDO(do gen.Dao) *usuariosDo {
	u.DO = *do.(*gen.DO)
	return u
}
