package features

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/thyagobudal/go-base/fiberfx"
)

// @Summary Hello com logging
// @Description Endpoint de exemplo que demonstra o uso de logging
// @Tags Logger
// @Produce json
// @Success 200
// @Router /api/logged/hello [get]
func helloLoggedHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello with logging!",
	})
}

func NewLoggerExampleFeature() *fiberfx.Module {
	return &fiberfx.Module{
		Routes: []func(app *fiber.App){
			func(app *fiber.App) {
				group := app.Group("/logged", loggerMiddleware)
				group.Get("/hello", helloLoggedHandler)
			},
		},
	}
}

func loggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("⚡️ [%s] Route accessed: %s\n", time.Now().Format(time.RFC3339), c.Path())

	err := c.Next()

	duration := time.Since(start)
	fmt.Printf("⚡️ [%s] Request completed in %v\n", time.Now().Format(time.RFC3339), duration)

	return err
}
