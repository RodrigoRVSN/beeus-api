package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	configs "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	configs.LoadEnv()

	configs.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
