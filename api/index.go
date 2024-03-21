package api

import (
	"fiber-ferdi/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
 
    routes.AuthRoute(app)
    routes.ModelRoute(app)
    routes.OpRoute(app)
    routes.PartRoute(app)
    routes.PlantRoute(app)
    routes.ProductRoute(app)
}
