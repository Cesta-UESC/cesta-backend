package models

type User struct {
	ID            int    `gorm:"primaryKey;autoIncrement;column:usuario_id"`
	Name          string `gorm:"column:usuario_nome;size:64"`
	CleanPassword string `gorm:"column:usuario_senha;size:128"` // DEPRECATED
	Email         string `gorm:"column:usuario_email;size:255"`
}

func (User) TableName() string {
	return "tabela_usuarios"
}
