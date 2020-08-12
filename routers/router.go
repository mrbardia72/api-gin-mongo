package routers

import (
	"github.com/gin-gonic/gin"
    "github.com/mrbardia72/api-gin-mongo/controllers"
    "github.com/mrbardia72/api-gin-mongo/controllers/auth"
)

func RouterApp()  {
	r := gin.Default()

    v1 := r.Group("/api/v1")
    {
        v1.POST("/signup", auth.Signup)
        v1.POST("/login", auth.Login)
        v1.GET("/all", controllers.AllUsers)
        v1.DELETE("/remove/:email", controllers.RemoveUsers)
        v1.PUT("/edit/:email",controllers.UpdateUsers)
    }

    // Handle error response when a route is not defined
    r.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"code":404,"message": "Not found"})
    })
    r.Run(":5000")
}