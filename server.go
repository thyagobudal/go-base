package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

func NewFiberApp() *fiber.App {
	return fiber.New()
}

func StartServer(app *fiber.App, lifecycle fx.Lifecycle, port string) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Starting server on :%s\n", port)
			go func() {
				if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return app.Shutdown()
		},
	})
}

func RegisterSwagger(app *fiber.App) {
	app.Static("/docs", "../docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/docs/swagger.json", // Local path to the swagger.json file
		DeepLinking: false,
	}))
}

func ServerModule(port string) fx.Option {
	return fx.Options(
		fx.Invoke(func(app *fiber.App, lifecycle fx.Lifecycle) {
			StartServer(app, lifecycle, port)
		}),
		fx.Invoke(RegisterSwagger),
	)
}
