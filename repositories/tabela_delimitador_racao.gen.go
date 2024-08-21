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

func newDelimitadorRacao(db *gorm.DB, opts ...gen.DOOption) delimitadorRacao {
	_delimitadorRacao := delimitadorRacao{}

	_delimitadorRacao.delimitadorRacaoDo.UseDB(db, opts...)
	_delimitadorRacao.delimitadorRacaoDo.UseModel(&models.DelimitadorRacao{})

	tableName := _delimitadorRacao.delimitadorRacaoDo.TableName()
	_delimitadorRacao.ALL = field.NewAsterisk(tableName)
	_delimitadorRacao.DelimitadorID = field.NewInt32(tableName, "delimitador_id")
	_delimitadorRacao.DelimitadorDescricao = field.NewString(tableName, "delimitador_descricao")
	_delimitadorRacao.DelimitadorDataRegistro = field.NewTime(tableName, "delimitador_data_registro")
	_delimitadorRacao.DelimitadorEmUso = field.NewInt32(tableName, "delimitador_em_uso")
	_delimitadorRacao.DelimitadorOficial = field.NewInt32(tableName, "delimitador_oficial")

	_delimitadorRacao.fillFieldMap()

	return _delimitadorRacao
}

type delimitadorRacao struct {
	delimitadorRacaoDo

	ALL                     field.Asterisk
	DelimitadorID           field.Int32
	DelimitadorDescricao    field.String
	DelimitadorDataRegistro field.Time
	DelimitadorEmUso        field.Int32
	DelimitadorOficial      field.Int32

	fieldMap map[string]field.Expr
}

func (d delimitadorRacao) Table(newTableName string) *delimitadorRacao {
	d.delimitadorRacaoDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d delimitadorRacao) As(alias string) *delimitadorRacao {
	d.delimitadorRacaoDo.DO = *(d.delimitadorRacaoDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *delimitadorRacao) updateTableName(table string) *delimitadorRacao {
	d.ALL = field.NewAsterisk(table)
	d.DelimitadorID = field.NewInt32(table, "delimitador_id")
	d.DelimitadorDescricao = field.NewString(table, "delimitador_descricao")
	d.DelimitadorDataRegistro = field.NewTime(table, "delimitador_data_registro")
	d.DelimitadorEmUso = field.NewInt32(table, "delimitador_em_uso")
	d.DelimitadorOficial = field.NewInt32(table, "delimitador_oficial")

	d.fillFieldMap()

	return d
}

func (d *delimitadorRacao) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *delimitadorRacao) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 5)
	d.fieldMap["delimitador_id"] = d.DelimitadorID
	d.fieldMap["delimitador_descricao"] = d.DelimitadorDescricao
	d.fieldMap["delimitador_data_registro"] = d.DelimitadorDataRegistro
	d.fieldMap["delimitador_em_uso"] = d.DelimitadorEmUso
	d.fieldMap["delimitador_oficial"] = d.DelimitadorOficial
}

func (d delimitadorRacao) clone(db *gorm.DB) delimitadorRacao {
	d.delimitadorRacaoDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d delimitadorRacao) replaceDB(db *gorm.DB) delimitadorRacao {
	d.delimitadorRacaoDo.ReplaceDB(db)
	return d
}

type delimitadorRacaoDo struct{ gen.DO }

type IDelimitadorRacaoDo interface {
	gen.SubQuery
	Debug() IDelimitadorRacaoDo
	WithContext(ctx context.Context) IDelimitadorRacaoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDelimitadorRacaoDo
	WriteDB() IDelimitadorRacaoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDelimitadorRacaoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDelimitadorRacaoDo
	Not(conds ...gen.Condition) IDelimitadorRacaoDo
	Or(conds ...gen.Condition) IDelimitadorRacaoDo
	Select(conds ...field.Expr) IDelimitadorRacaoDo
	Where(conds ...gen.Condition) IDelimitadorRacaoDo
	Order(conds ...field.Expr) IDelimitadorRacaoDo
	Distinct(cols ...field.Expr) IDelimitadorRacaoDo
	Omit(cols ...field.Expr) IDelimitadorRacaoDo
	Join(table schema.Tabler, on ...field.Expr) IDelimitadorRacaoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDelimitadorRacaoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDelimitadorRacaoDo
	Group(cols ...field.Expr) IDelimitadorRacaoDo
	Having(conds ...gen.Condition) IDelimitadorRacaoDo
	Limit(limit int) IDelimitadorRacaoDo
	Offset(offset int) IDelimitadorRacaoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDelimitadorRacaoDo
	Unscoped() IDelimitadorRacaoDo
	Create(values ...*models.DelimitadorRacao) error
	CreateInBatches(values []*models.DelimitadorRacao, batchSize int) error
	Save(values ...*models.DelimitadorRacao) error
	First() (*models.DelimitadorRacao, error)
	Take() (*models.DelimitadorRacao, error)
	Last() (*models.DelimitadorRacao, error)
	Find() ([]*models.DelimitadorRacao, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DelimitadorRacao, err error)
	FindInBatches(result *[]*models.DelimitadorRacao, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DelimitadorRacao) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDelimitadorRacaoDo
	Assign(attrs ...field.AssignExpr) IDelimitadorRacaoDo
	Joins(fields ...field.RelationField) IDelimitadorRacaoDo
	Preload(fields ...field.RelationField) IDelimitadorRacaoDo
	FirstOrInit() (*models.DelimitadorRacao, error)
	FirstOrCreate() (*models.DelimitadorRacao, error)
	FindByPage(offset int, limit int) (result []*models.DelimitadorRacao, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDelimitadorRacaoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d delimitadorRacaoDo) Debug() IDelimitadorRacaoDo {
	return d.withDO(d.DO.Debug())
}

func (d delimitadorRacaoDo) WithContext(ctx context.Context) IDelimitadorRacaoDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d delimitadorRacaoDo) ReadDB() IDelimitadorRacaoDo {
	return d.Clauses(dbresolver.Read)
}

func (d delimitadorRacaoDo) WriteDB() IDelimitadorRacaoDo {
	return d.Clauses(dbresolver.Write)
}

func (d delimitadorRacaoDo) Session(config *gorm.Session) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Session(config))
}

func (d delimitadorRacaoDo) Clauses(conds ...clause.Expression) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d delimitadorRacaoDo) Returning(value interface{}, columns ...string) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d delimitadorRacaoDo) Not(conds ...gen.Condition) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d delimitadorRacaoDo) Or(conds ...gen.Condition) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d delimitadorRacaoDo) Select(conds ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d delimitadorRacaoDo) Where(conds ...gen.Condition) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d delimitadorRacaoDo) Order(conds ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d delimitadorRacaoDo) Distinct(cols ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d delimitadorRacaoDo) Omit(cols ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d delimitadorRacaoDo) Join(table schema.Tabler, on ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d delimitadorRacaoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d delimitadorRacaoDo) RightJoin(table schema.Tabler, on ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d delimitadorRacaoDo) Group(cols ...field.Expr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d delimitadorRacaoDo) Having(conds ...gen.Condition) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d delimitadorRacaoDo) Limit(limit int) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d delimitadorRacaoDo) Offset(offset int) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d delimitadorRacaoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d delimitadorRacaoDo) Unscoped() IDelimitadorRacaoDo {
	return d.withDO(d.DO.Unscoped())
}

func (d delimitadorRacaoDo) Create(values ...*models.DelimitadorRacao) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d delimitadorRacaoDo) CreateInBatches(values []*models.DelimitadorRacao, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d delimitadorRacaoDo) Save(values ...*models.DelimitadorRacao) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d delimitadorRacaoDo) First() (*models.DelimitadorRacao, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DelimitadorRacao), nil
	}
}

func (d delimitadorRacaoDo) Take() (*models.DelimitadorRacao, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DelimitadorRacao), nil
	}
}

func (d delimitadorRacaoDo) Last() (*models.DelimitadorRacao, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DelimitadorRacao), nil
	}
}

func (d delimitadorRacaoDo) Find() ([]*models.DelimitadorRacao, error) {
	result, err := d.DO.Find()
	return result.([]*models.DelimitadorRacao), err
}

func (d delimitadorRacaoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DelimitadorRacao, err error) {
	buf := make([]*models.DelimitadorRacao, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d delimitadorRacaoDo) FindInBatches(result *[]*models.DelimitadorRacao, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d delimitadorRacaoDo) Attrs(attrs ...field.AssignExpr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d delimitadorRacaoDo) Assign(attrs ...field.AssignExpr) IDelimitadorRacaoDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d delimitadorRacaoDo) Joins(fields ...field.RelationField) IDelimitadorRacaoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d delimitadorRacaoDo) Preload(fields ...field.RelationField) IDelimitadorRacaoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d delimitadorRacaoDo) FirstOrInit() (*models.DelimitadorRacao, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DelimitadorRacao), nil
	}
}

func (d delimitadorRacaoDo) FirstOrCreate() (*models.DelimitadorRacao, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DelimitadorRacao), nil
	}
}

func (d delimitadorRacaoDo) FindByPage(offset int, limit int) (result []*models.DelimitadorRacao, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d delimitadorRacaoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d delimitadorRacaoDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d delimitadorRacaoDo) Delete(models ...*models.DelimitadorRacao) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *delimitadorRacaoDo) withDO(do gen.Dao) *delimitadorRacaoDo {
	d.DO = *do.(*gen.DO)
	return d
}
