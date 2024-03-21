package main

import (
	"fiber-ferdi/api"
	"fiber-ferdi/configs"
	"fiber-ferdi/domains/plant/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

    err := godotenv.Load()
		if err != nil {
			log.Fatal("Gagal memuat file .env")
		}
		
	app := fiber.New()
    app.Static("/real-time", "./html/jajal.html")
	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization ",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		  ExposeHeaders:  "Link, X-Pagination",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	// Database connection
	configs.ConnDB()

	// Welcome message
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rest Api with Fiber Golang")
	})
	
	app.Get("/ws", handlers.HandleWebSocket)
	
    // Route
	api.SetupRoutes(app)

	// Start the server
    port := os.Getenv("PORT")
    if port == "" {
        // Port default jika tidak diatur
        port = "3000"
    }

	app.Listen(":" + port)
}
