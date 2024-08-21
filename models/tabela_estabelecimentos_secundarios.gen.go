package models

const TableNameEstabelecimentosSecundarios = "tabela_estabelecimentos_secundarios"

// EstabelecimentosSecundarios mapped from table <tabela_estabelecimentos_secundarios>
type EstabelecimentosSecundarios struct {
	EstabelecimentoSecID   int32  `gorm:"column:estabelecimento_sec_id;primaryKey;autoIncrement:true" json:"estabelecimento_sec_id"`
	EstabelecimentoSecNome string `gorm:"column:estabelecimento_sec_nome;not null" json:"estabelecimento_sec_nome"`
}

// TableName EstabelecimentosSecundarios's table name
func (*EstabelecimentosSecundarios) TableName() string {
	return TableNameEstabelecimentosSecundarios
}
