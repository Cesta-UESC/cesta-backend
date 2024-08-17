package model

const TableNameMes = "tabela_mes"

// Mes mapped from table <tabela_mes>
type Mes struct {
	MesID   int32  `gorm:"column:mes_id;primaryKey" json:"mes_id"`
	MesNome string `gorm:"column:mes_nome;not null" json:"mes_nome"`
}

// TableName Mes's table name
func (*Mes) TableName() string {
	return TableNameMes
}
