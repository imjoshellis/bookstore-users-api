package services

import (
	"bookstore/users/domain/users"
	"bookstore/users/utils/errors"
)

// CreateUser is the service for creating users
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser is the service for getting users
func GetUser(id int64) (*users.User, *errors.RestErr) {
	if id <= 0 {
		return nil, errors.NewBadRequestError("invalid id")
	}
	user := users.User{ID: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
