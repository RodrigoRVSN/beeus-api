package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Recebaaaaaaaaaaaaaaaaaa")
	})

	app.Listen(":4444")
}
