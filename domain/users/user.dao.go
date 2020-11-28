// Package users holds the users model
package users

import (
	"fmt"
	"users/utils/date"
	"users/utils/errors"
)

var userDB = make(map[int64]*User)

// Get tries to get a user from the database
func (u *User) Get() *errors.RestErr {
	result := userDB[u.ID]
	if result == nil {
		msg := fmt.Sprintf("user %d not found", u.ID)
		return errors.NewBadRequestError(msg)
	}
	u.ID = result.ID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return nil
}

// Save tries to save a user to the database
func (u *User) Save() *errors.RestErr {
	u.DateCreated = date.GetNowString()
	u.ID = int64(len(userDB)) + 1
	userDB[u.ID] = u
	result := userDB[u.ID]
	if result == nil {
		msg := "there was a problem saving to the db"
		return errors.NewServerError(msg)
	}
	return nil
}
