package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Firstname string     `json:"firstname" gorm:"type:varchar(100)"`
	Lastname  string     `json:"lastname" gorm:"type:varchar(100)"`
	Username  string     `json:"username" gorm:"type:varchar(100)"`
	Email     string     `json:"email" gorm:"type:varchar(100)"`
	Password  string     `json:"password"`
	Profile   string     `json:"profile" gorm:"type:varchar(255)"`
	Roles     []string   `json:"roles" gorm:"json"`
	Verified  bool       `json:"verified" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	model     *gorm.DB   `json:"-" gorm:"-"`
}

func (u *User) Save() error {
	return u.model.Save(u).Error
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
