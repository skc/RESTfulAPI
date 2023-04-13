package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"name":    name,
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping/:name", Handler)
	r.Run(":5001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
