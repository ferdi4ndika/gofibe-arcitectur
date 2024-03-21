package routes

import (
	"fiber-ferdi/domains/auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Post("api/login", handlers.Login)
	app.Post("api/register", handlers.Register)
}
