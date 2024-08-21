package models

const TableNamePesquisasCidades = "tabela_pesquisas_cidades"

// PesquisasCidades mapped from table <tabela_pesquisas_cidades>
type PesquisasCidades struct {
	PesquisaID        int32   `gorm:"column:pesquisa_id;primaryKey" json:"pesquisa_id"`
	CidadeID          int32   `gorm:"column:cidade_id;primaryKey" json:"cidade_id"`
	VariacaoMensal    float64 `gorm:"column:variacao_mensal" json:"variacao_mensal"`
	VariacaoSemestral float64 `gorm:"column:variacao_semestral" json:"variacao_semestral"`
	VariacaoAnual     float64 `gorm:"column:variacao_anual" json:"variacao_anual"`
	GastoMensalCesta  float64 `gorm:"column:gasto_mensal_cesta" json:"gasto_mensal_cesta"`
	TempoTrabalho     float64 `gorm:"column:tempo_trabalho" json:"tempo_trabalho"`
}

// TableName PesquisasCidades's table name
func (*PesquisasCidades) TableName() string {
	return TableNamePesquisasCidades
}
