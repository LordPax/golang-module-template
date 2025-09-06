package auth

import (
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"

	jwt "github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	*core.Provider
	tokenModel    *TokenModel
	dotenvService *dotenv.DotenvService
}

func NewAuthService(module *AuthModule) *AuthService {
	return &AuthService{
		Provider:      core.NewProvider("AuthService"),
		tokenModel:    module.Get("TokenModel").(*TokenModel),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (as *AuthService) FindAll() ([]*Token, error) {
	return as.tokenModel.FindAll()
}

func (as *AuthService) FindByID(id string) (*Token, error) {
	return as.tokenModel.FindByID(id)
}

func (as *AuthService) FindOneBy(field string, value any) (*Token, error) {
	return as.tokenModel.FindOneBy(field, value)
}

func (as *AuthService) CreateToken(token *Token) error {
	return as.tokenModel.Create(token)
}

func (as *AuthService) DeleteTokensByUserID(userID string) error {
	return as.tokenModel.DeleteTokensByUserID(userID)
}

func (as *AuthService) ParseJWTToken(tokenString string) (*UserClaims, error) {
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
