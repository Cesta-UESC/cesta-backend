package model

import (
	"time"
)

const TableNamePesquisas = "tabela_pesquisas"

// Pesquisas mapped from table <tabela_pesquisas>
type Pesquisas struct {
	PesquisaID        int32     `gorm:"column:pesquisa_id;primaryKey;autoIncrement:true" json:"pesquisa_id"`
	DelimitadorID     int32     `gorm:"column:delimitador_id" json:"delimitador_id"`
	SalarioID         int32     `gorm:"column:salario_id" json:"salario_id"`
	PesquisaFechada   int32     `gorm:"column:pesquisa_fechada;not null" json:"pesquisa_fechada"`
	PesquisaData      time.Time `gorm:"column:pesquisa_data;not null" json:"pesquisa_data"`
	PesquisaDetalhada int32     `gorm:"column:pesquisa_detalhada" json:"pesquisa_detalhada"`
}

// TableName Pesquisas's table name
func (*Pesquisas) TableName() string {
	return TableNamePesquisas
}
