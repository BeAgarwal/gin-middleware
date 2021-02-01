package models

import "github.com/gin-gonic/gin"

// ProfileProcessor to process incoming user profile from server
type ProfileProcessor func(userDetails string)

type ConfigStruct struct {
	Name            string
	Url             string
	ClientId        string
	ClientSecret    string
	RedirectUri     string
	AuthEndpoint    string
	ProfileEndpoint string
}

type AuthenticateStruct func(configName string) gin.HandlerFunc
type CallbackStruct func(configName string) gin.HandlerFunc

type UniAuthInstance struct {
	Authenticate AuthenticateStruct
	Callback     CallbackStruct
}
