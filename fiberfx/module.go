package fiberfx

import "github.com/gofiber/fiber/v2"

// Module representa um módulo da aplicação
type Module struct {
	// Providers do módulo
	Providers []interface{}
	// Invocadores do módulo
	Invokes []interface{}
	// Rotas do módulo
	Routes      []func(app *fiber.App)
	Middlewares []func(app *fiber.App)
}

// NewModule cria um novo módulo
func NewModule(opts ...ModuleOption) *Module {
	m := &Module{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// ModuleOption é uma função que configura um módulo
type ModuleOption func(*Module)

// WithProviders adiciona providers ao módulo
func WithProviders(providers ...interface{}) ModuleOption {
	return func(m *Module) {
		m.Providers = append(m.Providers, providers...)
	}
}

// WithInvokes adiciona invokes ao módulo
func WithInvokes(invokes ...interface{}) ModuleOption {
	return func(m *Module) {
		m.Invokes = append(m.Invokes, invokes...)
	}
}
