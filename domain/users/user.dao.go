package users

import (
	"bookstore/users/utils/errors"
	"fmt"
)

var userDb = make(map[int64]*User)

// Get tries to get a user from the database
func (u *User) Get() *errors.RestErr {
	result := userDb[u.ID]
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
	current := userDb[u.ID]
	if current != nil {
		msg := fmt.Sprintf("user %d already exists", u.ID)
		return errors.NewBadRequestError(msg)
	}
	userDb[u.ID] = u
	result := userDb[u.ID]
	if result == nil {
		msg := "there was a problem saving to the db"
		return errors.NewServerError(msg)
	}
	return nil
}
