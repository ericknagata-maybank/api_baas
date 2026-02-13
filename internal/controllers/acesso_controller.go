package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AcessoController struct {
	// depois você injeta o service aqui (ex: AcessoService)
}

func NewAcessoController() *AcessoController {
	return &AcessoController{}
}

// Exemplo: endpoint externo para criação de conta
func (c *AcessoController) CreateAccount(ctx echo.Context) error {
	type CreateAccountRequest struct {
		Document string `json:"document" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}

	req := new(CreateAccountRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]string{
				"code":    "INVALID_BODY",
				"message": "Invalid request body",
			},
		})
	}

	// TODO: validação mais forte (validator)
	// TODO: chamar service de criação de conta
	// ex: result, err := c.service.CreateAccount(req)

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status": "accepted",
		"data": map[string]string{
			"document": req.Document,
			"name":     req.Name,
			"email":    req.Email,
		},
	})
}
