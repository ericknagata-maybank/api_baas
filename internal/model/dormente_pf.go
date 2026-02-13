package model

type DormentePf struct {
	ID            int64  `json:"id"`
	CPF           string `json:"cpf"`
	Documento     string `json:"documento"`
	DataEmissao   string `json:"data_emissao"`
	OrgaoEmissor  string `json:"orgao_emissor"`
	EstadoEmissor string `json:"estado_emissor"`

	TipoPassaporte        *string `json:"tipo_passaporte"`
	PaisEmissorPassaporte *string `json:"pais_emissor_passaporte"`
	Nacionalidade         *string `json:"nacionalidade_passaporte"`
	Naturalidade          *string `json:"naturalidade_passaporte"`
	ValidadePassaporte    *string `json:"validade_passaporte"`

	Nome           string  `json:"nome"`
	Idoc           *string `json:"idoc"`
	Sexo           string  `json:"sexo"`
	DataNascimento string  `json:"data_nascimento"`
	EstadoCivil    string  `json:"estado_civil"`
	Politico       bool    `json:"politico"`
	Renda          float64 `json:"renda"`
	NomeMae        string  `json:"nome_mae"`

	Cep         string  `json:"cep"`
	Numero      string  `json:"numero"`
	Endereco    string  `json:"endereco"`
	Estado      string  `json:"estado"`
	Cidade      string  `json:"cidade"`
	Bairro      string  `json:"bairro"`
	Complemento *string `json:"complemento"`

	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Senha    string `json:"senha"`

	ConfirmationURL  string `json:"confirmation_url"`
	DataHoraCadastro string `json:"data_hora_cadastro"`
}
