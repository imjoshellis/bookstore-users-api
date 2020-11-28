// Package users holds the users model
package users

import (
	"context"
	"fmt"
	"time"
	"users/data"
	"users/utils/date"
	"users/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get tries to get a user from the database
func (u *User) Get() *errors.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var res User
	err := data.UsersDB.FindOne(ctx, bson.D{{Key: "_id", Value: u.ID}}).Decode(&res)
	if err != nil {
		panic(err)
	}
	if res.FirstName == "" {
		msg := fmt.Sprintf("user %s not found", u.ID)
		return errors.NewBadRequestError(msg)
	}
	u.FirstName = res.FirstName
	u.LastName = res.LastName
	u.Email = res.Email
	u.DateCreated = res.DateCreated
	return nil
}

// Save tries to save a user to the database
func (u *User) Save() *errors.RestErr {
	u.DateCreated = date.GetNowString()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := data.UsersDB.InsertOne(ctx, bson.M{"firstName": u.FirstName, "lastName": u.LastName, "email": u.Email, "dateCreated": u.DateCreated})
	if err != nil {
		panic(err)
	}
	if res == nil {
		msg := "there was a problem saving to the db"
		return errors.NewServerError(msg)
	}
	u.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
