package models

const TableNameUsuarios = "tabela_usuarios"

// Usuarios mapped from table <tabela_usuarios>
type Usuarios struct {
	UsuarioID    int32  `gorm:"column:usuario_id;primaryKey;autoIncrement:true" json:"usuario_id"`
	UsuarioNome  string `gorm:"column:usuario_nome;not null" json:"usuario_nome"`
	UsuarioSenha string `gorm:"column:usuario_senha;not null" json:"usuario_senha"`
	UsuarioEmail string `gorm:"column:usuario_email;not null" json:"usuario_email"`
}

// TableName Usuarios's table name
func (*Usuarios) TableName() string {
	return TableNameUsuarios
}
