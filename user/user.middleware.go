package user

import (
	"golang-api/core"
	"golang-api/log"
	"golang-api/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
	*core.Provider
	userService  *UserService
	tokenService *token.TokenService
	logService   *log.LogService
}

func NewUserMiddleware(module *UserModule) *UserMiddleware {
	return &UserMiddleware{
		Provider:     core.NewProvider("UserMiddleware"),
		userService:  module.Get("UserService").(*UserService),
		tokenService: module.Get("TokenService").(*token.TokenService),
		logService:   module.Get("LogService").(*log.LogService),
	}
}

func (um *UserMiddleware) FindOne(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(name)
		user, err := um.userService.FindByID(id)
		tags := []string{"UserMiddleware", "FindOne"}
		if err != nil {
			um.logService.Errorf(tags, "User %s not found: %v", id, err)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func (um *UserMiddleware) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.MustGet("connectedUser").(*User)
		tags := []string{"UserMiddleware", "IsAdmin"}
		if !user.IsRole("admin") {
			um.logService.Errorf(tags, "User %s is not admin", user.ID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (um *UserMiddleware) IsMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.MustGet("user").(*User)
		connectedUser, _ := c.MustGet("connectedUser").(*User)
		tags := []string{"UserMiddleware", "IsMe"}
		if user.ID != connectedUser.ID && !connectedUser.IsRole(ROLE_ADMIN) {
			um.logService.Errorf(tags, "User %s is not allowed to access user %s", connectedUser.ID, user.ID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (um *UserMiddleware) IsLoggedIn(mandatory bool) gin.HandlerFunc {
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

		token, err := um.tokenService.FindOneBy("access_token", accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaim, err := um.tokenService.ParseJWTToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := um.userService.FindByID(userClaim.UserID)
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
