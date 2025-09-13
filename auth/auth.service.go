package auth

import (
	"golang-api/core"
	"golang-api/dotenv"
	"golang-api/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	*core.Provider
	dotenvService *dotenv.DotenvService
}

func NewAuthService(module *AuthModule) *AuthService {
	return &AuthService{
		Provider:      core.NewProvider("AuthService"),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (as *AuthService) SetAuthCookies(c *gin.Context, token *token.Token) {
	cookieSecure, _ := strconv.ParseBool(as.dotenvService.Get("COOKIE_SECURE"))
	c.SetCookie("access_token", token.AccessToken, ACCESS_TOKEN_TTL, "/", "", cookieSecure, true)
	c.SetCookie("refresh_token", token.RefreshToken, REFRESH_TOKEN_TTL, "/", "", cookieSecure, true)
}

func (as *AuthService) ClearAuthCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}
