package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/mrbardia72/api-gin-mongo/forms"
    repo_user "github.com/mrbardia72/api-gin-mongo/repository/user"
) 

// AllUsers controller handles
func AllUsers(c *gin.Context) {
    res,err := repo_user.AllUser()
    if err != nil {
        c.JSON(400, gin.H{"error": "empty table"})
        c.Abort()
        return
    }
    c.JSON(201, res)
}

// RemoveUsers controller handles
func RemoveUsers(c *gin.Context)  {
    
	email := c.Params.ByName("email")
    err := repo_user.RemoveUser(email)
    if err != nil {
        c.JSON(400, gin.H{"error": "empty collection"})
        c.Abort()
        return
    }
    c.JSON(201,gin.H{"Success":"Success remove email "  +email})
}

func UpdateUsers(c *gin.Context)  {
    
    var data forms.SignupUserCommand
    if c.BindJSON(&data) != nil {
        c.JSON(406, gin.H{"message": "Provide relevant fields"})
        c.Abort()
        return
        }

    email := c.Params.ByName("email")
    err := repo_user.UpdateUser(data,email)
    if err != nil {
        c.JSON(400, gin.H{"error": "no exists  collection for update"})
        c.Abort()
        return
    }
    c.JSON(201,gin.H{"Success":"Success update profile "})
}