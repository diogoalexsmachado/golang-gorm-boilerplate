package controllers

import (
    "github.com/gin-gonic/gin"
    "newsfeed-go/middlewares"
)

const userId string = "foobar420"

func Auth(c *gin.Context) {
    token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId)
    if err != nil {

    }
    c.JSON(200, token)
}

