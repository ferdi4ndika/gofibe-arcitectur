package routes

import (
	"fiber-ferdi/domains/op/handlers"
	"fiber-ferdi/middleware"

	"github.com/gofiber/fiber/v2"
)

func OpRoute(app *fiber.App) {
	group := app.Group("/api/master/operation", middleware.Authentication("Data"))
	group.Get("/", handlers.GetAllOp)
	group.Get("/:id", handlers.GetOpById)
	group.Post("/", handlers.StoreOp)
	group.Put("/:id", handlers.UpdateOp)
	group.Delete("/:id", handlers.DeleteOp)
}
