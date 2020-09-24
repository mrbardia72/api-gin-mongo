package user

import (
	"log"

	"github.com/mrbardia72/api-gin-mongo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// all user
func RemoveUser(email string) (err error) {
	// Connect to the user collection
	collection := models.DBConnect.Use(models.DatabaseName, "user")
	// Assign result to error object while saving user
	err = collection.Remove(bson.M{"email": email})
	if err != nil {
		switch err {
		default:
			log.Println("Failed delete user: ", err)
			return

		case mgo.ErrNotFound:
			log.Println("collection not found: ", err)
			return
		}
	}
	return nil
}
