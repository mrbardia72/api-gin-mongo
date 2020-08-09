package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/mrbardia72/api-gin-mongo/forms"
    "github.com/mrbardia72/api-gin-mongo/models"
    "github.com/mrbardia72/api-gin-mongo/helpers"
    "github.com/mrbardia72/api-gin-mongo/services"
) 
 
// Import the userModel from the models
var userLoginModel = new(models.UserModel)

// Login allows a user to login a user and get
// access token
func Login(c *gin.Context) {
    var data forms.LoginUserCommand

    // Bind the request body data to var data and check if all details are provided
    if c.BindJSON(&data) != nil {
        c.JSON(406, gin.H{"message": "Provide required details"})
        c.Abort()
        return
    }

    result, err := userLoginModel.GetUserByEmail(data.Email)

    if result.Email == "" {
        c.JSON(404, gin.H{"message": "User account was not found"})
        c.Abort()
        return
    }

    if err != nil {
        c.JSON(400, gin.H{"message": "Problem logging into your account"})
        c.Abort()
        return
    }

    // Get the hashed password from the saved document
    hashedPassword := []byte(result.Password)
    // Get the password provided in the request.body
    password := []byte(data.Password)

    err = helpers.PasswordCompare(password, hashedPassword)

    if err != nil {
        c.JSON(403, gin.H{"message": "Invalid user credentials"})
        c.Abort()
        return
    }

    jwtToken, err2 := services.GenerateToken(data.Email)

    // If we fail to generate token for access
    if err2 != nil {
        c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
        c.Abort()
        return
    }

    c.JSON(200, gin.H{"message": "Log in success", "token": jwtToken})
}