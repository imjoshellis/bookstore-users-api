// Package users holds the users controller
package users

import (
	"net/http"
	"users/domain/users"
	"users/services"
	"users/utils/errs"

	"github.com/gofiber/fiber/v2"
)

const (
	base    = 10
	intSize = 64
)

// GetUser returns an individual user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	res, getErr := services.GetUser(id)
	if getErr != nil {
		return c.Status(getErr.Status).JSON(getErr)
	}
	return c.Status(http.StatusOK).JSON(res)
}

// CreateUser creates a user
func CreateUser(c *fiber.Ctx) error {
	var user users.User
	if err := c.BodyParser(&user); err != nil {
		restErr := errs.NewBadRequestError("Invalid JSON Body")
		return c.Status(restErr.Status).JSON(restErr)
	}
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return c.Status(saveErr.Status).JSON(saveErr)
	}
	return c.Status(http.StatusCreated).JSON(res)
}

// SearchUser searches for a user
func SearchUser(c *fiber.Ctx) error {
	// ??? implement response
	return c.Status(http.StatusNotImplemented).SendString("")
}
