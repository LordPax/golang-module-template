package code

import (
	"golang-api/core"
	"golang-api/email"
	"golang-api/query"
)

type ICodeService interface {
	core.IProvider
	FindAll(q query.QueryFilter) ([]*Code, error)
	FindByID(id string) (*Code, error)
	FindOneBy(field string, value any) (*Code, error)
	FindOneByCodeAndEmail(code, email string) (*Code, error)
	Create(code *Code) error
	Update(code *Code) error
	Delete(id string) error
	DeleteBy(field string, value any) error
	DeleteExpiredCodes() error
	SendVerifCodeEmail(receiver, code string) error
	SendResetCodeEmail(receiver, code string) error
}

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

func (cs *CodeService) SendVerifCodeEmail(receiver, code string) error {
	path := "code/template/verification.html"
	subject := "Vérifier votre adresse e-mail"
	params := map[string]any{"code": code}
	return cs.emailService.SendHtmlTemplate(receiver, path, subject, params)
}

func (cs *CodeService) SendResetCodeEmail(receiver, code string) error {
	path := "code/template/reset.html"
	subject := "Réinitialiser votre mot de passe"
	params := map[string]any{"code": code}
	return cs.emailService.SendHtmlTemplate(receiver, path, subject, params)
}
