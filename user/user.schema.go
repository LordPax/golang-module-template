package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Firstname string     `json:"firstname,omitempty" gorm:"type:varchar(100)"`
	Lastname  string     `json:"lastname,omitempty" gorm:"type:varchar(100)"`
	Username  string     `json:"username,omitempty" gorm:"type:varchar(100)"`
	Email     string     `json:"email,omitempty" gorm:"type:varchar(100)"`
	Password  string     `json:"password,omitempty"`
	Profile   string     `json:"profile,omitempty" gorm:"type:varchar(255)"`
	Roles     []string   `json:"roles,omitempty" gorm:"json"`
	Verified  bool       `json:"verified,omitempty" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) IsRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) Secure() {
	u.Password = ""
	u.Verified = false
}
