package app

import (
	"bookstore/users/controllers/ping"
	"bookstore/users/controllers/users"
)

// MapRoutes maps routes to controllers
func MapRoutes() {
	// Ping route for connection testing
	router.GET("/ping", ping.Ping)

	router.GET("/users/:userId", users.GetUser)
	router.POST("/users", users.CreateUser)
}
