package user

import (
	"github.com/mrbardia72/api-gin-mongo/models"
	"gopkg.in/mgo.v2/bson"
)

// UpdateUserPass handles updating user password
func UpdateUserPass(email string, password string) (err error) {
	
	collection := models.DBConnect.Use(models.DatabaseName, "user")
	err = collection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})
	return err
}