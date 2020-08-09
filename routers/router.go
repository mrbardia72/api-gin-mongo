package routers

import (
	"github.com/gin-gonic/gin"
    "github.com/mrbardia72/api-gin-mongo/controllers"
)

func RouterApp()  {
	r := gin.Default()

    v1 := r.Group("/api/v1")
    {
        v1.POST("/signup", controllers.Signup)
        v1.POST("/login", controllers.Login)
        v1.GET("/all", controllers.AllUsers)
    }

    // Handle error response when a route is not defined
    r.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"message": "404 Not found"})
    })
    r.Run(":5000")
}