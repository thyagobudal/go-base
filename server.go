package baseserver

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
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

func ServerModule(port string) fx.Option {
	return fx.Invoke(func(app *fiber.App, lifecycle fx.Lifecycle) {
		StartServer(app, lifecycle, port)
	})
}