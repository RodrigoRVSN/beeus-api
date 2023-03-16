package main

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":4444")
}
