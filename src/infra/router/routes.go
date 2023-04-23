package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	authController "github.com/rodrigoRVSN/beeus-api/src/application/auth/controllers"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/users/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Get("/users", userController.ListUsers)

	app.Post("/auth/register", authController.CreateUser)
	app.Post("/auth/login", authController.SignInUser)
}
