package user

import (
    "gopkg.in/mgo.v2/bson"
	"github.com/mrbardia72/api-gin-mongo/models"
)

// UserModel defines the model structure
type UserModel struct{}

// all user
func (u *UserModel) AllUsers() (user []models.User,err error)  {
    // Connect to the user collection
    collection := models.DBConnect.Use(models.DatabaseName, "user")
    // Assign result to error object while saving user
    err = collection.Find(bson.M{}).All(&user)
    return user, err

}