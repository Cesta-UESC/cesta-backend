package model

const TableNamePrecos = "tabela_precos"

// Precos mapped from table <tabela_precos>
type Precos struct {
	PrecosID             int32   `gorm:"column:precos_id;primaryKey;autoIncrement:true" json:"precos_id"`
	MedidaID             int32   `gorm:"column:medida_id;not null" json:"medida_id"`
	ProdutoID            int32   `gorm:"column:produto_id;not null" json:"produto_id"`
	ColetaID             int32   `gorm:"column:coleta_id;not null" json:"coleta_id"`
	PrecosMediaObservado float64 `gorm:"column:precos_media_observado" json:"precos_media_observado"`
	PrecosMedia          float64 `gorm:"column:precos_media" json:"precos_media"`
	PrecosTotal          float64 `gorm:"column:precos_total" json:"precos_total"`
}

// TableName Precos's table name
func (*Precos) TableName() string {
	return TableNamePrecos
}
