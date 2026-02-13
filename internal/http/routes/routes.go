package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/ericknagata-maybank/api_baas/internal/controllers"
)

func Register(e *echo.Echo, controller *controllers.AcessoController) {

	println("Registering routes.")

	v1 := e.Group("/v1")

	v1.POST("/accounts", controller.CreateAccount)
	v1.POST("/contas-pf", controller.CreateContaPf)
}
