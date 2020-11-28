// Package ping is for testing connection
package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping sends pong to any request
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
