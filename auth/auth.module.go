package auth

import (
	auth_lang "golang-api/auth/lang"
	"golang-api/code"
	"golang-api/database"
	"golang-api/dotenv"
	"golang-api/email"
	"golang-api/gin"
	"golang-api/lang"
	"golang-api/log"
	"golang-api/token"
	"golang-api/user"

	"github.com/LordPax/godular/core"
)

var module *AuthModule

type AuthModule struct {
	*core.Module
}

func NewAuthModule() *AuthModule {
	module := &AuthModule{
		Module: core.NewModule("AuthModule"),
	}

	module.AddModule(lang.Module())
	module.AddModule(gin.Module())
	module.AddModule(dotenv.Module())
	module.AddModule(email.Module())
	module.AddModule(database.Module())
	module.AddModule(log.Module())
	module.AddModule(token.Module())
	module.AddModule(user.Module())
	module.AddModule(code.Module())
	module.AddProvider(NewAuthService(module))
	module.AddProvider(NewAuthController(module))

	langService := module.Get("LangService").(lang.ILangService)
	langService.AddStrings(auth_lang.EN_US, "en_US", "en_GB", "en_CA")
	langService.AddStrings(auth_lang.FR_FR, "fr_FR", "fr_CA")

	return module
}

func Module() *AuthModule {
	if module == nil {
		module = NewAuthModule()
	}
	return module
}
