package token_test

import (
	"golang-api/core"
	"golang-api/token"
)

type TokenModelMock struct {
	*core.ModelMock[*token.Token]
}

func NewTokenModelMock(module core.IModule) *TokenModelMock {
	return &TokenModelMock{
		ModelMock: core.NewModelMock[*token.Token]("TokenModel"),
	}
}

func (um *TokenModelMock) DeleteByUserID(userID string) error {
	um.MethodCalled("DeleteByUserID", userID)
	um.CallFunc("DeleteByUserID")
	return nil
}
