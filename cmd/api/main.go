package main

import (
	"net/http"

	"github.com/ericknagata-maybank/api_baas/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	routes.Register(e)

	e.Logger.Fatal(e.Start(":8080"))
}
