package auth_test

import (
	"golang-api/auth"
	"golang-api/code"
	code_test "golang-api/code/test"
	"golang-api/core"
	dotenv_test "golang-api/dotenv/test"
	email_test "golang-api/email/test"
	log_test "golang-api/log/test"
	user_test "golang-api/user/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCode(t *testing.T, expected *code.Code, actual *code.Code) {
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Email, actual.Email)
	assert.IsType(t, expected.Code, actual.Code)
}

func NewAuthModuleTest() *auth.AuthModule {
	module := &auth.AuthModule{
		Module: core.NewModule("AuthTestModule"),
	}

	module.AddProvider(dotenv_test.NewDotenvServiceMock(""))
	module.AddProvider(email_test.NewEmailServiceMock(module))
	module.AddProvider(log_test.NewLogServiceMock(module))
	module.AddProvider(code_test.NewCodeServiceMock(module))
	module.AddProvider(auth.NewAuthService(module))

	return module
}

func TestSetupAuthModule(t *testing.T) {
	module := NewAuthModuleTest()
	dotenvService := module.Get("DotenvService").(*dotenv_test.DotenvServiceMock)
	emailService := module.Get("EmailService").(*email_test.EmailServiceMock)
	logService := module.Get("LogService").(*log_test.LogServiceMock)
	codeService := module.Get("CodeService").(*code_test.CodeServiceMock)
	authService := module.Get("AuthService").(*auth.AuthService)

	assert.NotNil(t, module, "Auth module should be created")
	assert.NotNil(t, dotenvService, "DotenvService should be created")
	assert.NotNil(t, emailService, "EmailService should be created")
	assert.NotNil(t, logService, "LogService should be created")
	assert.NotNil(t, codeService, "CodeService should be created")
	assert.NotNil(t, authService, "AuthService should be created")
}

func TestAuthService_SendWelcomeEmail(t *testing.T) {
	module := NewAuthModuleTest()
	authService := module.Get("AuthService").(*auth.AuthService)
	emailService := module.Get("EmailService").(*email_test.EmailServiceMock)
	dotenvService := module.Get("DotenvService").(*dotenv_test.DotenvServiceMock)

	receiver := "test@example.com"
	name := "User"
	companyName := "TestCompany"
	path := "auth/template/welcome.html"
	subject := "Bienvenue sur " + companyName + " !"
	params := map[string]any{
		"name":    name,
		"company": companyName,
	}

	dotenvService.MockMethod("Get", func(params ...any) any { return companyName })
	emailService.MockMethod("SendHtmlTemplate", nil)

	authService.SendWelcomeEmail(receiver, name)

	dotenvCalled := dotenvService.IsMethodCalled("Get")
	if !assert.Equal(t, true, dotenvCalled, "DotenvService Get method should be called") {
		return
	}
	dotenvParams := dotenvService.IsParamsEqual("Get", "NAME")
	if !assert.Equal(t, true, dotenvParams, "DotenvService Get method should be called with 'NAME' parameter") {
		return
	}
	called := emailService.IsMethodCalled("SendHtmlTemplate")
	if !assert.Equal(t, true, called, "SendHtmlTemplate method should be called") {
		return
	}
	paramsMatch := emailService.IsParamsEqual("SendHtmlTemplate", receiver, path, subject, params)
	if !assert.Equal(t, true, paramsMatch, "SendHtmlTemplate method should be called with correct parameters") {
		return
	}
}

func TestAuthService_SendWelcomeAndVerif(t *testing.T) {
	module := NewAuthModuleTest()
	authService := module.Get("AuthService").(*auth.AuthService)
	codeService := module.Get("CodeService").(*code_test.CodeServiceMock)
	dotenvService := module.Get("DotenvService").(*dotenv_test.DotenvServiceMock)

	companyName := "TestCompany"
	expectedUser := user_test.CreateUser()
	userId := expectedUser.ID
	receiver := expectedUser.Email
	expectedCode := code_test.CreateCode(userId, receiver)
	var createdCode *code.Code

	dotenvService.MockMethod("Get", func(params ...any) any { return companyName })
	codeService.MockMethod("Create", func(params ...any) any {
		createdCode = params[0].(*code.Code)
		return nil
	})
	codeService.MockMethod("SendVerifCodeEmail", nil)

	authService.SendWelcomeAndVerif(expectedUser)

	called := codeService.IsMethodCalled("SendVerifCodeEmail")
	if !assert.Equal(t, true, called, "CodeService SendVerifCodeEmail method should be called") {
		return
	}
	paramsMatch := codeService.IsParamsEqual("SendVerifCodeEmail", receiver, createdCode.Code)
	if !assert.Equal(t, true, paramsMatch, "CodeService SendVerifCodeEmail method should be called with correct parameters") {
		return
	}
	codeCalled := codeService.IsMethodCalled("Create")
	if !assert.Equal(t, true, codeCalled, "CodeService Create method should be called") {
		return
	}

	testCode(t, expectedCode, createdCode)
}
