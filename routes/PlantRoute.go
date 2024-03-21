package routes

import (
	"fiber-ferdi/domains/plant/handlers"
	"fiber-ferdi/middleware"

	"github.com/gofiber/fiber/v2"
)

func PlantRoute(app *fiber.App) {
	// group := app.Group("/plant")/api/master/plant
	group := app.Group("/api/master/plant", middleware.Authentication("Plant"))
	group.Get("/", handlers.GetAllPlant)
	group.Get("/:id", handlers.GetPlantById)
	group.Post("/", handlers.StorePlant)
	group.Put("/:id", handlers.UpdatePlant)
	group.Delete("/:id", handlers.DeletePlant)
}
