package model

import (
	"time"
)

const TableNameAuxiliarCronograma = "tabela_auxiliar_cronograma"

// AuxiliarCronograma mapped from table <tabela_auxiliar_cronograma>
type AuxiliarCronograma struct {
	CronogramaID int32     `gorm:"column:cronograma_id;not null" json:"cronograma_id"`
	MesID        int32     `gorm:"column:mes_id;not null" json:"mes_id"`
	InicioColeta time.Time `gorm:"column:inicio_coleta;not null" json:"inicio_coleta"`
	FimColeta    time.Time `gorm:"column:fim_coleta;not null" json:"fim_coleta"`
}

// TableName AuxiliarCronograma's table name
func (*AuxiliarCronograma) TableName() string {
	return TableNameAuxiliarCronograma
}
