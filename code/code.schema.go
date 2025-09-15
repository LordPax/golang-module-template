package code

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Code struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"type:varchar(100)"`
	Email     string    `json:"email" gorm:"type:varchar(100)"`
	Code      string    `json:"code" gorm:"type:varchar(10)"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCode(userID, email string) *Code {
	return &Code{
		UserID:    userID,
		Email:     email,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}
}

func (v *Code) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New().String()
	return nil
}

func (v *Code) IsExpired() bool {
	return time.Now().After(v.ExpiresAt)
}

// GenerateCode generates a random 5-digit code code.
func (v *Code) GenerateCode() {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	v.Code = strconv.Itoa(10000 + randomGenerator.Intn(90000))
}
