package controllers

import (
    "github.com/gin-gonic/gin"
    repo_user "github.com/mrbardia72/api-gin-mongo/repository/user"
) 
 
// Import the userModel from the models
// var userModel = new(models.UserModel)

// AllUsers controller handles
func AllUsers(c *gin.Context) {

    res,err := repo_user.AllUsers()
    if err != nil {
        c.JSON(400, gin.H{"error": "empty table"})
        c.Abort()
        return
    }
    c.JSON(201, res)
} 