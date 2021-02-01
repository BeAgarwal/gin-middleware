package gin_middleware

import (
	"fmt"
	"github.com/UniAuth/gin-middleware/models"
	"github.com/gin-gonic/gin"
)

var uniAuthInstance models.UniAuthInstance

func saveConfigurations(c models.ConfigStruct) {
	fmt.Print(c)
}

func AuthenticateMiddleware(configName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	mojo for redirecting user to auth endpoint
	}
}

func CallBackMiddleware(configName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	mojo for receiving callback from user
	}
}

func Init(config models.ConfigStruct) models.UniAuthInstance {
	uniAuthInstance = models.UniAuthInstance{
		Callback:     CallBackMiddleware,
		Authenticate: AuthenticateMiddleware,
	}
	fmt.Println("UniAuth Module Injected")
	saveConfigurations(config)

	return uniAuthInstance
}
