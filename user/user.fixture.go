package user

import (
	"fmt"
	"math/rand"

	"github.com/LordPax/godular/core"
	"github.com/jaswdr/faker/v2"
)

type UserFixture struct {
	*core.Provider
	userService IUserService
	userModel   IUserModel

	fake         faker.Faker
	userNb       int
	userRoles    []string
	userPassword string
}

func NewUserFixture(module core.IModule) *UserFixture {
	service := &UserFixture{
		Provider:    core.NewProvider("UserFixture"),
		userService: module.Get("UserService").(IUserService),
		userModel:   module.Get("UserModel").(IUserModel),

		fake:         faker.New(),
		userNb:       10,
		userRoles:    []string{ROLE_USER, ROLE_ADMIN},
		userPassword: "password",
	}

	module.On("db:fixtures", service.LoadUsers)

	return service
}

func (uf *UserFixture) LoadUsers() error {
	if err := uf.userModel.ClearTable(); err != nil {
		return err
	}

	for i := 0; i < uf.userNb; i++ {
		role := uf.userRoles[rand.Intn(len(uf.userRoles))]
		firstname := uf.fake.Person().FirstName()
		lastname := uf.fake.Person().LastName()

		user := User{
			Firstname: firstname,
			Lastname:  lastname,
			Username:  uf.fake.Person().FirstName(),
			Email:     uf.fake.Internet().Email(),
			Profile:   fmt.Sprintf("https://api.dicebear.com/9.x/initials/svg?seed=%s%s", firstname, lastname),
			Roles:     []string{role},
			Verified:  true,
		}

		if err := user.HashPassword(uf.userPassword); err != nil {
			return err
		}

		if err := uf.userService.Create(&user); err != nil {
			return err
		}

		fmt.Printf("User %s created with password %s\n", user.Username, uf.userPassword)
	}

	return nil
}
