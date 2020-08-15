package user

import (
    "gopkg.in/mgo.v2/bson"
	"github.com/mrbardia72/api-gin-mongo/models"
)

// all user
func AllUser() (user []models.User,err error)  {
    // Connect to the user collection
    collection := models.DBConnect.Use(models.DatabaseName, "user")
    // Assign result to error object while saving user
    err = collection.Find(bson.M{}).All(&user) //all profiles
    // err = collection.Find(bson.M{"name": id}).One(&user)
    return user, err

}