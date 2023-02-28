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

	// Listen on Port
	app.Listen(":3000")

}
