package models

const TableNamePesquisaResultadosProdutos = "tabela_pesquisa_resultados_produtos"

// PesquisaResultadosProdutos mapped from table <tabela_pesquisa_resultados_produtos>
type PesquisaResultadosProdutos struct {
	ProdutoID                int32   `gorm:"column:produto_id;primaryKey" json:"produto_id"`
	CidadeID                 int32   `gorm:"column:cidade_id;primaryKey" json:"cidade_id"`
	PesquisaID               int32   `gorm:"column:pesquisa_id;primaryKey" json:"pesquisa_id"`
	ProdutoVariacaoMensal    float64 `gorm:"column:produto_variacao_mensal" json:"produto_variacao_mensal"`
	ProdutoVariacaoSemestral float64 `gorm:"column:produto_variacao_semestral" json:"produto_variacao_semestral"`
	ProdutoVariacaoAnual     float64 `gorm:"column:produto_variacao_anual" json:"produto_variacao_anual"`
	ProdutoPrecoMedio        float64 `gorm:"column:produto_preco_medio" json:"produto_preco_medio"`
	ProdutoPrecoTotal        float64 `gorm:"column:produto_preco_total" json:"produto_preco_total"`
	ProdutoTempoTrabalho     float64 `gorm:"column:produto_tempo_trabalho" json:"produto_tempo_trabalho"`
}

// TableName PesquisaResultadosProdutos's table name
func (*PesquisaResultadosProdutos) TableName() string {
	return TableNamePesquisaResultadosProdutos
}
