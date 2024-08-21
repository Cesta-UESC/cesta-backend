package models

const TableNameRacaoMinima = "tabela_racao_minima"

// RacaoMinima mapped from table <tabela_racao_minima>
type RacaoMinima struct {
	DelimitadorID            int32   `gorm:"column:delimitador_id;primaryKey" json:"delimitador_id"`
	ProdutoID                int32   `gorm:"column:produto_id;primaryKey" json:"produto_id"`
	MedidaID                 int32   `gorm:"column:medida_id;primaryKey" json:"medida_id"`
	RacaoMinimaQuantidade    float64 `gorm:"column:racao_minima_quantidade;not null" json:"racao_minima_quantidade"`
	RacaoMinimaTransformador float64 `gorm:"column:racao_minima_transformador;not null" json:"racao_minima_transformador"`
	RacaoMinimaMedida        int32   `gorm:"column:racao_minima_medida;not null" json:"racao_minima_medida"`
}

// TableName RacaoMinima's table name
func (*RacaoMinima) TableName() string {
	return TableNameRacaoMinima
}
