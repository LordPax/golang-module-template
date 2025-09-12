package token

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	UserID       string    `json:"user_id" gorm:"index"`
	AccessToken  string    `json:"access_token" gorm:"type:varchar(255)"`
	RefreshToken string    `json:"refresh_token" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	return nil
}

func (t *Token) GenerateAccessToken(jwtKey string) error {
	now := time.Now()

	claims := &UserClaims{
		UserID: t.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		return fmt.Errorf("failed to sign the token: %w", err)
	}

	t.AccessToken = accessTokenString

	return nil
}

func (t *Token) GenerateRefreshToken(jwtKey string) error {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken.Claims = jwt.MapClaims{
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
		"sub": 1,
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(jwtKey))
	if err != nil {
		return fmt.Errorf("failed to sign the token: %w", err)
	}

	t.RefreshToken = refreshTokenString

	return nil
}

func (t *Token) GenerateTokens(jwtKey string) error {
	if err := t.GenerateAccessToken(jwtKey); err != nil {
		return fmt.Errorf("failed to generate access token: %w", err)
	}

	if err := t.GenerateRefreshToken(jwtKey); err != nil {
		return fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return nil
}
