package models

import (
	"time"
)

const TableNameDelimitadorRacao = "tabela_delimitador_racao"

// DelimitadorRacao mapped from table <tabela_delimitador_racao>
type DelimitadorRacao struct {
	DelimitadorID           int32     `gorm:"column:delimitador_id;primaryKey;autoIncrement:true" json:"delimitador_id"`
	DelimitadorDescricao    string    `gorm:"column:delimitador_descricao;not null" json:"delimitador_descricao"`
	DelimitadorDataRegistro time.Time `gorm:"column:delimitador_data_registro;not null" json:"delimitador_data_registro"`
	DelimitadorEmUso        int32     `gorm:"column:delimitador_em_uso;not null" json:"delimitador_em_uso"`
	DelimitadorOficial      int32     `gorm:"column:delimitador_oficial;not null" json:"delimitador_oficial"`
}

// TableName DelimitadorRacao's table name
func (*DelimitadorRacao) TableName() string {
	return TableNameDelimitadorRacao
}
