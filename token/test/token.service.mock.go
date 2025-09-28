package token_test

import (
	"golang-api/core"
	"golang-api/token"
)

type TokenServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewTokenServiceMock(module core.IModule) *TokenServiceMock {
	return &TokenServiceMock{
		Provider: core.NewProvider("TokenService"),
		Mockable: core.NewMockable(),
	}
}

func (as *TokenServiceMock) FindByID(id string) (*token.Token, error) {
	as.MethodCalled("FindByID", id)
	return as.CallFunc("FindByID").(*token.Token), nil
}

func (as *TokenServiceMock) FindOneBy(field string, value any) (*token.Token, error) {
	as.MethodCalled("FindOneBy", field, value)
	return as.CallFunc("FindOneBy").(*token.Token), nil
}

func (as *TokenServiceMock) Create(token *token.Token) error {
	as.MethodCalled("Create", token)
	as.CallFunc("Create")
	return nil
}

func (as *TokenServiceMock) Delete(id string) error {
	as.MethodCalled("Delete", id)
	as.CallFunc("Delete")
	return nil
}

func (as *TokenServiceMock) DeleteByUserID(userID string) error {
	as.MethodCalled("DeleteByUserID", userID)
	as.CallFunc("DeleteByUserID")
	return nil
}

func (as *TokenServiceMock) Update(token *token.Token) error {
	as.MethodCalled("Update", token)
	as.CallFunc("Update")
	return nil
}

func (as *TokenServiceMock) ParseJWTToken(tokenString string) (*token.UserClaims, error) {
	as.MethodCalled("ParseJWTToken", tokenString)
	return as.CallFunc("ParseJWTToken").(*token.UserClaims), nil
}
