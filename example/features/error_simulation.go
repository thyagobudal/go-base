package features

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/thyagobudal/go-base/fiberfx"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// @Summary Simula um erro
// @Description Endpoint para simular um erro e testar o APM
// @Tags Error
// @Produce json
// @Success 200
// @Router /api/simulate-error [get]
func simulateErrorHandler(c *fiber.Ctx) error {
	logger := c.Locals(fiberfx.LoggerKey).(*zap.Logger)

	fiberfx.LogTrace(c.Context(), logger, "iniciando simulação de erro",
		zap.String("path", c.Path()),
		zap.String("method", c.Method()),
	)

	err := errors.New("erro simulado para teste do APM")

	logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel)).
		Error("erro simulado",
			zap.Error(err),
			zap.String("path", c.Path()),
			zap.String("method", c.Method()),
		)

	if e := apm.CaptureError(c.Context(), err); e != nil {
		e.Send()
	}

	fiberfx.LogTrace(c.Context(), logger, "finalizando simulação de erro",
		zap.String("error", err.Error()),
	)

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewErrorSimulationFeature() *fiberfx.Module {
	return &fiberfx.Module{
		Routes: []func(app *fiber.App){
			func(app *fiber.App) {
				app.Get("/simulate-error", simulateErrorHandler)
			},
		},
	}
}
