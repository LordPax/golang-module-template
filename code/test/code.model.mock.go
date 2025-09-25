package code_test

import (
	"golang-api/code"
	"golang-api/core"
	"golang-api/query"
)

type CodeModelMock struct {
	*core.ModelMock[*code.Code]
}

func NewCodeModelMock(module *code.CodeModule) *CodeModelMock {
	return &CodeModelMock{
		ModelMock: core.NewModelMock[*code.Code]("CodeModel"),
	}
}

func (cm *CodeModelMock) QueryFindAll(q query.QueryFilter) ([]*code.Code, error) {
	cm.MethodCalled("QueryFindAll", q)
	return cm.CallFunc("QueryFindAll").([]*code.Code), nil
}

func (cm *CodeModelMock) FindOneByCodeAndEmail(c, email string) (*code.Code, error) {
	cm.MethodCalled("FindOneByCodeAndEmail", c, email)
	return cm.CallFunc("FindOneByCodeAndEmail").(*code.Code), nil
}

func (cm *CodeModelMock) DeleteExpiredCodes() error {
	cm.MethodCalled("DeleteExpiredCodes")
	cm.CallFunc("DeleteExpiredCodes")
	return nil
}
