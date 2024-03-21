package routes

import (
	"fiber-ferdi/domains/product/handlers"
	"fiber-ferdi/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App) {
	group := app.Group("/api/master/product", middleware.Authentication("Data"))
	group.Get("/", handlers.GetAllProduct)
	group.Get("/:id", handlers.GetProductById)
	group.Post("/", handlers.StoreProduct)
	group.Put("/:id", handlers.UpdateProduct)
	group.Delete("/:id", handlers.DeleteProduct)
}
