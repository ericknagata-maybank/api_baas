package repository

import (
	"database/sql"

	"github.com/ericknagata-maybank/api_baas/internal/model"
)

type DormentePfRepository struct {
	db *sql.DB
}

func NewDormentePfRepository(db *sql.DB) *DormentePfRepository {
	return &DormentePfRepository{db: db}
}

func (r *DormentePfRepository) Create(pf *model.DormentePf) error {
	query := `
		INSERT INTO dormente_pf (
			cpf, documento, data_emissao, orgao_emissor, estado_emissor,
			tipo_passaporte, pais_emissor_passaporte, nacionalidade_passaporte, naturalidade_passaporte, validade_passaporte,
			nome, idoc, sexo, data_nascimento, estado_civil, politico, renda, nome_mae,
			cep, numero, endereco, estado, cidade, bairro, complemento,
			email, telefone, senha, confirmation_url, data_hora_cadastro
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		pf.CPF, pf.Documento, pf.DataEmissao, pf.OrgaoEmissor, pf.EstadoEmissor,
		pf.TipoPassaporte, pf.PaisEmissorPassaporte, pf.Nacionalidade, pf.Naturalidade, pf.ValidadePassaporte,
		pf.Nome, pf.Idoc, pf.Sexo, pf.DataNascimento, pf.EstadoCivil, pf.Politico, pf.Renda, pf.NomeMae,
		pf.Cep, pf.Numero, pf.Endereco, pf.Estado, pf.Cidade, pf.Bairro, pf.Complemento,
		pf.Email, pf.Telefone, pf.Senha, pf.ConfirmationURL, pf.DataHoraCadastro,
	)

	return err
}
