package user_test

import (
	"fmt"
	"golang-api/user"

	"github.com/jaswdr/faker/v2"
)

var fake = faker.New()

func CreateUser() *user.User {
	firstname := fake.Person().FirstName()
	lastname := fake.Person().LastName()

	return &user.User{
		ID:        fake.UUID().V4(),
		Firstname: firstname,
		Lastname:  lastname,
		Username:  fake.Person().FirstName(),
		Email:     fake.Internet().Email(),
		Password:  fake.Internet().Password(),
		Profile:   fmt.Sprintf("https://api.dicebear.com/9.x/initials/svg?seed=%s%s", firstname, lastname),
		Roles:     []string{"user"},
		Verified:  true,
	}
}

func CreateManyUsers(n int) []*user.User {
	users := make([]*user.User, n)
	for i := 0; i < n; i++ {
		users[i] = CreateUser()
	}
	return users
}
