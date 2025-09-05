package user

import (
	"golang-api/core"
)

type UserService struct {
	*core.Provider
	userModel *UserModel
}

func NewUserService(module *UserModule) *UserService {
	return &UserService{
		Provider:  core.NewProvider("UserService"),
		userModel: module.Get("UserModel").(*UserModel),
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
