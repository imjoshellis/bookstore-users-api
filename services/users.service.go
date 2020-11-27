package services

import (
	"bookstore/users/domain/users"
	"bookstore/users/utils/errors"
)

// CreateUser is the service for creating users
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := errors.NewBadRequestError("")
	return &user, nil
}
