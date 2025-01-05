package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/thyagobudal/go-base/fiberfx"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// healthFeature represents the health check functionality
func healthFeature() *fiberfx.Module {
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

// errorSimulationFeature represents the error simulation functionality
func errorSimulationFeature() *fiberfx.Module {
	return &fiberfx.Module{
		Routes: []func(app *fiber.App){
			func(app *fiber.App) {
				app.Get("/simulate-error", simulateErrorHandler)
			},
		},
	}
}

// simulateErrorHandler handles the error simulation endpoint
func simulateErrorHandler(c *fiber.Ctx) error {
	// Obtém o logger do contexto da aplicação
	logger := c.Locals(fiberfx.LoggerKey).(*zap.Logger)

	// Registra trace no início da requisição
	fiberfx.LogTrace(c.Context(), logger, "iniciando simulação de erro",
		zap.String("path", c.Path()),
		zap.String("method", c.Method()),
	)

	err := errors.New("erro simulado para teste do APM")

	// Usa AddStacktrace do Zap
	logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel)).
		Error("erro simulado",
			zap.Error(err),
			zap.String("path", c.Path()),
			zap.String("method", c.Method()),
		)

	// Registra no APM
	if e := apm.CaptureError(c.Context(), err); e != nil {
		e.Send()
	}

	// Registra trace no fim da requisição
	fiberfx.LogTrace(c.Context(), logger, "finalizando simulação de erro",
		zap.String("error", err.Error()),
	)

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func main() {
	cfg := fiberfx.Config{
		Port:        "3000",
		ServiceName: "MyAPI",
		Environment: "development",
		EnableAPM:   true,
	}

	// Initialize features
	features := []*fiberfx.Module{
		healthFeature(),
		errorSimulationFeature(),
	}

	// Start application with features
	app := fiberfx.NewFxApp(cfg, features...)
	app.Run()
}
