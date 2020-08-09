package user

import (
    "gopkg.in/mgo.v2/bson"
	"github.com/mrbardia72/api-gin-mongo/models"
)

// GetUserByEmail handles fetching user by email
func  GetUserByEmail(email string) (user models.User, err error) {
    // Connect to the user collection
    collection := models.DBConnect.Use(models.DatabaseName, "user")
    // Assign result to error object while saving user
    err = collection.Find(bson.M{"email": email}).One(&user)
    return user, err
}