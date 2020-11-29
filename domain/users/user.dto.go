package users

import (
	"strings"
	"users/utils/errs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct models user
type User struct {
	ID          primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	FirstName   string             `bson:"firstName" json:"firstName"`
	LastName    string             `bson:"lastName" json:"lastName"`
	Email       string             `bson:"email" json:"email"`
	DateCreated string             `bson:"dateCreated" json:"dateCreated"`
}

// Validate checks to make sure a user is valid
func (u *User) Validate() *errs.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errs.NewBadRequestError("invalid email address")
	}
	return nil
}
