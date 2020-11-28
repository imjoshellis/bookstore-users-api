// Package services has business logic
package services

import (
	"users/domain/users"
	"users/utils/errs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser is the service for creating users
func CreateUser(user users.User) (*users.User, *errs.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser is the service for getting users
func GetUser(id string) (*users.User, *errs.RestErr) {
	objectID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return nil, errs.NewBadRequestError("invalid id")
	}
	user := users.User{ID: objectID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
