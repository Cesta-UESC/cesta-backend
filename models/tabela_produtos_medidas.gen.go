package models

const TableNameProdutosMedidas = "tabela_produtos_medidas"

// ProdutosMedidas mapped from table <tabela_produtos_medidas>
type ProdutosMedidas struct {
	ProdutoID        int32   `gorm:"column:produto_id;primaryKey" json:"produto_id"`
	MedidaID         int32   `gorm:"column:medida_id;primaryKey" json:"medida_id"`
	MedidaPesquisada float64 `gorm:"column:medida_pesquisada;not null" json:"medida_pesquisada"`
}

// TableName ProdutosMedidas's table name
func (*ProdutosMedidas) TableName() string {
	return TableNameProdutosMedidas
}
