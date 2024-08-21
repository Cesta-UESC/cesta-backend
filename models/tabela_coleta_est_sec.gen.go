package models

const TableNameColetaEstSec = "tabela_coleta_est_sec"

// ColetaEstSec mapped from table <tabela_coleta_est_sec>
type ColetaEstSec struct {
	ColetaEstSecID int32 `gorm:"column:coleta_est_sec_id;primaryKey;autoIncrement:true" json:"coleta_est_sec_id"`
	ColetaID       int32 `gorm:"column:coleta_id;not null" json:"coleta_id"`
	EstHasSecID    int32 `gorm:"column:est_has_sec_id;not null" json:"est_has_sec_id"`
	ProdutoID      int32 `gorm:"column:produto_id;not null" json:"produto_id"`
}

// TableName ColetaEstSec's table name
func (*ColetaEstSec) TableName() string {
	return TableNameColetaEstSec
}
