package user

import (
	"golang-api/core"
	"golang-api/query"
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

func (us *UserService) FindAll(q query.QueryFilter) ([]*User, error) {
	return us.userModel.QueryFindAll(q)
}

func (us *UserService) FindByID(id string) (*User, error) {
	return us.userModel.FindByID(id, []string{})
}

func (us *UserService) FindOneBy(field string, value any) (*User, error) {
	return us.userModel.FindOneBy(field, value, []string{})
}

func (us *UserService) Create(user *User) error {
	return us.userModel.Create(user)
}

func (us *UserService) Update(user *User) error {
	return us.userModel.UpdateByID(user.ID, user)
}

func (us *UserService) Delete(id string) error {
	return us.userModel.DeleteByID(id)
}

func (us *UserService) Save(user *User) error {
	return us.userModel.Save(user)
}

func (us *UserService) IsUserExists(email, username string) bool {
	emailExists, _ := us.userModel.CountBy("email", email)
	usernameExists, _ := us.userModel.CountBy("username", username)
	return emailExists > 0 || usernameExists > 0
}
