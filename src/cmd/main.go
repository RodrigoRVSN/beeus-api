package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rodrigoRVSN/beeus-api/src/config"
	"github.com/rodrigoRVSN/beeus-api/src/infra/container"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	app := fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		context.NewFiberContext(ctx)
		return ctx.Next()
	})

	config.LoadEnv()

	container := container.InitializeContainer()

	app.Use(cors.New())
	routes.SetupRoutes(app, container)

	app.Listen(":" + os.Getenv("PORT"))
}
