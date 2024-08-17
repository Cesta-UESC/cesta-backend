package model

const TableNameProdutos = "tabela_produtos"

// Produtos mapped from table <tabela_produtos>
type Produtos struct {
	ProdutoID               int32  `gorm:"column:produto_id;primaryKey;autoIncrement:true" json:"produto_id"`
	ProdutoNome             string `gorm:"column:produto_nome;not null" json:"produto_nome"`
	ProdutoCesta            int32  `gorm:"column:produto_cesta;not null" json:"produto_cesta"`
	ProdutoNomeVisualizacao string `gorm:"column:produto_nome_visualizacao;not null" json:"produto_nome_visualizacao"`
	ProdutoTipo             int32  `gorm:"column:produto_tipo;not null" json:"produto_tipo"`
}

// TableName Produtos's table name
func (*Produtos) TableName() string {
	return TableNameProdutos
}
