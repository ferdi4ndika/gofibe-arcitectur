package routes

import (
	"fiber-ferdi/domains/model/handlers"
	"fiber-ferdi/middleware"

	"github.com/gofiber/fiber/v2"
)

func ModelRoute(app *fiber.App) {
	group := app.Group("/api/master/model", middleware.Authentication("Data"))
	group.Get("/", handlers.GetAllModel)
	group.Get("/:id", handlers.GetModelById)
	group.Post("/", handlers.StoreModel)
	group.Put("/:id", handlers.UpdateModel)
	group.Delete("/:id", handlers.DeleteModel)
}
