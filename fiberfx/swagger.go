package fiberfx

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// RegisterSwagger registra o Swagger no Fiber.
func RegisterSwagger(app *fiber.App) {
	// Serve static files from swagger folder
	app.Static("/docs", "../docs")

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/docs/swagger.json",
	}))
}
