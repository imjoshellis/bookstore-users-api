package app

import (
	"users/controllers/ping"
	"users/controllers/users"

	"github.com/gofiber/fiber/v2"
)

// MapRoutes maps routes to controllers
func MapRoutes(app *fiber.App) {
	// Ping route for connection testing
	app.Get("/ping", ping.Ping)

	app.Post("/users", users.CreateUser)
	app.Get("/users/:id", users.GetUser)
}
