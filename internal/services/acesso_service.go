package services

import (
	"strings"
	"time"

	"github.com/ericknagata-maybank/api_baas/internal/model"
	repositories "github.com/ericknagata-maybank/api_baas/internal/repository"
)

type AcessoService struct {
	repo *repositories.DormentePfRepository
}

func NewAcessoService(repo *repositories.DormentePfRepository) *AcessoService {
	return &AcessoService{repo: repo}
}

func normalizarGenero(genero string) string {
	g := strings.ToUpper(strings.TrimSpace(genero))

	switch g {
	case "M", "1", "MASCULINO":
		return "MASCULINO"
	case "F", "2", "FEMININO":
		return "FEMININO"
	default:
		return "MASCULINO"
	}
}

func (s *AcessoService) CriarContaPf(req *CreateContaPfRequest) error {
	pf := &model.DormentePf{
		CPF:                   req.NaturalPerson.DocumentNumber,
		Documento:             req.NaturalPerson.IdentityDocument,
		DataEmissao:           req.NaturalPerson.IssueDate,
		OrgaoEmissor:          req.NaturalPerson.IssuingAgency,
		EstadoEmissor:         req.NaturalPerson.IssuingState,
		TipoPassaporte:        req.NaturalPerson.PassportType,
		PaisEmissorPassaporte: req.NaturalPerson.PassportIssuingCountry,
		Nacionalidade:         req.NaturalPerson.PassportNationality,
		Naturalidade:          req.NaturalPerson.PassportPlaceOfBirth,
		ValidadePassaporte:    req.NaturalPerson.PassportExpirationDate,

		Nome:           req.NaturalPerson.Name,
		Idoc:           req.NaturalPerson.Idoc,
		Sexo:           normalizarGenero(req.NaturalPerson.Gender),
		DataNascimento: req.NaturalPerson.BirthDate,
		EstadoCivil:    req.NaturalPerson.IdMaritalStatus,
		Politico:       req.NaturalPerson.Political,
		Renda:          req.NaturalPerson.Rent,
		NomeMae:        req.NaturalPerson.NameMother,

		Cep:         req.Address.ZipCode,
		Numero:      req.Address.AddressNumber,
		Endereco:    req.Address.Address,
		Estado:      req.Address.State,
		Cidade:      req.Address.City,
		Bairro:      req.Address.Neighborhood,
		Complemento: req.Address.Complement,

		Email:    req.Contact.Email,
		Telefone: req.Contact.PhoneNumber,
		Senha:    req.UserAdmin.Password,

		ConfirmationURL:  req.ConfirmationUrl,
		DataHoraCadastro: time.Now().Format("2006-01-02 15:04:05"),
	}

	return s.repo.Create(pf)
}
