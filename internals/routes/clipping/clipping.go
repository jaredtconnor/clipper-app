package clippingRoutes

import (
	"github.com/gofiber/fiber/v2"

	clippingHandler "github.com/jaredtconnor/clipper-app/internals/handlers/clipping"
)

func SetupClippingRoutes(router fiber.Router) {

	clipping := router.Group("/clipping")

	//Creating clipping
	clipping.Post("/", clippingHandler.CreateClipping)

	// Read all clippings
	clipping.Get("/", clippingHandler.GetClippings)

	// Read one note
	clipping.Get("/:clippingId", clippingHandler.GetClipping)

	// Update one clipping
	clipping.Put("/:clippingId", clippingHandler.UpdateClipping)

	// Delete note
	clipping.Delete("/:clippingId", clippingHandler.DeleteClipping)

}
