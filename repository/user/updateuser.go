package user


import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/mrbardia72/api-gin-mongo/models"
    "log"
    "github.com/mrbardia72/api-gin-mongo/forms"
    "github.com/mrbardia72/api-gin-mongo/helpers"
)

// all user
func UpdateUser(data forms.SignupUserCommand,email string) (err error)  {
    // var usera models.User
    // var data forms.SignupUserCommand
    collection := models.DBConnect.Use(models.DatabaseName, "user")
//  
filter := bson.M{"email": email}
	change := bson.M{
		"$set": bson.M{
            "name":        data.Name,
            "email":       data.Email,
            "password": helpers.GeneratePasswordHash([]byte(data.Password)),
		},
    }
    // 
	err = collection.Update(filter, change)
    if err != nil {
        switch err {
        default:
            log.Println("Failed update user: ", err)
            return
            
        case mgo.ErrNotFound:
            log.Println("collection not found: ", err)
            return
        }
    }
    return nil
}