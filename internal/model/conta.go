package model

import "time"

type Conta struct {
	ID        string    `json:"id"`
	UsuarioID string    `json:"usuario_id"`
	Agencia   string    `json:"agencia"`
	Numero    string    `json:"numero"`
	Status    string    `json:"status"` // PENDING, ACTIVE, BLOCKED
	CriadoEm  time.Time `json:"criado_em"`
}
