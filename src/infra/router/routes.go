package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/controller/user"
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
)

func SetupRoutes(app *fiber.App, uc *userUseCase.CreateUserUseCase) {
	app.Use(cors.New())

	userHandler := userController.NewUserController(*uc)

	app.Get("/users", userHandler.FindAllUsers)

	app.Post("/auth/login", userHandler.SignInUser)
	app.Post("/auth/register", userHandler.CreateUser)
}
