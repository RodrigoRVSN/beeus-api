package routes

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/controller/user"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	"github.com/rodrigoRVSN/beeus-api/src/infra/middleware"
)

func UserRoutes(app *fiber.App, userController userController.UserControllerInterface) {
	app.Post("/auth/login", context.AdaptHandler(userController.SignInUser))
	app.Post("/auth/register", context.AdaptHandler(userController.CreateUser))

	app.Get("/user/me", context.AdaptHandler(middleware.AuthMiddleware), context.AdaptHandler(userController.Me))

	app.Get("/users/ranking", context.AdaptHandler(middleware.AuthMiddleware), context.AdaptHandler(userController.GetRankedUsers))
}
