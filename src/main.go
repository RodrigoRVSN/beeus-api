package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	database "github.com/rodrigoRVSN/beeus-api/src/infra"
	loadEnv "github.com/rodrigoRVSN/beeus-api/src/infra/config"
)

func main() {
	loadEnv.LoadEnv()

	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("vasdjiasdhasioasdioasdjioasdjioasdjiopasj")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
