package email

import (
	"bytes"
	"context"
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"
	"golang-api/log"
	"os"
	"text/template"

	brevo "github.com/getbrevo/brevo-go/lib"
)

type EmailService struct {
	*core.Provider
	dotenvService *dotenv.DotenvService
	logService    *log.LogService
	client        *brevo.APIClient
	ctx           context.Context
	tags          []string
}

func NewEmailService(module *EmailModule) *EmailService {
	return &EmailService{
		Provider:      core.NewProvider("EmailService"),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
		logService:    module.Get("LogService").(*log.LogService),
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

func (es *EmailService) SendHtmlTemplate(email string, template EmailTemplate, params map[string]any) error {
	body := es.LoadHtmlTemplate(template.Path, params)
	return es.SendHtml(email, template.Subject, body)
}

func (es *EmailService) LoadHtmlTemplate(filePath string, params map[string]interface{}) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		es.logService.Errorf(es.tags, "Error reading template file: %v", err)
		return "<p>Error loading template</p>"
	}

	tmpl := template.New("email")
	tmpl, err = tmpl.Parse(string(content))
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

// // Send sends an email using the Brevo API
// func (es *EmailService) SendTemplate(email string, template int64, params map[string]any) error {
// 	sender := es.dotenvService.Get("BREVO_SENDER")
// 	company := es.dotenvService.Get("NAME")
// 	to := []brevo.SendSmtpEmailTo{{Email: email}}

// 	body := brevo.SendSmtpEmail{
// 		TemplateId: template,
// 		Sender: &brevo.SendSmtpEmailSender{
// 			Name:  company,
// 			Email: sender,
// 		},
// 		To:     to,
// 		Params: params,
// 	}

// 	response, httpResp, err := es.client.TransactionalEmailsApi.SendTransacEmail(es.ctx, body)
// 	if err != nil {
// 		return fmt.Errorf("error sending email: %w, response: %v, http response: %v", err, response, httpResp)
// 	}

// 	return nil
// }
