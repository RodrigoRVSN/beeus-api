package routes

import (
	"github.com/gofiber/fiber/v2"
	documentationController "github.com/rodrigoRVSN/beeus-api/src/application/controller/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	"github.com/rodrigoRVSN/beeus-api/src/infra/middleware"
)

func DocumentationRoutes(app *fiber.App, documentationController *documentationController.DocumentationController) {
	app.Get("/documentation", context.AdaptHandler(documentationController.ListAllDocumentations))
	app.Get("/documentation/:documentationId", middleware.AuthMiddleware, context.AdaptHandler(documentationController.GetDocumentation))
	app.Post("/documentation", middleware.AuthMiddleware, context.AdaptHandler(documentationController.CreateDocumentation))
	app.Delete("/documentation/:documentationId", middleware.AuthMiddleware, context.AdaptHandler(documentationController.DeleteDocumentation))
	app.Put("/documentation/:documentationId", middleware.AuthMiddleware, context.AdaptHandler(documentationController.EditDocumentation))
}
