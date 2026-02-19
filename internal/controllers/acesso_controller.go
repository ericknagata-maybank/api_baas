package controllers

import (
	"net/http"

	"github.com/ericknagata-maybank/api_baas/internal/services"
	"github.com/labstack/echo/v4"
)

type AcessoController struct {
	service *services.AcessoService
}

func NewAcessoController(service *services.AcessoService) *AcessoController {
	return &AcessoController{
		service: service}
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

func (c *AcessoController) CreateContaPf(ctx echo.Context) error {
	req := new(services.CreateContaPfRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	if err := c.service.CriarContaPf(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  "db error",
			"detail": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "created",
	})

}

func (c *AcessoController) CreateContaPfj(ctx echo.Context) error {
	req := new(services.CreateContaPjRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	if err := c.service.CriarContaPj(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  "db error",
			"detail": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "created",
	})

}
