package clippingRoutes

import "github.com/gofiber/fiber/v2"

func SetupClippingRoutes(router fiber.Router) {

	clipping := router.Group("/clipping")

	// Create a clipping
	clipping.Post("/", func(c *fiber.Ctx) error {})

	// read all clippings
	clipping.Get("/", func(c *fiber.Ctx) error {})

	// Read specific clipping
	clipping.Get("/:clippingID", func(c *fiber.Ctx) error {})

	// Updated specific clipping
	clipping.Put("/:clippingID", func(c *fiber.Ctx) error {})

	// Delete specific clipping
	clipping.Delete("/:clippingID", func(c *fiber.Ctx) error {})
}
