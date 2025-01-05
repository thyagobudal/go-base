package main

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
		features.NewHealthFeature(),
		features.NewErrorSimulationFeature(),
	)

	app.Run()
}
