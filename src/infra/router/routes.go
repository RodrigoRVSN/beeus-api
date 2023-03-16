package routes

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/rodrigoRVSN/beeus-api/src/models/users/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", userController.ListUsers)
	app.Post("/users/register", userController.CreateUser)
}
