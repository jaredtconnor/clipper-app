package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaredtconnor/clipper-app/database"
)

func main() {

	app := fiber.New()

	// database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	app.Listen(":3000")

}
