package model

import "time"

type PessoaFisica struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Documento string    `json:"documento"`
	CriadoEm  time.Time `json:"criado_em"`
}
