package server

import (
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmfiber/v2"
	"go.uber.org/fx"
)

func NewTracer() (*apm.Tracer, error) {
	tracer, err := apm.NewTracer("base-server", "1.0")
	if err != nil {
		return nil, err
	}
	apm.DefaultTracer = tracer
	return tracer, nil
}

func TracingModule() fx.Option {
	return fx.Provide(NewTracer)
}

func RegisterAPM(app *fiber.App) {
	app.Use(apmfiber.Middleware())
}
