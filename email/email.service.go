package email

import (
	"context"
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"

	brevo "github.com/getbrevo/brevo-go/lib"
)

type EmailService struct {
	*core.Provider
	dotenvService *dotenv.DotenvService
	client        *brevo.APIClient
	ctx           context.Context
}

func NewEmailService(module *EmailModule) *EmailService {
	return &EmailService{
		Provider:      core.NewProvider("EmailService"),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
		ctx:           context.Background(),
	}
}

func (es *EmailService) OnInit() error {
	fmt.Printf("Initializing %s\n", es.GetName())
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
	to := []brevo.SendSmtpEmailTo{{Email: email}}

	body := brevo.SendSmtpEmail{
		HtmlContent: content,
		Subject:     subject,
		Sender: &brevo.SendSmtpEmailSender{
			Name:  sender,
			Email: sender,
		},
		To: to,
		Params: map[string]interface{}{
			"subject": subject,
		},
	}

	response, httpResp, err := es.client.TransactionalEmailsApi.SendTransacEmail(es.ctx, body)
	if err != nil {
		return fmt.Errorf("error sending email: %w, response: %v, http response: %v", err, response, httpResp)
	}

	return nil
}
