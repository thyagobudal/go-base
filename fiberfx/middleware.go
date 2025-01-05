package fiberfx

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RegisterMiddlewares registra middlewares globais.
func RegisterMiddlewares(app *fiber.App, logger *zap.Logger) {
	app.Use(func(c *fiber.Ctx) error {
		logger.Info("Request received", zap.String("path", c.Path()))
		return c.Next()
	})
}
