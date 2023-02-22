package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {

	app := fiber.App()
	api := app.Group("api")

	user := api.Group("user")

	user.Get("/", func(c *fiber.Ctx) {})

	user.Get("/:userId", func(c *fiber.Ctx) {})

	user.Put("/:userId", func(c *fiber.Ctx) {})

	api := app.Group("/api", logger.New())

}
