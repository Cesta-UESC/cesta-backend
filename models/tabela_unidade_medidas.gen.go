package models

const TableNameUnidadeMedidas = "tabela_unidade_medidas"

// UnidadeMedidas mapped from table <tabela_unidade_medidas>
type UnidadeMedidas struct {
	MedidaID        int32  `gorm:"column:medida_id;primaryKey;autoIncrement:true" json:"medida_id"`
	MedidaDescricao string `gorm:"column:medida_descricao;not null" json:"medida_descricao"`
	MedidaSimbolo   string `gorm:"column:medida_simbolo;not null" json:"medida_simbolo"`
}

// TableName UnidadeMedidas's table name
func (*UnidadeMedidas) TableName() string {
	return TableNameUnidadeMedidas
}
