package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	documentationController "github.com/rodrigoRVSN/beeus-api/src/application/controller/documentation"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/controller/user"
	documentationUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/documentation"
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
	"github.com/rodrigoRVSN/beeus-api/src/config"
	documentationRepository "github.com/rodrigoRVSN/beeus-api/src/infra/repository/documentation"
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

	documentationRepository := documentationRepository.NewDocumentationRepository(db)
	documentationUseCase := documentationUseCase.NewDocumentationUseCase(documentationRepository, userUseCase)
	documentationController := documentationController.NewDocumentationController(*documentationUseCase)

	app.Use(cors.New())
	routes.UserRoutes(app, userController)
	routes.DocumentationRoutes(app, documentationController)

	app.Listen(":" + os.Getenv("PORT"))
}
