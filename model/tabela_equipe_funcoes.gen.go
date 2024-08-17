package model

const TableNameEquipeFuncoes = "tabela_equipe_funcoes"

// EquipeFuncoes mapped from table <tabela_equipe_funcoes>
type EquipeFuncoes struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Funcao string `gorm:"column:funcao" json:"funcao"`
}

// TableName EquipeFuncoes's table name
func (*EquipeFuncoes) TableName() string {
	return TableNameEquipeFuncoes
}
