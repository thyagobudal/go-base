package server

import (
	"go.uber.org/fx"
)

func BaseServerModule(port string) fx.Option {
	return fx.Options(
		fx.Provide(NewFiberApp),
		ServerModule(port),
		LoggerModule(),
		TracingModule(),
	)
}

func NewApp(port string, appModules ...fx.Option) *fx.App {
	return fx.New(
		BaseServerModule(port),
		fx.Options(appModules...),
	)
}
