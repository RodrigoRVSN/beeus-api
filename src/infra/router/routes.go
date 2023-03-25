package routes

import (
	"github.com/gofiber/fiber/v2"
	authController "github.com/rodrigoRVSN/beeus-api/src/app/auth/controllers"
	userController "github.com/rodrigoRVSN/beeus-api/src/app/users/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", userController.ListUsers)

	app.Post("/auth/register", authController.CreateUser)
	app.Post("/auth/login", authController.SignInUser)
}
