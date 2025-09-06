package auth

import jwt "github.com/golang-jwt/jwt/v5"

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

type LoginSuccessResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshSuccessResponse struct {
	AccessToken string `json:"access_token"`
}
