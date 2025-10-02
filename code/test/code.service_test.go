package code_test

import (
	"golang-api/code"
	"golang-api/core"
	email_test "golang-api/email/test"
	"golang-api/query"
	user_test "golang-api/user/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCode(t *testing.T, expected *code.Code, actual *code.Code) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Code, actual.Code)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.ExpiresAt, actual.ExpiresAt)
}

func NewCodeModuleTest() *code.CodeModule {
	module := &code.CodeModule{
		Module: core.NewModule("CodeTestModule"),
	}

	module.AddProvider(email_test.NewEmailServiceMock(module))
	module.AddProvider(NewCodeModelMock(module))
	module.AddProvider(code.NewCodeService(module))

	return module
}

func TestSetupCodeModule(t *testing.T) {
	module := NewCodeModuleTest()
	emailService := module.Get("EmailService").(*email_test.EmailServiceMock)
	codeModel := module.Get("CodeModel").(*CodeModelMock)
	codeService := module.Get("CodeService").(*code.CodeService)

	assert.NotNil(t, module, "Code module should be created")
	assert.NotNil(t, emailService, "EmailServiceMock should be created")
	assert.NotNil(t, codeModel, "CodeModelMock should be created")
	assert.NotNil(t, codeService, "CodeService should be created")
}

func TestCodeService_FindAll(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	q := query.QueryFilter{}
	nbCodes := 3
	expectedCodes := CreateManyCodes(nbCodes)

	codeModel.MockMethod("QueryFindAll", func(params ...any) any { return expectedCodes })

	newCodes, _ := codeService.FindAll(q)

	called := codeModel.IsMethodCalled("QueryFindAll")
	if !assert.Equal(t, true, called, "QueryFindAll method should be called") {
		return
	}

	if !assert.Len(t, newCodes, nbCodes, "Number of codes should be equal to expected") {
		return
	}
	for i := 0; i < nbCodes; i++ {
		testCode(t, expectedCodes[i], newCodes[i])
	}
}

func TestCodeService_FindByID(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	user := user_test.CreateUser()
	expectedCode := CreateCode(user.ID, user.Email)

	codeModel.MockMethod("FindByID", func(params ...any) any {
		return expectedCode
	})

	newCode, _ := codeService.FindByID(expectedCode.ID)

	called := codeModel.IsMethodCalled("FindByID")
	if !assert.Equal(t, true, called, "FindByID method should be called") {
		return
	}
	params := codeModel.IsParamsEqual("FindByID", expectedCode.ID)
	if !assert.Equal(t, true, params, "FindByID parameter should be the code ID") {
		return
	}

	testCode(t, expectedCode, newCode)
}

func TestCodeService_FindOneBy(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	user := user_test.CreateUser()
	expectedCode := CreateCode(user.ID, user.Email)

	codeModel.MockMethod("FindOneBy", func(params ...any) any {
		return expectedCode
	})

	newCode, _ := codeService.FindOneBy("email", expectedCode.Email)

	called := codeModel.IsMethodCalled("FindOneBy")
	if !assert.Equal(t, true, called, "FindOneBy method should be called") {
		return
	}
	params := codeModel.IsParamsEqual("FindOneBy", "email", expectedCode.Email)
	if !assert.Equal(t, true, params, "FindOneBy parameters should be the field and value") {
		return
	}

	testCode(t, expectedCode, newCode)
}

func TestCodeService_FindOneByCodeAndEmail(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	user := user_test.CreateUser()
	expectedCode := CreateCode(user.ID, user.Email)

	codeModel.MockMethod("FindOneByCodeAndEmail", func(params ...any) any {
		return expectedCode
	})

	newCode, _ := codeService.FindOneByCodeAndEmail(expectedCode.Code, expectedCode.Email)

	called := codeModel.IsMethodCalled("FindOneByCodeAndEmail")
	if !assert.Equal(t, true, called, "FindOneByCodeAndEmail method should be called") {
		return
	}
	params := codeModel.IsParamsEqual("FindOneByCodeAndEmail", expectedCode.Code, expectedCode.Email)
	if !assert.Equal(t, true, params, "FindOneByCodeAndEmail parameters should be the code and email") {
		return
	}

	testCode(t, expectedCode, newCode)
}

func TestCodeService_Create(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	user := user_test.CreateUser()
	expectedCode := CreateCode(user.ID, user.Email)
	var createdCode *code.Code

	codeModel.MockMethod("Create", func(params ...any) any {
		createdCode = params[0].(*code.Code)
		return nil
	})

	err := codeService.Create(expectedCode)
	if !assert.Nil(t, err, "Create should not return an error") {
		return
	}

	called := codeModel.IsMethodCalled("Create")
	if !assert.Equal(t, true, called, "Create method should be called") {
		return
	}
	params := codeModel.GetMethodParams("Create")
	if !assert.Len(t, params, 1, "Create should have one parameter") {
		return
	}
	paramCode, ok := params[0].(*code.Code)
	if !assert.Equal(t, true, ok, "Create parameter should be a code") {
		return
	}

	testCode(t, expectedCode, paramCode)
	testCode(t, expectedCode, createdCode)
}

func TestCodeService_Update(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	user := user_test.CreateUser()
	expectedCode := CreateCode(user.ID, user.Email)
	var updatedCode *code.Code

	codeModel.MockMethod("UpdateByID", func(params ...any) any {
		updatedCode = params[1].(*code.Code)
		return nil
	})

	err := codeService.Update(expectedCode)
	if !assert.Nil(t, err, "Update should not return an error") {
		return
	}

	called := codeModel.IsMethodCalled("UpdateByID")
	if !assert.Equal(t, true, called, "UpdateByID method should be called") {
		return
	}
	params := codeModel.GetMethodParams("UpdateByID")
	if !assert.Len(t, params, 2, "UpdateByID should have two parameters") {
		return
	}
	assert.Equal(t, params[0], expectedCode.ID, "First parameter should be the code ID")
	paramCode, ok := params[1].(*code.Code)
	if !assert.Equal(t, true, ok, "Second parameter should be a code") {
		return
	}

	testCode(t, expectedCode, paramCode)
	testCode(t, expectedCode, updatedCode)
}

func TestCodeService_Delete(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	user := user_test.CreateUser()
	expectedCode := CreateCode(user.ID, user.Email)

	codeModel.MockMethod("DeleteByID", nil)

	err := codeService.Delete(expectedCode.ID)
	if !assert.Nil(t, err, "Delete should not return an error") {
		return
	}

	called := codeModel.IsMethodCalled("DeleteByID")
	if !assert.Equal(t, true, called, "DeleteByID method should be called") {
		return
	}
	params := codeModel.IsParamsEqual("DeleteByID", expectedCode.ID)
	if !assert.Equal(t, true, params, "DeleteByID parameter should be the code ID") {
		return
	}
}

func TestCodeService_DeleteBy(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	field := "email"
	value := "test@example.com"

	codeModel.MockMethod("DeleteBy", nil)

	err := codeService.DeleteBy(field, value)
	if !assert.Nil(t, err, "DeleteBy should not return an error") {
		return
	}

	called := codeModel.IsMethodCalled("DeleteBy")
	if !assert.Equal(t, true, called, "DeleteBy method should be called") {
		return
	}
	params := codeModel.IsParamsEqual("DeleteBy", field, value)
	if !assert.Equal(t, true, params, "DeleteBy parameters should be the field and value") {
		return
	}
}

func TestCodeService_DeleteExpiredCodes(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	codeModel := module.Get("CodeModel").(*CodeModelMock)

	codeModel.MockMethod("DeleteExpiredCodes", nil)

	err := codeService.DeleteExpiredCodes()
	if !assert.Nil(t, err, "DeleteExpiredCodes should not return an error") {
		return
	}

	called := codeModel.IsMethodCalled("DeleteExpiredCodes")
	if !assert.Equal(t, true, called, "DeleteExpiredCodes method should be called") {
		return
	}
}

func TestCodeService_SendVerifCodeEmail(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	emailService := module.Get("EmailService").(*email_test.EmailServiceMock)

	receiver := "test@example.com"
	testCode := "123456"
	path := "code/template/verification.html"
	subject := "Vérifier votre adresse e-mail"

	emailService.MockMethod("SendHtmlTemplate", nil)

	err := codeService.SendVerifCodeEmail(receiver, testCode)
	if !assert.Nil(t, err, "SendVerifCodeEmail should not return an error") {
		return
	}

	called := emailService.IsMethodCalled("SendHtmlTemplate")
	if !assert.Equal(t, true, called, "SendHtmlTemplate method should be called") {
		return
	}

	params := emailService.GetMethodParams("SendHtmlTemplate")
	if !assert.Len(t, params, 4, "SendHtmlTemplate should have four parameters") {
		return
	}

	assert.Equal(t, receiver, params[0], "First parameter should be the receiver email")
	assert.Equal(t, path, params[1], "Second parameter should be the template path")
	assert.Equal(t, subject, params[2], "Third parameter should be the email subject")
	assert.Equal(t, map[string]any{"code": testCode}, params[3], "Fourth parameter should be the template params")
}

func TestCodeService_SendResetCodeEmail(t *testing.T) {
	module := NewCodeModuleTest()
	codeService := module.Get("CodeService").(*code.CodeService)
	emailService := module.Get("EmailService").(*email_test.EmailServiceMock)

	receiver := "test@example.com"
	testCode := "123456"
	path := "code/template/reset.html"
	subject := "Réinitialiser votre mot de passe"

	emailService.MockMethod("SendHtmlTemplate", nil)

	err := codeService.SendResetCodeEmail(receiver, testCode)
	if !assert.Nil(t, err, "SendResetCodeEmail should not return an error") {
		return
	}

	called := emailService.IsMethodCalled("SendHtmlTemplate")
	if !assert.Equal(t, true, called, "SendHtmlTemplate method should be called") {
		return
	}

	params := emailService.GetMethodParams("SendHtmlTemplate")
	if !assert.Len(t, params, 4, "SendHtmlTemplate should have four parameters") {
		return
	}

	assert.Equal(t, receiver, params[0], "First parameter should be the receiver email")
	assert.Equal(t, path, params[1], "Second parameter should be the template path")
	assert.Equal(t, subject, params[2], "Third parameter should be the email subject")
	assert.Equal(t, map[string]any{"code": testCode}, params[3], "Fourth parameter should be the template params")
}
