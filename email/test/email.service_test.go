package email_test

import (
	dotenv_test "golang-api/dotenv/test"
	"golang-api/email"
	log_test "golang-api/log/test"
	"testing"

	"github.com/LordPax/godular/core"
	"github.com/stretchr/testify/assert"
)

func NewEmailModuleTest() *email.EmailModule {
	module := &email.EmailModule{
		Module: core.NewModule("EmailTestModule"),
	}

	module.AddProvider(dotenv_test.NewDotenvServiceMock(module))
	module.AddProvider(log_test.NewLogServiceMock(module))
	module.AddProvider(email.NewEmailService(module))

	return module
}

func TestSetupEmailModule(t *testing.T) {
	module := NewEmailModuleTest()
	dotenvService := module.Get("DotenvService").(*dotenv_test.DotenvServiceMock)
	logService := module.Get("LogService").(*log_test.LogServiceMock)
	emailService := module.Get("EmailService").(*email.EmailService)

	assert.NotNil(t, module, "Email module should be created")
	assert.NotNil(t, dotenvService, "DotenvService should be created")
	assert.NotNil(t, logService, "LogService should be created")
	assert.NotNil(t, emailService, "EmailService should be created")
}
