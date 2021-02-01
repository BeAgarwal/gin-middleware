package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"alive": true})
	})

	app.Run(":3000")
}
