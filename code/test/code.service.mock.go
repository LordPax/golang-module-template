package code_test

import (
	"golang-api/code"
	"golang-api/core"
	"golang-api/query"
)

type CodeServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewCodeServiceMock(module *code.CodeModule) *CodeServiceMock {
	return &CodeServiceMock{
		Provider: core.NewProvider("CodeService"),
		Mockable: core.NewMockable(),
	}
}

func (cs *CodeServiceMock) FindAll(q query.QueryFilter) ([]*code.Code, error) {
	cs.MethodCalled("FindAll", q)
	return cs.CallFunc("FindAll").([]*code.Code), nil
}

func (cs *CodeServiceMock) FindByID(id string) (*code.Code, error) {
	cs.MethodCalled("FindByID", id)
	return cs.CallFunc("FindByID").(*code.Code), nil
}

func (cs *CodeServiceMock) FindOneBy(field string, value any) (*code.Code, error) {
	cs.MethodCalled("FindOneBy", field, value)
	return cs.CallFunc("FindOneBy").(*code.Code), nil
}

func (cs *CodeServiceMock) FindOneByCodeAndEmail(c, email string) (*code.Code, error) {
	cs.MethodCalled("FindOneByCodeAndEmail", c, email)
	return cs.CallFunc("FindOneByCodeAndEmail").(*code.Code), nil
}

func (cs *CodeServiceMock) Create(code *code.Code) error {
	cs.MethodCalled("Create", code)
	cs.CallFunc("Create")
	return nil
}

func (cs *CodeServiceMock) Update(code *code.Code) error {
	cs.MethodCalled("Update", code)
	cs.CallFunc("Update")
	return nil
}

func (cs *CodeServiceMock) Delete(id string) error {
	cs.MethodCalled("Delete", id)
	cs.CallFunc("Delete")
	return nil
}

func (cs *CodeServiceMock) DeleteBy(field string, value any) error {
	cs.MethodCalled("DeleteBy", field, value)
	cs.CallFunc("DeleteBy")
	return nil
}

func (cs *CodeServiceMock) DeleteExpiredCodes() error {
	cs.MethodCalled("DeleteExpiredCodes")
	cs.CallFunc("DeleteExpiredCodes")
	return nil
}

func (cs *CodeServiceMock) SendVerifCodeEmail(receiver, code string) error {
	cs.MethodCalled("SendVerifCodeEmail", receiver, code)
	cs.CallFunc("SendVerifCodeEmail")
	return nil
}

func (cs *CodeServiceMock) SendResetCodeEmail(receiver, code string) error {
	cs.MethodCalled("SendResetCodeEmail", receiver, code)
	cs.CallFunc("SendResetCodeEmail")
	return nil
}
