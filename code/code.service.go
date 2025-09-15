package code

import (
	"golang-api/core"
	"golang-api/email"
	"golang-api/query"
)

type CodeService struct {
	*core.Provider
	codeModel    *CodeModel
	emailService *email.EmailService
}

func NewCodeService(module *CodeModule) *CodeService {
	return &CodeService{
		Provider:     core.NewProvider("CodeService"),
		codeModel:    module.Get("CodeModel").(*CodeModel),
		emailService: module.Get("EmailService").(*email.EmailService),
	}
}

func (cs *CodeService) FindAll(q query.QueryFilter) ([]*Code, error) {
	return cs.codeModel.QueryFindAll(q)
}

func (cs *CodeService) FindByID(id string) (*Code, error) {
	return cs.codeModel.FindByID(id)
}

func (cs *CodeService) FindOneBy(field string, value any) (*Code, error) {
	return cs.codeModel.FindOneBy(field, value)
}

func (cs *CodeService) FindOneByCodeAndEmail(code, email string) (*Code, error) {
	return cs.codeModel.FindOneByCodeAndEmail(code, email)
}

func (cs *CodeService) Create(code *Code) error {
	return cs.codeModel.Create(code)
}

func (cs *CodeService) Update(code *Code) error {
	return cs.codeModel.UpdateByID(code.ID, code)
}

func (cs *CodeService) Delete(id string) error {
	return cs.codeModel.DeleteByID(id)
}

func (cs *CodeService) DeleteBy(field string, value any) error {
	return cs.codeModel.DeleteBy(field, value)
}

func (cs *CodeService) DeleteExpiredCodes() error {
	return cs.codeModel.DeleteExpiredCodes()
}

func (cs *CodeService) SendCodeEmail(receiver, code string) error {
	template := email.EmailTemplate{
		Subject: "Vérifier votre adresse e-mail",
		Path:    "code/template/verification.html",
	}
	params := map[string]any{"code": code}
	return cs.emailService.SendHtmlTemplate(receiver, template, params)
}
