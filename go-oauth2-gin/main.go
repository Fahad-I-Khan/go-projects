package main

import (
	"go-oauth2-gin/config"
	_ "go-oauth2-gin/docs"
	"go-oauth2-gin/handlers"
	"go-oauth2-gin/routes"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:8080"}, // or whatever frontend is
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Content-Type"},
    AllowCredentials: true, // <- REQUIRED
	}))
	routes.RegisterRoutes(r)

	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)
}
