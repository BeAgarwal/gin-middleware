package main

import (
	"fmt"
	gin_middleware "github.com/UniAuth/gin-middleware"
	"github.com/UniAuth/gin-middleware/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// initialize gin
	app := gin.New()
	app.Use(gin.Logger())

	// initialize multiple UniAuth Application Configurations
	var configs = []models.ApplicationConfig{
		models.ApplicationConfig{
			Name:            "server1",
			Url:             "http://localhost:5000",
			ClientId:        "600ee6ec924dd75267384cb4",
			ClientSecret:    "986727d0-c253-4adb-a9b8-c233a89cdb25",
			RedirectUri:     "http://localhost:3000/callback",
			AuthEndpoint:    "account/o/login",
			ProfileEndpoint: "account/o/access",
			ProfileProcessor: func(details string) {
				fmt.Println(details)
			},
		},

		models.ApplicationConfig{
			Name:            "server2",
			Url:             "http://localhost:4000",
			ClientId:        "600ee6ec924dd714752a8f74",
			ClientSecret:    "986727d0-147a-258a-3a6a-q7s4f5s8z3a6",
			RedirectUri:     "http://localhost:3000/callback",
			AuthEndpoint:    "account/o/login",
			ProfileEndpoint: "account/o/access",
			ProfileProcessor: func(details string) {
				fmt.Println(details)
			},
		},
	}

	// initialize uniAuth middleware instances
	uniAuth := gin_middleware.Init(configs)

	// landing page
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"alive": true})
	})

	// url to initiate login
	app.GET("/login", uniAuth.Authenticate("server1"))

	app.Run(":3000")
}
