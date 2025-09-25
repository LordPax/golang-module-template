package email_test

import (
	"golang-api/core"
	"golang-api/email"
)

type EmailServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewEmailServiceMock(module *email.EmailModule) *EmailServiceMock {
	return &EmailServiceMock{
		Provider: core.NewProvider("EmailService"),
		Mockable: core.NewMockable(),
	}
}

func (es *EmailServiceMock) Authenticate() error {
	es.MethodCalled("Authenticate")
	es.CallFunc("Authenticate")
	return nil
}

// Send sends an email using the Brevo API
func (es *EmailServiceMock) Send(email, subject, content string) error {
	es.MethodCalled("Send", email, subject, content)
	es.CallFunc("Send")
	return nil
}

func (es *EmailServiceMock) SendHtml(email, subject, htmlContent string) error {
	es.MethodCalled("SendHtml", email, subject, htmlContent)
	es.CallFunc("SendHtml")
	return nil
}

func (es *EmailServiceMock) SendHtmlTemplate(email, path, subject string, params map[string]any) error {
	es.MethodCalled("SendHtmlTemplate", email, path, subject, params)
	es.CallFunc("SendHtmlTemplate")
	return nil
}

func (es *EmailServiceMock) LoadHtmlTemplate(filePath string, params map[string]interface{}) string {
	es.MethodCalled("LoadHtmlTemplate", filePath, params)
	return es.CallFunc("LoadHtmlTemplate").(string)
}
