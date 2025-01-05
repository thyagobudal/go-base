package features

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thyagobudal/go-base/fiberfx"
)

// @Summary Health Check
// @Description Verifica o status da API
// @Tags Health
// @Produce json
// @Success 200
// @Router /api/health [get]
func healthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

func NewHealthFeature() *fiberfx.Module {
	return &fiberfx.Module{
		Routes: []func(app *fiber.App){
			func(app *fiber.App) {
				app.Get("/health", healthCheckHandler)
			},
		},
	}
}
