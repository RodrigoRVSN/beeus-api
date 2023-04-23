package routes

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/controller/user"
)

func UserRoutes(app *fiber.App, userController *userController.UserController) {
	app.Get("/users", userController.FindAllUsers)

	app.Post("/auth/login", userController.SignInUser)
	app.Post("/auth/register", userController.CreateUser)
}
