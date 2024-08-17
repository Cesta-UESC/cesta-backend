package model

const TableNameCronogramaColetas = "tabela_cronograma_coletas"

// CronogramaColetas mapped from table <tabela_cronograma_coletas>
type CronogramaColetas struct {
	CronogramaID int32 `gorm:"column:cronograma_id;primaryKey;autoIncrement:true" json:"cronograma_id"`
	Ano          int32 `gorm:"column:ano;not null" json:"ano"`
}

// TableName CronogramaColetas's table name
func (*CronogramaColetas) TableName() string {
	return TableNameCronogramaColetas
}
