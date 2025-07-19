package main

import (
	"go-oauth2-gin/config"
	"go-oauth2-gin/handlers"
	"go-oauth2-gin/routes"
	 _ "go-oauth2-gin/docs"

	"github.com/gin-gonic/gin"
)

// @title OAuth2 API
// @version 1.0
// @description A demo of OAuth2 login with Google using Gin and Swagger
// @host localhost:8080
// @BasePath /api/v1
func main() {
	config.ConnectDB()
	handlers.InitOAuthConfig()
	config.InitOAuth()
	
	


	r := gin.Default()
	routes.RegisterRoutes(r)

	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)
}
