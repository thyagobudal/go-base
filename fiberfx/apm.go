package fiberfx

import (
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmfiber"
)

func RegisterAPMMiddleware(app *fiber.App, enableAPM bool, serviceName string) {
	if enableAPM {
		apm.DefaultTracer.Service.Name = serviceName
		app.Use(apmfiber.Middleware(apmfiber.WithTracer(apm.DefaultTracer)))
	}
}
