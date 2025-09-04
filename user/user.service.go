package user

import (
	"golang-api/core"

	"gorm.io/gorm"
)

type UserService struct {
	*core.Provider
	userModel *UserModel
	model     *gorm.DB
}

func NewUserService(userModel *UserModel) *UserService {
	return &UserService{
		Provider:  core.NewProvider("UserService"),
		userModel: userModel,
	}
}

func (us *UserService) FindAll() ([]*User, error) {
	return us.userModel.FindAll()
}

func (us *UserService) FindByID(id string) (*User, error) {
	return us.userModel.FindByID(id)
}

func (us *UserService) FindOneBy(field string, value any) (*User, error) {
	return us.userModel.FindOneBy(field, value)
}

func (us *UserService) Create(user *User) error {
	return us.userModel.Create(user)
}
