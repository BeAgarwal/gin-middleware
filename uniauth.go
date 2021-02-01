package gin_middleware

import (
	"fmt"
	"github.com/UniAuth/gin-middleware/minion"
	"github.com/UniAuth/gin-middleware/models"
	"github.com/gin-gonic/gin"
)

var uniAuthInstance models.UniAuthInstance
var appConfigs []models.ApplicationConfig

// AuthenticateMiddleware redirects user to authorization page
func AuthenticateMiddleware(configName string) gin.HandlerFunc {
	// return instance of gin middleware
	return func(c *gin.Context) {
		// check if configName is valid
		config, err := minion.GetConfigByName(appConfigs, configName)
		if err != nil {
			fmt.Println(fmt.Sprintf("Config named %s was not found, using %s", configName, config.Name))
		}

		// generate url of UniAuth Server
		authLink := fmt.Sprintf("%s/%s?client_id=%s&redirect_uri=%s", config.Url, config.AuthEndpoint, config.ClientId, config.RedirectUri)

		// take a permanent redirect to auth server
		c.Redirect(302, authLink)
	}
}

// CallBackMiddleware handles the response from OAuth server and fetches user details
func CallBackMiddleware(configName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	mojo for receiving callback from user
	}
}

// Init acts as a constructor to initiate the auth module
func Init(configs []models.ApplicationConfig) models.UniAuthInstance {
	uniAuthInstance = models.UniAuthInstance{
		Callback:     CallBackMiddleware,
		Authenticate: AuthenticateMiddleware,
	}
	fmt.Println("UniAuth Module Injected")
	appConfigs = configs

	return uniAuthInstance
}
