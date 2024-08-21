package models

const TableNameAuxiliarPrecos = "tabela_auxiliar_precos"

// AuxiliarPrecos mapped from table <tabela_auxiliar_precos>
type AuxiliarPrecos struct {
	PrecosID     int32   `gorm:"column:precos_id;not null" json:"precos_id"`
	PrecoProduto float64 `gorm:"column:preco_produto;not null" json:"preco_produto"`
}

// TableName AuxiliarPrecos's table name
func (*AuxiliarPrecos) TableName() string {
	return TableNameAuxiliarPrecos
}
