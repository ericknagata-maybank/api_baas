package repository

import (
	"database/sql"

	"github.com/ericknagata-maybank/api_baas/internal/model"
)

type DormentePjRepository struct {
	db *sql.DB
}

func NewDormentePjRepository(db *sql.DB) *DormentePjRepository {
	return &DormentePjRepository{db: db}
}

func (r *DormentePjRepository) Create(pj *model.DormentePJ) error {
	query := `
		INSERT INTO dormente_pj (
			razao_social, nome_fantasia, cnpj, agencia, cnae, faturamento, contribuicao,
			inscricao_estadual, telefone, email, data_abertura, cep, endereco, numero,
			complemento, bairro, cidade, estado, representante_nome, representante_cpf,
			representante_data_nascimento, representante_telefone, representante_nome_mae,
			representante_sexo, representante_email, representante_estado_civil,
			representante_cep, representante_endereco, representante_numero,
			representante_complemento, representante_bairro, representante_cidade,
			representante_estado, representante_documento, representante_data_emissao,
			representante_orgao_emissor, representante_estado_emissor, datahora,
			usado, datahora_usado, arquivo, aparelho, permite_criacao_automatica,
			senha, imagecnpj, imagecomprovante, imagecontrato, imagecomprovante_endereco,
			representante_imageself, representante_imagedoc, representante_imagedoc_verso,
			representante_imagecomprovante_endereco, representante, representante_politico,
			representante_idoc, representante_procuracao
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		pj.Razao_social, pj.Nome_fantasia, pj.Cnpj, pj.Agencia, pj.Cnae, pj.Faturamento, pj.Contribuicao,
		pj.Inscricao_estadual, pj.Telefone, pj.Email, pj.Data_abertura, pj.Cep, pj.Endereco, pj.Numero,
		pj.Complemento, pj.Bairro, pj.Cidade, pj.Estado, pj.Representante_nome, pj.Representante_cpf,
		pj.Representante_data_nascimento, pj.Representante_telefone, pj.Representante_nome_mae,
		pj.Representante_sexo, pj.Representante_email, pj.Representante_estado_civil,
		pj.Representante_cep, pj.Representante_endereco, pj.Representante_numero,
		pj.Representante_complemento, pj.Representante_bairro, pj.Representante_cidade,
		pj.Representante_estado, pj.Representante_documento, pj.Representante_data_emissao,
		pj.Representante_orgao_emissor, pj.Representante_estado_emissor, pj.Datahora,
		pj.Usado, pj.Datahora_usado, pj.Arquivo, pj.Aparelho, pj.Permite_criacao_automatica,
		pj.Senha, pj.Imagecnpj, pj.Imagecomprovante, pj.Imagecontrato, pj.Imagecomprovante_endereco,
		pj.Representante_imageself, pj.Representante_imagedoc, pj.Representante_imagedoc_verso,
		pj.Representante_imagecomprovante_endereco, pj.Representante, pj.Representante_politico,
		pj.Representante_idoc, pj.Representante_procuracao,
	)
	return err
}
