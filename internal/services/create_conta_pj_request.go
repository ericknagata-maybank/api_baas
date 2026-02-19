package services

type CreateContaPjRequest struct {
	RazaoSocial                 string `json:"razaoSocial"`
	NomeFantasia                string `json:"nomeFantasia"`
	CNPJ                        string `json:"cnpj"`
	Agencia                     string `json:"agencia"`
	CNAE                        string `json:"cnae"`
	Faturamento                 string `json:"faturamento"`
	Contribuicao                string `json:"contribuicao"`
	InscricaoEstadual           string `json:"inscricaoEstadual"`
	Telefone                    string `json:"telefone"`
	Email                       string `json:"email"`
	DataAbertura                string `json:"dataAbertura"`
	Cep                         string `json:"cep"`
	Endereco                    string `json:"endereco"`
	Numero                      string `json:"numero"`
	Complemento                 string `json:"complemento"`
	Bairro                      string `json:"bairro"`
	Cidade                      string `json:"cidade"`
	Estado                      string `json:"estado"`
	RepresentanteNome           string `json:"representanteNome"`
	RepresentanteCPF            string `json:"representanteCpf"`
	RepresentanteDataNascimento string `json:"representanteDataNascimento"`
	RepresentanteTelefone       string `json:"representanteTelefone"`
	RepresentanteNomeMae        string `json:"representanteNomeMae"`
	RepresentanteSexo           string `json:"representanteSexo"`
	RepresentanteEmail          string `json:"representanteEmail"`
	RepresentanteEstadoCivil    string `json:"representanteEstadoCivil"`
	RepresentanteCep            string `json:"representanteCep"`
	RepresentanteEndereco       string `json:"representanteEndereco"`
	RepresentanteNumero         string `json:"representanteNumero"`
	RepresentanteComplemento    string `json:"representanteComplemento"`
	RepresentanteBairro         string `json:"representanteBairro"`
	RepresentanteCidade         string `json:"representanteCidade"`
	RepresentanteEstado         string `json:"representanteEstado"`
	RepresentanteDocumento      string `json:"representanteDocumento"`
	RepresentanteDataEmissao    string `json:"representanteDataEmissao"`
	RepresentanteOrgaoEmissor   string `json:"representanteOrgaoEmissor"`
	RepresentanteEstadoEmissor  string `json:"representanteEstadoEmissor"`
	Senha                       string `json:"senha"`
}
