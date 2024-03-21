package routes

import (
	"fiber-ferdi/domains/part/handlers"
	"fiber-ferdi/middleware"

	"github.com/gofiber/fiber/v2"
)

func PartRoute(app *fiber.App) {
	group := app.Group("/api/master/part", middleware.Authentication("Data"))
	group.Get("/", handlers.GetAllPart)
	group.Get("/excle", handlers.ExportExcelHandler)
	group.Get("/:id", handlers.GetPartById)
	group.Post("/", handlers.StorePart)
	group.Put("/:id", handlers.UpdatePart)
	group.Delete("/:id", handlers.DeletePart)
	
}
