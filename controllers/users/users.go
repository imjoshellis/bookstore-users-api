package users

import (
	"bookstore/users/domain/users"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser returns an individual user
func GetUser(c *gin.Context) {
	// ???: implement response
	c.String(http.StatusNotImplemented, "")
}

// CreateUser creates a user
func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// ???: handle error
		return
	}
	fmt.Println(user, string(bytes), err)
	// ???: implement response
	c.String(http.StatusNotImplemented, "")
}

// SearchUser searches for a user
func SearchUser(c *gin.Context) {
	// ???: implement response
	c.String(http.StatusNotImplemented, "")
}
