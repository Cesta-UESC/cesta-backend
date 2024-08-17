package model

import (
	"time"
)

const TableNameEstabelecimentos = "tabela_estabelecimentos"

// Estabelecimentos mapped from table <tabela_estabelecimentos>
type Estabelecimentos struct {
	EstabelecimentoID          int32     `gorm:"column:estabelecimento_id;primaryKey;autoIncrement:true" json:"estabelecimento_id"`
	BairroID                   int32     `gorm:"column:bairro_id;not null" json:"bairro_id"`
	EstabelecimentoNome        string    `gorm:"column:estabelecimento_nome;not null" json:"estabelecimento_nome"`
	EstabelecimentoData        time.Time `gorm:"column:estabelecimento_data;not null" json:"estabelecimento_data"`
	EstabelecimentoEndereco    string    `gorm:"column:estabelecimento_endereco" json:"estabelecimento_endereco"`
	EstabelecimentoContato     string    `gorm:"column:estabelecimento_contato" json:"estabelecimento_contato"`
	EstabelecimentoTelefone    string    `gorm:"column:estabelecimento_telefone" json:"estabelecimento_telefone"`
	EstabelecimentoReferencial string    `gorm:"column:estabelecimento_referencial" json:"estabelecimento_referencial"`
	EstabelecimentoAtivo       bool      `gorm:"column:estabelecimento_ativo;default:1" json:"estabelecimento_ativo"`
}

// TableName Estabelecimentos's table name
func (*Estabelecimentos) TableName() string {
	return TableNameEstabelecimentos
}
