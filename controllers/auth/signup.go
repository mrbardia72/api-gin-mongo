package auth


import (
    "github.com/gin-gonic/gin"
    "github.com/mrbardia72/api-gin-mongo/forms"
    repo_user "github.com/mrbardia72/api-gin-mongo/repository/user"
) 
 

// Signup controller handles registering a user
func Signup(c *gin.Context) {

    var data forms.SignupUserCommand
    if c.BindJSON(&data) != nil {
        c.JSON(406, gin.H{"message": "Provide relevant fields"})
        c.Abort()
        return
    }

    result, _ := repo_user.GetUserByEmail(data.Email)
    if result.Email != "" {
        c.JSON(403, gin.H{"message": "Email is already in use"})
        c.Abort()
        return
    }

    err := repo_user.Signup(data)
    if err != nil {
        c.JSON(400, gin.H{"message": "Problem creating an account"})
        c.Abort()
        return
    }

    c.JSON(201, gin.H{"message": "New user account registered"})
}
