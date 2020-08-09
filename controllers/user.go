package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/mrbardia72/api-gin-mongo/models"
) 
 
// Import the userModel from the models
var userModel = new(models.UserModel)

// AllUsers controller handles
func AllUsers(c *gin.Context) {

    res,err := userModel.AllUsers()
    if err != nil {
        c.JSON(400, gin.H{"error": "empty table"})
        c.Abort()
        return
    }
    c.JSON(201, res)
}