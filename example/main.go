package main

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	baseserver "github.com/thyagobudal/go-base"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
	"go.uber.org/fx"
)

var tracer *apm.Tracer

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample server for a Fiber application.
// @host localhost:8081
// @BasePath /
func RegisterAppRoutes(app *fiber.App) {
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy"})
	})

	app.Get("/api/trace", func(c *fiber.Ctx) error {
		tx := tracer.StartTransaction("handleRequest", "request")
		defer tx.End()

		ctx := apm.ContextWithTransaction(c.Context(), tx)
		logrus.WithContext(ctx).Info("Handling /api/trace request")

		result := functionA(ctx)
		return c.JSON(fiber.Map{"result": result})
	})
}

func functionA(ctx context.Context) string {
	span, ctx := apm.StartSpan(ctx, "functionA", "custom")
	defer span.End()

	logrus.WithContext(ctx).Info("Executing functionA")
	time.Sleep(1 * time.Second)

	return functionB(ctx)
}

func functionB(ctx context.Context) string {
	span, ctx := apm.StartSpan(ctx, "functionB", "sql")
	defer span.End()

	logrus.WithContext(ctx).Info("Executing functionB")
	time.Sleep(1 * time.Second)

	return functionC(ctx)
}

func functionC(ctx context.Context) string {
	span, ctx := apm.StartSpan(ctx, "functionC", "custom")
	defer span.End()
	logrus.WithContext(ctx).Info("Executing functionC")
	time.Sleep(500 * time.Millisecond)

	return "Tracing complete"
}

func AppRoutesModule() fx.Option {
	return fx.Invoke(RegisterAppRoutes)
}

func main() {
	logrus.AddHook(&apmlogrus.Hook{})

	tracer = apm.DefaultTracer

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	app := baseserver.NewApp(
		port,
		AppRoutesModule(),
	)

	app.Run()
}
