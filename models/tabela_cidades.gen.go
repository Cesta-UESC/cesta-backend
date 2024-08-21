package models

import (
	"time"
)

const TableNameCidades = "tabela_cidades"

// Cidades mapped from table <tabela_cidades>
type Cidades struct {
	CidadeID   int32     `gorm:"column:cidade_id;primaryKey;autoIncrement:true" json:"cidade_id"`
	CidadeNome string    `gorm:"column:cidade_nome;not null" json:"cidade_nome"`
	CidadeData time.Time `gorm:"column:cidade_data;not null" json:"cidade_data"`
}

// TableName Cidades's table name
func (*Cidades) TableName() string {
	return TableNameCidades
}
