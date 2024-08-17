package model

const TableNameEstabelecimentoHasSecundario = "tabela_estabelecimento_has_secundario"

// EstabelecimentoHasSecundario mapped from table <tabela_estabelecimento_has_secundario>
type EstabelecimentoHasSecundario struct {
	EstHasSecID          int32 `gorm:"column:est_has_sec_id;primaryKey;autoIncrement:true" json:"est_has_sec_id"`
	EstabelecimentoID    int32 `gorm:"column:estabelecimento_id;not null" json:"estabelecimento_id"`
	EstabelecimentoSecID int32 `gorm:"column:estabelecimento_sec_id;not null" json:"estabelecimento_sec_id"`
}

// TableName EstabelecimentoHasSecundario's table name
func (*EstabelecimentoHasSecundario) TableName() string {
	return TableNameEstabelecimentoHasSecundario
}
