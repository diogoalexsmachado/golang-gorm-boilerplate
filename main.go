package main
// only need mysql OR sqlite 
// both are included here for reference
import (
    "github.com/gin-gonic/gin"
    "newsfeed-go/controllers"
    "newsfeed-go/db"
    "newsfeed-go/middlewares"
)

var err error

func main() {
    db.Init()

    r := gin.Default()

    v1 := r.Group("/v1")
    v1.Use(middlewares.AuthHandler())
    {
        v1.GET("/people", controllers.GetPeople)
        v1.GET("/people/:id", controllers.GetPerson)
        v1.POST("/people", controllers.CreatePerson)
        v1.PUT("/people/:id", controllers.UpdatePerson)
        v1.DELETE("/people/:id", controllers.DeletePerson)
    }

    r.GET("/auth", controllers.Auth)

    r.Run(":8080")

    defer db.CloseDB()
}


