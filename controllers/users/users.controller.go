// Package users holds the users controller
package users

import (
	"bookstore/users/domain/users"
	"bookstore/users/services"
	"bookstore/users/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	base    = 10
	intSize = 64
)

// GetUser returns an individual user
func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), base, intSize)
	if err != nil {
		parseErr := errors.NewBadRequestError("invalid user id")
		c.JSON(parseErr.Status, parseErr)
		return
	}
	res, getErr := services.GetUser(id)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated, res)
	c.String(http.StatusNotImplemented, "")
}

// CreateUser creates a user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, res)
}

// SearchUser searches for a user
func SearchUser(c *gin.Context) {
	// ??? implement response
	c.String(http.StatusNotImplemented, "")
}
