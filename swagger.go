package server

import (
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

func RegisterSwagger(app *fiber.App) {

	cmd := exec.Command("swag", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("swag init failed: %v", err)
	}

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/docs/swagger.json",
		DeepLinking: false,
	}))
}

func SwaggerModule() fx.Option {
	return fx.Invoke(RegisterSwagger)
}
