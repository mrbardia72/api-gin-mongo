package user

import (
	"github.com/mrbardia72/api-gin-mongo/forms"
	"github.com/mrbardia72/api-gin-mongo/helpers"
	"github.com/mrbardia72/api-gin-mongo/models"
	"gopkg.in/mgo.v2/bson"
)

// Signup handles registering a user
func Signup(data forms.SignupUserCommand) error {
	// Connect to the user collection
	collection := models.DBConnect.Use(models.DatabaseName, "user")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": helpers.GeneratePasswordHash([]byte(data.Password)),
		// This will come later when adding verification
		"is_verified": false,
	})

	return err
}
