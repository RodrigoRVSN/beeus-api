package routes

import (
	"github.com/gofiber/fiber/v2"
	documentationController "github.com/rodrigoRVSN/beeus-api/src/application/controller/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/infra/middleware"
)

func DocumentationRoutes(app *fiber.App, documentationController *documentationController.DocumentationController) {
	app.Get("/documentation", documentationController.ListAllDocumentations)
	app.Post("/documentation", middleware.AuthMiddleware, documentationController.CreateDocumentation)
	app.Delete("/documentation/:documentationId", middleware.AuthMiddleware, documentationController.DeleteDocumentation)
}
