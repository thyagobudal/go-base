package features

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thyagobudal/go-base/fiberfx"
)

func NewHealthFeature() *fiberfx.Module {
	return &fiberfx.Module{
		Routes: []func(app *fiber.App){
			func(app *fiber.App) {
				app.Get("/health", func(c *fiber.Ctx) error {
					return c.JSON(fiber.Map{"status": "ok"})
				})
			},
		},
	}
}
