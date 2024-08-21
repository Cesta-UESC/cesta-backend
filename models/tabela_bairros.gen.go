package models

const TableNameBairros = "tabela_bairros"

// Bairros mapped from table <tabela_bairros>
type Bairros struct {
	BairroID   int32  `gorm:"column:bairro_id;primaryKey;autoIncrement:true" json:"bairro_id"`
	BairroNome string `gorm:"column:bairro_nome;not null" json:"bairro_nome"`
	CidadeID   int32  `gorm:"column:cidade_id;not null" json:"cidade_id"`
}

// TableName Bairros's table name
func (*Bairros) TableName() string {
	return TableNameBairros
}
