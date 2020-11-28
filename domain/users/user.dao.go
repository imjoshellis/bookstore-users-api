// Package users holds the users model
package users

import (
	"context"
	"fmt"
	"time"
	"users/data/mongodb/usersdb"
	"users/utils/date"
	"users/utils/errs"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get tries to get a user from the database
func (u *User) Get() *errs.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: u.ID}}
	err := usersdb.Collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		msg := fmt.Sprintf("user %s not found", u.ID.Hex())
		return errs.NewBadRequestError(msg)
	}

	return nil
}

// Save tries to save a user to the database
func (u *User) Save() *errs.RestErr {
	u.DateCreated = date.GetNowString()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := usersdb.Collection.InsertOne(ctx, bson.M{"firstName": u.FirstName, "lastName": u.LastName, "email": u.Email, "dateCreated": u.DateCreated})
	if err != nil {
		log.Error(err)
		msg := "there was a problem saving to the db"
		return errs.NewServerError(msg)
	}
	u.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
