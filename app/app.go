package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

// StartApp sets up routes and starts a gin server
func StartApp() {
	MapRoutes()
	router.Run(":8080")
}
