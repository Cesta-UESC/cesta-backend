package models

import (
	"time"
)

const TableNameSalarios = "tabela_salarios"

// Salarios mapped from table <tabela_salarios>
type Salarios struct {
	SalarioID           int32     `gorm:"column:salario_id;primaryKey;autoIncrement:true" json:"salario_id"`
	SalarioValorBruto   float64   `gorm:"column:salario_valor_bruto;not null" json:"salario_valor_bruto"`
	SalarioValorLiquido float64   `gorm:"column:salario_valor_liquido;not null" json:"salario_valor_liquido"`
	SalarioEmUso        int32     `gorm:"column:salario_em_uso;not null" json:"salario_em_uso"`
	SalarioDataRegistro time.Time `gorm:"column:salario_data_registro;not null" json:"salario_data_registro"`
	SalarioNome         string    `gorm:"column:salario_nome;not null" json:"salario_nome"`
	SalarioSimbolo      string    `gorm:"column:salario_simbolo;not null" json:"salario_simbolo"`
}

// TableName Salarios's table name
func (*Salarios) TableName() string {
	return TableNameSalarios
}
