package controllers

import (
    "github.com/gin-gonic/gin"
)
 
func Default(c *gin.Context) {
    c.JSON(200, gin.H{"message": "test Hello world test"})
}
