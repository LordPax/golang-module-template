package user

import (
	"golang-api/core"
	"golang-api/query"

	"github.com/LordPax/sockevent"
)

type IUserService interface {
	core.IProvider
	FindAll(q query.QueryFilter) ([]*User, error)
	FindByID(id string) (*User, error)
	FindOneBy(field string, value any) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	IsUserExists(email, username string) bool
	CountStats(ws *sockevent.Websocket) map[string]int
}

type UserService struct {
	*core.Provider
	userModel IUserModel
}

func NewUserService(module *UserModule) *UserService {
	return &UserService{
		Provider:  core.NewProvider("UserService"),
		userModel: module.Get("UserModel").(IUserModel),
	}
}

func (us *UserService) FindAll(q query.QueryFilter) ([]*User, error) {
	return us.userModel.QueryFindAll(q)
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

func (us *UserService) Update(user *User) error {
	return us.userModel.UpdateByID(user.ID, user)
}

func (us *UserService) Delete(id string) error {
	return us.userModel.DeleteByID(id)
}

func (us *UserService) IsUserExists(email, username string) bool {
	emailExists, _ := us.userModel.CountBy("email", email)
	usernameExists, _ := us.userModel.CountBy("username", username)
	return emailExists > 0 || usernameExists > 0
}

func (us *UserService) CountStats(ws *sockevent.Websocket) map[string]int {
	totalUsers, _ := us.userModel.CountAll()

	loggedClients := ws.FilterClient(func(c *sockevent.Client) bool {
		return c.Get("logged").(bool)
	})

	annonClients := ws.FilterClient(func(c *sockevent.Client) bool {
		return !c.Get("logged").(bool)
	})

	return map[string]int{
		"loggedUsers": len(loggedClients),
		"annonUsers":  len(annonClients),
		"totalUsers":  int(totalUsers),
	}
}
