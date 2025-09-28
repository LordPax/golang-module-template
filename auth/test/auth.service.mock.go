package auth_test

import (
	"golang-api/core"
	"golang-api/token"
	"golang-api/user"

	"github.com/gin-gonic/gin"
)

type AuthServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewAuthService(module core.IModule) *AuthServiceMock {
	return &AuthServiceMock{
		Provider: core.NewProvider("AuthServiceMock"),
		Mockable: core.NewMockable(),
	}
}

func (as *AuthServiceMock) SetAuthCookies(c *gin.Context, token *token.Token) {
	as.MethodCalled("SetAuthCookies", c, token)
}

func (as *AuthServiceMock) ClearAuthCookies(c *gin.Context) {
	as.MethodCalled("ClearAuthCookies", c)
}

func (as *AuthServiceMock) SendWelcomeEmail(receiver, name string) error {
	as.MethodCalled("SendWelcomeEmail", receiver, name)
	as.CallFunc("SendWelcomeEmail")
	return nil
}

func (as *AuthServiceMock) SendWelcomeAndVerif(user *user.User) error {
	as.MethodCalled("SendWelcomeAndVerif", user)
	as.CallFunc("SendWelcomeAndVerif")
	return nil
}
