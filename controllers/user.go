package controllers

import (
    "github.com/gin-gonic/gin"
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