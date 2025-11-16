package email

import (
	"bytes"
	"context"
	"fmt"
	"golang-api/dotenv"
	"golang-api/log"
	"text/template"

	"github.com/LordPax/godular/core"
	brevo "github.com/getbrevo/brevo-go/lib"
)

type IEmailService interface {
	core.IProvider
	Authenticate() error
	Send(email, subject, content string) error
	SendHtml(email, subject, htmlContent string) error
	SendHtmlTemplate(email, content, subject string, params map[string]any) error
	LoadHtmlTemplate(content string, params map[string]interface{}) string
}

type EmailService struct {
	*core.Provider
	dotenvService dotenv.IDotenvService
	logService    log.ILogService
	client        *brevo.APIClient
	ctx           context.Context
	tags          []string
}

func NewEmailService(module core.IModule) *EmailService {
	return &EmailService{
		Provider:      core.NewProvider("EmailService"),
		dotenvService: module.Get("DotenvService").(dotenv.IDotenvService),
		logService:    module.Get("LogService").(log.ILogService),
		ctx:           context.Background(),
		tags:          []string{"Email"},
	}
}

func (es *EmailService) OnInit() error {
	return es.Authenticate()
}

func (es *EmailService) Authenticate() error {
	fmt.Println("Authenticating with Brevo API")

	apiKey := es.dotenvService.Get("BREVO_API_KEY")
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)
	es.client = brevo.NewAPIClient(cfg)

	return nil
}

// Send sends an email using the Brevo API
func (es *EmailService) Send(email, subject, content string) error {
	sender := es.dotenvService.Get("BREVO_SENDER")
	company := es.dotenvService.Get("NAME")
	to := []brevo.SendSmtpEmailTo{{Email: email}}

	body := brevo.SendSmtpEmail{
		TextContent: content,
		Subject:     subject,
		Sender: &brevo.SendSmtpEmailSender{
			Name:  company,
			Email: sender,
		},
		To: to,
	}

	response, httpResp, err := es.client.TransactionalEmailsApi.SendTransacEmail(es.ctx, body)
	if err != nil {
		return fmt.Errorf("error sending email: %w, response: %v, http response: %v", err, response, httpResp)
	}

	return nil
}

func (es *EmailService) SendHtml(email, subject, htmlContent string) error {
	sender := es.dotenvService.Get("BREVO_SENDER")
	company := es.dotenvService.Get("NAME")
	to := []brevo.SendSmtpEmailTo{{Email: email}}

	body := brevo.SendSmtpEmail{
		HtmlContent: htmlContent,
		Subject:     subject,
		Sender: &brevo.SendSmtpEmailSender{
			Name:  company,
			Email: sender,
		},
		To: to,
	}

	response, httpResp, err := es.client.TransactionalEmailsApi.SendTransacEmail(es.ctx, body)
	if err != nil {
		return fmt.Errorf("error sending email: %w, response: %v, http response: %v", err, response, httpResp)
	}

	return nil
}

func (es *EmailService) SendHtmlTemplate(email, content, subject string, params map[string]any) error {
	body := es.LoadHtmlTemplate(content, params)
	return es.SendHtml(email, subject, body)
}

func (es *EmailService) LoadHtmlTemplate(content string, params map[string]interface{}) string {
	var err error
	tmpl := template.New("email")
	tmpl, err = tmpl.Parse(content)
	if err != nil {
		es.logService.Errorf(es.tags, "Error parsing template: %v", err)
		return "<p>Error parsing template</p>"
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, params); err != nil {
		es.logService.Errorf(es.tags, "Error executing template: %v", err)
		return "<p>Error executing template</p>"
	}

	return buffer.String()
}
