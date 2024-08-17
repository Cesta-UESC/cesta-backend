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

func newPrecos(db *gorm.DB, opts ...gen.DOOption) precos {
	_precos := precos{}

	_precos.precosDo.UseDB(db, opts...)
	_precos.precosDo.UseModel(&model.Precos{})

	tableName := _precos.precosDo.TableName()
	_precos.ALL = field.NewAsterisk(tableName)
	_precos.PrecosID = field.NewInt32(tableName, "precos_id")
	_precos.MedidaID = field.NewInt32(tableName, "medida_id")
	_precos.ProdutoID = field.NewInt32(tableName, "produto_id")
	_precos.ColetaID = field.NewInt32(tableName, "coleta_id")
	_precos.PrecosMediaObservado = field.NewFloat64(tableName, "precos_media_observado")
	_precos.PrecosMedia = field.NewFloat64(tableName, "precos_media")
	_precos.PrecosTotal = field.NewFloat64(tableName, "precos_total")

	_precos.fillFieldMap()

	return _precos
}

type precos struct {
	precosDo

	ALL                  field.Asterisk
	PrecosID             field.Int32
	MedidaID             field.Int32
	ProdutoID            field.Int32
	ColetaID             field.Int32
	PrecosMediaObservado field.Float64
	PrecosMedia          field.Float64
	PrecosTotal          field.Float64

	fieldMap map[string]field.Expr
}

func (p precos) Table(newTableName string) *precos {
	p.precosDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p precos) As(alias string) *precos {
	p.precosDo.DO = *(p.precosDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *precos) updateTableName(table string) *precos {
	p.ALL = field.NewAsterisk(table)
	p.PrecosID = field.NewInt32(table, "precos_id")
	p.MedidaID = field.NewInt32(table, "medida_id")
	p.ProdutoID = field.NewInt32(table, "produto_id")
	p.ColetaID = field.NewInt32(table, "coleta_id")
	p.PrecosMediaObservado = field.NewFloat64(table, "precos_media_observado")
	p.PrecosMedia = field.NewFloat64(table, "precos_media")
	p.PrecosTotal = field.NewFloat64(table, "precos_total")

	p.fillFieldMap()

	return p
}

func (p *precos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *precos) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["precos_id"] = p.PrecosID
	p.fieldMap["medida_id"] = p.MedidaID
	p.fieldMap["produto_id"] = p.ProdutoID
	p.fieldMap["coleta_id"] = p.ColetaID
	p.fieldMap["precos_media_observado"] = p.PrecosMediaObservado
	p.fieldMap["precos_media"] = p.PrecosMedia
	p.fieldMap["precos_total"] = p.PrecosTotal
}

func (p precos) clone(db *gorm.DB) precos {
	p.precosDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p precos) replaceDB(db *gorm.DB) precos {
	p.precosDo.ReplaceDB(db)
	return p
}

type precosDo struct{ gen.DO }

type IPrecosDo interface {
	gen.SubQuery
	Debug() IPrecosDo
	WithContext(ctx context.Context) IPrecosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPrecosDo
	WriteDB() IPrecosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPrecosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPrecosDo
	Not(conds ...gen.Condition) IPrecosDo
	Or(conds ...gen.Condition) IPrecosDo
	Select(conds ...field.Expr) IPrecosDo
	Where(conds ...gen.Condition) IPrecosDo
	Order(conds ...field.Expr) IPrecosDo
	Distinct(cols ...field.Expr) IPrecosDo
	Omit(cols ...field.Expr) IPrecosDo
	Join(table schema.Tabler, on ...field.Expr) IPrecosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPrecosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPrecosDo
	Group(cols ...field.Expr) IPrecosDo
	Having(conds ...gen.Condition) IPrecosDo
	Limit(limit int) IPrecosDo
	Offset(offset int) IPrecosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPrecosDo
	Unscoped() IPrecosDo
	Create(values ...*model.Precos) error
	CreateInBatches(values []*model.Precos, batchSize int) error
	Save(values ...*model.Precos) error
	First() (*model.Precos, error)
	Take() (*model.Precos, error)
	Last() (*model.Precos, error)
	Find() ([]*model.Precos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Precos, err error)
	FindInBatches(result *[]*model.Precos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Precos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPrecosDo
	Assign(attrs ...field.AssignExpr) IPrecosDo
	Joins(fields ...field.RelationField) IPrecosDo
	Preload(fields ...field.RelationField) IPrecosDo
	FirstOrInit() (*model.Precos, error)
	FirstOrCreate() (*model.Precos, error)
	FindByPage(offset int, limit int) (result []*model.Precos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPrecosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p precosDo) Debug() IPrecosDo {
	return p.withDO(p.DO.Debug())
}

func (p precosDo) WithContext(ctx context.Context) IPrecosDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p precosDo) ReadDB() IPrecosDo {
	return p.Clauses(dbresolver.Read)
}

func (p precosDo) WriteDB() IPrecosDo {
	return p.Clauses(dbresolver.Write)
}

func (p precosDo) Session(config *gorm.Session) IPrecosDo {
	return p.withDO(p.DO.Session(config))
}

func (p precosDo) Clauses(conds ...clause.Expression) IPrecosDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p precosDo) Returning(value interface{}, columns ...string) IPrecosDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p precosDo) Not(conds ...gen.Condition) IPrecosDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p precosDo) Or(conds ...gen.Condition) IPrecosDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p precosDo) Select(conds ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p precosDo) Where(conds ...gen.Condition) IPrecosDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p precosDo) Order(conds ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p precosDo) Distinct(cols ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p precosDo) Omit(cols ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p precosDo) Join(table schema.Tabler, on ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p precosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p precosDo) RightJoin(table schema.Tabler, on ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p precosDo) Group(cols ...field.Expr) IPrecosDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p precosDo) Having(conds ...gen.Condition) IPrecosDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p precosDo) Limit(limit int) IPrecosDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p precosDo) Offset(offset int) IPrecosDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p precosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPrecosDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p precosDo) Unscoped() IPrecosDo {
	return p.withDO(p.DO.Unscoped())
}

func (p precosDo) Create(values ...*model.Precos) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p precosDo) CreateInBatches(values []*model.Precos, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p precosDo) Save(values ...*model.Precos) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p precosDo) First() (*model.Precos, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Precos), nil
	}
}

func (p precosDo) Take() (*model.Precos, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Precos), nil
	}
}

func (p precosDo) Last() (*model.Precos, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Precos), nil
	}
}

func (p precosDo) Find() ([]*model.Precos, error) {
	result, err := p.DO.Find()
	return result.([]*model.Precos), err
}

func (p precosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Precos, err error) {
	buf := make([]*model.Precos, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p precosDo) FindInBatches(result *[]*model.Precos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p precosDo) Attrs(attrs ...field.AssignExpr) IPrecosDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p precosDo) Assign(attrs ...field.AssignExpr) IPrecosDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p precosDo) Joins(fields ...field.RelationField) IPrecosDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p precosDo) Preload(fields ...field.RelationField) IPrecosDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p precosDo) FirstOrInit() (*model.Precos, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Precos), nil
	}
}

func (p precosDo) FirstOrCreate() (*model.Precos, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Precos), nil
	}
}

func (p precosDo) FindByPage(offset int, limit int) (result []*model.Precos, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p precosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p precosDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p precosDo) Delete(models ...*model.Precos) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *precosDo) withDO(do gen.Dao) *precosDo {
	p.DO = *do.(*gen.DO)
	return p
}
