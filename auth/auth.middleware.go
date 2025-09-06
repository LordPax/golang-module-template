package auth

import (
	"golang-api/core"
	"golang-api/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	*core.Provider
	authService *AuthService
	userService *user.UserService
}

func NewAuthMiddleware(module *AuthModule) *AuthMiddleware {
	return &AuthMiddleware{
		Provider:    core.NewProvider("AuthMiddleware"),
		authService: module.Get("AuthService").(*AuthService),
		userService: module.Get("UserService").(*user.UserService),
	}
}

func (am *AuthMiddleware) IsLoggedIn(mandatory bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if accessToken == "" || err != nil {
			accessToken = c.GetHeader("Authorization")
		}
		if accessToken == "" {
			accessToken = c.Query("token")
		}

		if !mandatory && accessToken == "" {
			c.Next()
			return
		}

		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if strings.Split(accessToken, " ")[0] == "Bearer" {
			accessToken = strings.Split(accessToken, " ")[1]
		}

		token, err := am.authService.FindOneBy("access_token", accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaim, err := am.authService.ParseJWTToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := am.userService.FindByID(userClaim.UserID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
			c.Abort()
			return
		}

		c.Set("connectedUser", user)
		c.Set("token", token)
		c.Next()
	}
}

func (am *AuthMiddleware) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.MustGet("connectedUser").(*user.User)

		if !user.IsRole("admin") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
