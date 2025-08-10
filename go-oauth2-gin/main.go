package main

import (
	"go-oauth2-gin/config"
	_ "go-oauth2-gin/docs"
	"go-oauth2-gin/handlers"
	"go-oauth2-gin/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title OAuth2 API
// @version 1.0
// @description A demo of OAuth2 login with Google using Gin and Swagger
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// 1. Connect to PostgreSQL database
	config.ConnectDB()
	// 2. Connect to Redis
	config.ConnectRedis()
	// 3. Initialize Google OAuth configs
	handlers.InitOAuthConfig()
	config.InitOAuth()

	r := gin.Default()

	// 4. Configure and add the Redis session store middleware
	// The key "secret" is used to sign the session cookies. This should be a strong
	// and secret key in a production environment.
	redisStore, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"), "", "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatal("Failed to create Redis session store:", err)
	}
	r.Use(sessions.Sessions("go_oauth2_gin_session", redisStore))

	// 5. Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Or your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// 6. Register application routes
	routes.RegisterRoutes(r)

	// 7. Start the server
	port := config.GetEnv("PORT", "8080")
	log.Printf("✅ Server starting on port %s...", port)
	r.Run(":" + port)
}
