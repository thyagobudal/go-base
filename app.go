package baseserver

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AppParams struct {
	fx.In
	Config *Config
	Logger *Logger
	App    *fiber.App
}

func NewApp() fx.Option {
	return fx.Options(
		fx.Provide(NewFiberApp),
	)
}

func NewFiberApp(config *Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName: config.AppName,
	})
}
