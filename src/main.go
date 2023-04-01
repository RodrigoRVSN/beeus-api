package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	configs "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	routes "github.com/rodrigoRVSN/beeus-api/src/infra/router"
)

func main() {
	app := fiber.New()

	configs.LoadEnv()
	configs.ConnectDb()

	app.Use(cors.New())
	routes.SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
