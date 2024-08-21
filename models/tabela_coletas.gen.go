package models

import (
	"time"
)

const TableNameColetas = "tabela_coletas"

// Coletas mapped from table <tabela_coletas>
type Coletas struct {
	ColetaID          int32     `gorm:"column:coleta_id;primaryKey;autoIncrement:true" json:"coleta_id"`
	EstabelecimentoID int32     `gorm:"column:estabelecimento_id;not null" json:"estabelecimento_id"`
	PesquisaID        int32     `gorm:"column:pesquisa_id;not null" json:"pesquisa_id"`
	ColetaData        time.Time `gorm:"column:coleta_data;not null" json:"coleta_data"`
	ColetaPrecoCesta  float64   `gorm:"column:coleta_preco_cesta" json:"coleta_preco_cesta"`
	ColetaFechada     int32     `gorm:"column:coleta_fechada" json:"coleta_fechada"`
}

// TableName Coletas's table name
func (*Coletas) TableName() string {
	return TableNameColetas
}
