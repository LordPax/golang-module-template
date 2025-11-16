package token_test

import (
	"golang-api/token"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type TokenModelMock struct {
	*common.ModelMock[*token.Token]
}

func NewTokenModelMock(module core.IModule) *TokenModelMock {
	return &TokenModelMock{
		ModelMock: common.NewModelMock[*token.Token]("TokenModel"),
	}
}

func (um *TokenModelMock) DeleteByUserID(userID string) error {
	um.MethodCalled("DeleteByUserID", userID)
	um.CallFunc("DeleteByUserID")
	return nil
}
