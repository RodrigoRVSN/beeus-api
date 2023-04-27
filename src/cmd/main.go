package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/controller/user"
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
	"github.com/rodrigoRVSN/beeus-api/src/config"
	userRepository "github.com/rodrigoRVSN/beeus-api/src/infra/repository/user"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	app := fiber.New()

	config.LoadEnv()
	db := config.ConnectDb()

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	userController := userController.NewUserController(*userUseCase)

	app.Use(cors.New())
	routes.UserRoutes(app, userController)

	app.Listen(":" + os.Getenv("PORT"))
}
