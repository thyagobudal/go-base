package fiberfx

import (
	//_ "example.com/myapi/docs" // Import necessário para o Swag CLI
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// RegisterSwagger registra o Swagger no Fiber.
func RegisterSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
