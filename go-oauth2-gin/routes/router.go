package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-oauth2-gin/config"    
	"go-oauth2-gin/handlers"  
)
import ginSwagger "github.com/swaggo/gin-swagger"
import swaggerFiles "github.com/swaggo/files"



func RegisterRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		// ✅ Health check with actual DB ping
		v1.GET("/health", func(c *gin.Context) {
			sqlDB, err := config.DB.DB()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Database connection error",
					"error":  err.Error(),
				})
				return
			}

			err = sqlDB.Ping()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Database not reachable",
					"error":  err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "Database connected ✅"})
		})

		// Auth routes using handlers
		auth := v1.Group("/auth")
		{
			auth.GET("/login", handlers.GoogleLogin)
			auth.GET("/callback", handlers.GoogleCallback)
		}
		protected := v1.Group("/protected")
		{
			protected.Use(handlers.CookieAuthMiddleware())
			protected.GET("/dashboard", handlers.Dashboard)
		}
	}
}

