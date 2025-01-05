package main

// @title API Documentation
// @version 1.0
// @description This is a sample API documentation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your-email@domain.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api

import (
	"github.com/thyagobudal/go-base/example/features"
	"github.com/thyagobudal/go-base/fiberfx"
)

func main() {
	cfg := fiberfx.Config{
		Port:        "3000",
		ServiceName: "MyAPI",
		Environment: "development",
		EnableAPM:   true,
	}

	app := fiberfx.NewFxApp(cfg,
		features.GetModules()...,
	)

	app.Run()
}
