package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaredtconnor/clipper-app/database"
	"github.com/jaredtconnor/clipper-app/router"
)

func main() {

	app := fiber.New()

	// // Connect to database
	database.ConnectDB()

	// // Setup router
	router.SetupRouter(app)

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	// Listen on Port
	app.Listen(":3000")

}
