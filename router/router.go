package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	clippingRoutes "github.com/jaredtconnor/clipper-app/internals/routes/clipping"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/api", logger.New())

	clippingRoutes.SetupClippingRoutes(api)

	fmt.Println("Setup Router")

}
