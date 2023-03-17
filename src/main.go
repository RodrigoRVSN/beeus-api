package main

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/rodrigoRVSN/beeus-api/src/infra"
	"os"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("vasdjiasdhasioasdioasdjioasdjioasdjiopasj")
	})

	app.Listen(":%s", os.Getenv("PORT"))
}
