package clippingRoutes

import (
	"github.com/gofiber/fiber/v2"
	clippingHandler "github.com/jaredtconnor/clipper-app/internal/handlers/clipping"
)

func SetupClippingRoutes(router fiber.Router) {

	clipping := router.Group("/clipping")

	// Create a clipping
	clipping.Post("/", clippingHandler.CreateClipping)

	// read all clippings
	clipping.Get("/", clippingHandler.GetClippings)

	// Read specific clipping
	clipping.Get("/:clippingId", clippingHandler.GetClipping)

	// Updated specific clipping
	clipping.Put("/:clippingId", clippingHandler.UpdateClipping)

	// Delete specific clipping
	clipping.Delete("/:clippingId", clippingHandler.DeleteClipping)
}
