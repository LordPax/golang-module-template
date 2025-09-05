package user

import (
	"golang-api/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
	*core.Provider
	userService *UserService
}

func NewUserMiddleware(module *UserModule) *UserMiddleware {
	return &UserMiddleware{
		Provider:    core.NewProvider("UserMiddleware"),
		userService: module.Get("UserService").(*UserService),
	}
}

func (um *UserMiddleware) GetUser(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(name)
		user, err := um.userService.FindByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
