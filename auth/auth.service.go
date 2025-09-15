package auth

import (
	"fmt"
	codeM "golang-api/code"
	"golang-api/core"
	"golang-api/dotenv"
	"golang-api/email"
	"golang-api/log"
	"golang-api/token"
	"golang-api/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	*core.Provider
	dotenvService *dotenv.DotenvService
	emailService  *email.EmailService
	logService    *log.LogService
	codeService   *codeM.CodeService
}

func NewAuthService(module *AuthModule) *AuthService {
	return &AuthService{
		Provider:      core.NewProvider("AuthService"),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
		emailService:  module.Get("EmailService").(*email.EmailService),
		logService:    module.Get("LogService").(*log.LogService),
		codeService:   module.Get("CodeService").(*codeM.CodeService),
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

func (as *AuthService) SendWelcomeEmail(receiver, name string) error {
	company := as.dotenvService.Get("NAME")
	template := email.EmailTemplate{
		Subject: "Bienvenue sur " + company + " !",
		Path:    "auth/template/welcome.html",
	}
	params := map[string]any{
		"name":    name,
		"company": company,
	}
	return as.emailService.SendHtmlTemplate(receiver, template, params)
}

// SendWelcomeAndVerif sends welcome and verification emails to the user.
func (as *AuthService) SendWelcomeAndVerif(user *user.User) error {
	tags := []string{"AuthService", "SendWelcomeAndVerif"}

	code := codeM.NewCode(user.ID, user.Email)
	code.GenerateCode()

	if err := as.codeService.Create(code); err != nil {
		as.logService.Errorf(tags, "Failed to create code for user %s: %v", user.Email, err)
		return fmt.Errorf("Failed to create code")
	}

	if err := as.codeService.SendCodeEmail(user.Email, code.Code); err != nil {
		as.logService.Errorf(tags, "Failed to send code email to user %s: %v", user.Email, err)
		return fmt.Errorf("Failed to send code email")
	}

	if err := as.SendWelcomeEmail(user.Email, user.Firstname); err != nil {
		as.logService.Errorf(tags, "Failed to send welcome email to user %s: %v", user.Email, err)
		return fmt.Errorf("Failed to send welcome email")
	}

	return nil
}
