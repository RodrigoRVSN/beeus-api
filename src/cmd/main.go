package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/config"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	app := fiber.New()

	config.LoadConfig()

	routes.SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
