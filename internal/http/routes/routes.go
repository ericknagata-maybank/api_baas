package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/ericknagata-maybank/api_baas/internal/controllers"
)

func Register(e *echo.Echo) {

	println("Registering routes.")

	v1 := e.Group("/v1")

	acessoController := controllers.NewAcessoController()

	// Porta de entrada externa do BaaS
	v1.POST("/accounts", acessoController.CreateAccount)
}
