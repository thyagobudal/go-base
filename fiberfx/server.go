package fiberfx

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Chave para o logger no contexto local
const LoggerKey = "logger"

// ServerParams são os parâmetros necessários para criar o servidor
type ServerParams struct {
	fx.In

	Config  Config
	Modules []*Module `group:"modules"`
}

// NewFxApp cria uma nova aplicação fx
func NewFxApp(cfg Config, modules ...*Module) *fx.App {
	var moduleOpts []fx.Option

	// Adiciona providers padrão
	moduleOpts = append(moduleOpts,
		fx.Provide(
			func() Config { return cfg },
			NewLogger,
			NewServer, // Adiciona NewServer como provider
		),
	)

	// Registra o primeiro módulo
	if len(modules) > 0 {
		moduleOpts = append(moduleOpts, fx.Provide(fx.Annotated{
			Group:  "modules",
			Target: func() *Module { return modules[0] },
		}))
	}

	// Registra cada módulo adicional
	for _, m := range modules[1:] {
		moduleOpts = append(moduleOpts, fx.Provide(fx.Annotated{
			Group:  "modules",
			Target: func() *Module { return m },
		}))
	}

	// Invoca o servidor para garantir que ele seja iniciado
	moduleOpts = append(moduleOpts, fx.Invoke(func(*fiber.App) {}))

	return fx.New(moduleOpts...)
}

// NewServer agora é um construtor fx
func NewServer(lc fx.Lifecycle, params ServerParams) *fiber.App {
	app := fiber.New()

	// Configura o logger
	logger, err := NewLogger(params.Config.ServiceName, params.Config.Environment)
	if err != nil {
		panic(err)
	}

	// Armazena o logger no contexto local do Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(LoggerKey, logger)
		return c.Next()
	})

	// Registra middlewares globais
	RegisterMiddlewares(app, logger)

	// Registra o Elastic APM se habilitado
	RegisterAPMMiddleware(app, params.Config.EnableAPM, params.Config.ServiceName)

	RegisterSwagger(app)

	// Register module-specific middlewares first
	for _, module := range params.Modules {
		for _, middleware := range module.Middlewares {
			middleware(app)
		}
	}

	// Registra as rotas de cada módulo
	for _, module := range params.Modules {
		for _, route := range module.Routes {
			route(app)
		}
	}

	// Adiciona hooks do lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(":" + params.Config.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
