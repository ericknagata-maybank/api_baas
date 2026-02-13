package model

import "time"

type Representante struct {
	ID       string    `json:"id"`
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	Telefone string    `json:"telefone"`
	CriadoEm time.Time `json:"criado_em"`
}
