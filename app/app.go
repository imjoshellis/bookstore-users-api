// Package app is the entry point
package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

// StartApp sets up routes and starts a gin server
func StartApp() {
	MapRoutes()
	router.Run("localhost:8080")
}
