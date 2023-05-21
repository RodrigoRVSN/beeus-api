package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/container"
)

func SetupRoutes(app *fiber.App, container *container.Container) {
	UserRoutes(app, container.UserController)
	DocumentationRoutes(app, container.DocumentationController)
}
