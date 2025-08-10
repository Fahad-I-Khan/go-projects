package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go-oauth2-gin/config"
	"go-oauth2-gin/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func InitOAuthConfig() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  config.EnvGoogleRedirectURL,
		ClientID:     config.EnvGoogleClientID,
		ClientSecret: config.EnvGoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

// GoogleLogin godoc
// @Summary Login with Google
// @Tags Auth
// @Success 302 {string} string "redirect to Google"
// @Router /api/v1/auth/login [get]
func GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("randomstate", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback godoc
// @Summary Callback from Google OAuth
// @Tags Auth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/callback [get]
func GoogleCallback(c *gin.Context) {
	// 1. Get the code from Google
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found in URL"})
		return
	}

	// 2. Exchange code for access token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed"})
		return
	}

	// 3. Use token to fetch user info
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	var userInfo struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.Unmarshal(data, &userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	// 4. Save or update user in DB
	var user models.User
	result := config.DB.Where("email = ?", userInfo.Email).First(&user)

	if result.Error != nil {
		// New user, create
		user = models.User{
			Email:     userInfo.Email,
			FullName:  userInfo.Name,
			AvatarURL: userInfo.Picture,
			Provider:  "google",
		}
		config.DB.Create(&user)
	}

	// 5. Save user information to Redis session
	// The `sessions.Default(c)` gets the session from the context
	session := sessions.Default(c)
	session.Set("user_id", user.ID) // Store the user ID in the session
	session.Save()                  // Save the session to Redis

	// 6. Respond to client
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No session"})
			return
		}

		// Fetch the user from the DB and attach it to the context
		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User not found"})
			return
		}
		c.Set("currentUser", user)

		c.Next()
	}
}

func Dashboard(c *gin.Context) {
	// The user object is now available in the context, thanks to the AuthRequired middleware.
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not found in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the dashboard",
		"user":    user,
	})
}

// Used for tutorial How to add middlerware to protected routes

// func CookieAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		cookie, err := c.Request.Cookie("session_token")
// 		if err != nil || cookie.Value == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No session"})
// 			return
// 		}

// 		// Optionally verify if session ID exists in DB or cache
// 		c.Next()
// 	}
// }

// func Dashboard(c *gin.Context) {
//     // ✅ Try to get the session_token cookie
//     sessionToken, err := c.Cookie("session_token")
//     if err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No session"})
//         return
//     }

//     // ✅ Just return the token for now to confirm it works
//     c.JSON(http.StatusOK, gin.H{
//         "message": "Welcome to dashboard",
//         "user":    sessionToken,
//     })
// }
