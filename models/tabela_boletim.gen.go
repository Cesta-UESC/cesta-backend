package models

const TableNameBoletim = "tabela_boletim"

// Boletim mapped from table <tabela_boletim>
type Boletim struct {
	BoletimID   int32  `gorm:"column:boletim_id;primaryKey" json:"boletim_id"`
	BoletimNome string `gorm:"column:boletim_nome;not null" json:"boletim_nome"`
}

// TableName Boletim's table name
func (*Boletim) TableName() string {
	return TableNameBoletim
}
