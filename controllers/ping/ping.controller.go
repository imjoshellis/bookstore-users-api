// Package ping is for testing connection
package ping

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Ping sends pong to any request
func Ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("pong")
}
