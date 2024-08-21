package api

import "time"

type LoginData struct {
	Username string `validate:"required,max=320"`
	Password string `validate:"required,max=100"`
}

type LoginResponse struct {
	Token      string    `json:"token"`
	Type       string    `json:"type"`
	Expiration time.Time `json:"expiration"`
}
