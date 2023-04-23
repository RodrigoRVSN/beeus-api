package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
	"github.com/rodrigoRVSN/beeus-api/src/config"
	"github.com/rodrigoRVSN/beeus-api/src/infra/repository"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	app := fiber.New()

	config.LoadEnv()
	db := config.ConnectDb()

	userRepository := repository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)

	routes.SetupRoutes(app, userUseCase)

	app.Listen(":" + os.Getenv("PORT"))
}
