package fiberfx

import "github.com/gofiber/fiber/v2"

// RouteRegistrar é uma função que registra rotas no Fiber.
type RouteRegistrar func(app *fiber.App)

// RegisterRoutes registra todas as rotas fornecidas.
func RegisterRoutes(app *fiber.App, registrars ...RouteRegistrar) {
	for _, registrar := range registrars {
		registrar(app)
	}
}
