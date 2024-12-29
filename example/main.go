package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	baseserver "github.com/thyagobudal/go-base"
	"go.uber.org/fx"
)

func RegisterAppRoutes(app *fiber.App) {
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy"})
	})
}

func AppRoutesModule() fx.Option {
	return fx.Invoke(RegisterAppRoutes)
}

func main() {

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
