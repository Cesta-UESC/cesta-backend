package models

const TableNameEquipe = "tabela_equipe"

// Equipe mapped from table <tabela_equipe>
type Equipe struct {
	ID              int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NomeCompleto    string `gorm:"column:nome_completo;not null" json:"nome_completo"`
	Email           string `gorm:"column:email;not null" json:"email"`
	FuncaoID        int32  `gorm:"column:funcao_id;not null" json:"funcao_id"`
	MostrarHome     bool   `gorm:"column:mostrar_home;default:1" json:"mostrar_home"`
	MostrarContatos bool   `gorm:"column:mostrar_contatos;default:1" json:"mostrar_contatos"`
}

// TableName Equipe's table name
func (*Equipe) TableName() string {
	return TableNameEquipe
}
