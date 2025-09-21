package token

import (
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"

	jwt "github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	core.IProvider
	FindByID(id string) (*Token, error)
	FindOneBy(field string, value any) (*Token, error)
	Create(token *Token) error
	Delete(id string) error
	DeleteByUserID(userID string) error
	Update(token *Token) error
	ParseJWTToken(tokenString string) (*UserClaims, error)
}

type TokenService struct {
	*core.Provider
	tokenModel    *TokenModel
	dotenvService *dotenv.DotenvService
}

func NewTokenService(module *TokenModule) *TokenService {
	return &TokenService{
		Provider:      core.NewProvider("TokenService"),
		tokenModel:    module.Get("TokenModel").(*TokenModel),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (as *TokenService) FindByID(id string) (*Token, error) {
	return as.tokenModel.FindByID(id)
}

func (as *TokenService) FindOneBy(field string, value any) (*Token, error) {
	return as.tokenModel.FindOneBy(field, value)
}

func (as *TokenService) Create(token *Token) error {
	return as.tokenModel.Create(token)
}

func (as *TokenService) Delete(id string) error {
	return as.tokenModel.DeleteByID(id)
}

func (as *TokenService) DeleteByUserID(userID string) error {
	return as.tokenModel.DeleteByUserID(userID)
}

func (as *TokenService) Update(token *Token) error {
	return as.tokenModel.UpdateByID(token.ID, token)
}

func (as *TokenService) ParseJWTToken(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}
	jwtKey := []byte(as.dotenvService.Get("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("the token is invalid")
	}

	return claims, nil
}
