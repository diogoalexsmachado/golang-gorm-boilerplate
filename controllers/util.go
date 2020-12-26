package controllers

import (
    "github.com/gin-gonic/gin"
    "strconv"
)

func page(c *gin.Context) int {
    page := c.Query("p")
    var p int
    p, err := strconv.Atoi(page)
    if err != nil {
        p = 0
    }
    return p
}
