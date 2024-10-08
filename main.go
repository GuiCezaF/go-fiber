package main

import (
	"github.com/GuiCezaF/go-fiber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("API is running...")
		return err
	})

	app.Listen(":3000")
}
