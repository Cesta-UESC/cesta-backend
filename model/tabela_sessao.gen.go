package model

const TableNameSessao = "tabela_sessao"

// Sessao mapped from table <tabela_sessao>
type Sessao struct {
	SessaoID      string `gorm:"column:sessao_id;primaryKey" json:"sessao_id"`
	SessaoUsuario int32  `gorm:"column:sessao_usuario;not null" json:"sessao_usuario"`
	SessaoIP      string `gorm:"column:sessao_ip;not null" json:"sessao_ip"`
	SessaoTempo   int64  `gorm:"column:sessao_tempo;not null" json:"sessao_tempo"`
}

// TableName Sessao's table name
func (*Sessao) TableName() string {
	return TableNameSessao
}
