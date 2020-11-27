package app

import (
	"bookstore/users/controllers/ping"
	"bookstore/users/controllers/users"
)

// MapRoutes maps routes to controllers
func MapRoutes() {
	// Ping route for connection testing
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
}
