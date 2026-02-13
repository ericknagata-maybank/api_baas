package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/ericknagata-maybank/api_baas/internal/config"
	"github.com/ericknagata-maybank/api_baas/internal/controllers"
	"github.com/ericknagata-maybank/api_baas/internal/http/routes"
	"github.com/ericknagata-maybank/api_baas/internal/repository"
	"github.com/ericknagata-maybank/api_baas/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewDormentePfRepository(db)
	services := services.NewAcessoService(repo)
	controller := controllers.NewAcessoController(services)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Register(e, controller)

	e.Logger.Fatal(e.Start(":8080"))
}
