package model

const TableNameTiposProdutos = "tabela_tipos_produtos"

// TiposProdutos mapped from table <tabela_tipos_produtos>
type TiposProdutos struct {
	TipoID   int32  `gorm:"column:tipo_id;primaryKey;autoIncrement:true" json:"tipo_id"`
	TipoNome string `gorm:"column:tipo_nome;not null" json:"tipo_nome"`
}

// TableName TiposProdutos's table name
func (*TiposProdutos) TableName() string {
	return TableNameTiposProdutos
}
