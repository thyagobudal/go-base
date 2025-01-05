package fiberfx

import "github.com/gofiber/fiber/v2"

// Module representa um módulo da aplicação
type Module struct {
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
