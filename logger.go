package baseserver

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func LoggerModule() fx.Option {
	return fx.Provide(NewLogger)
}
