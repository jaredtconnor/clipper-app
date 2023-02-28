package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	clippingRoutes "github.com/jaredtconnor/clipper-app/internal/routes/clipping"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/api", logger.New())

	clippingRoutes.SetupClippingRoutes(api)

}
