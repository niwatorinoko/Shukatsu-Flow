package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"shukatsu-flow/api/internal/config"
	postgresdb "shukatsu-flow/api/internal/infrastructure/db/postgres"
	postgresrepository "shukatsu-flow/api/internal/infrastructure/repository/postgres"
	httpHandler "shukatsu-flow/api/internal/interface/http/handler"
	httpRouter "shukatsu-flow/api/internal/interface/http/router"
	companyUsecase "shukatsu-flow/api/internal/usecase/company"
)

func main() {
	if envLoadError := config.LoadDotEnv(".env", "services/api/.env"); envLoadError != nil {
		log.Fatal(envLoadError)
	}

	echoServer := echo.New()

	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	databaseConnection, databaseConnectionError := postgresdb.NewConnection()
	if databaseConnectionError != nil {
		log.Fatal(databaseConnectionError)
	}
	defer databaseConnection.Close()

	companyRepository := postgresrepository.NewCompanyRepository(databaseConnection)
	companyApplicationUsecase := companyUsecase.NewUsecase(companyRepository)
	companyHandler := httpHandler.NewCompanyHandler(companyApplicationUsecase)

	httpRouter.RegisterRoutes(echoServer, companyHandler)

	appPort := config.GetEnv("APP_PORT", "8080")

	log.Fatal(echoServer.Start(":" + appPort))
}
