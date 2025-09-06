package user

import "golang-api/core"

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

func castUsers(users []core.IEntity) []*User {
	castedUsers := make([]*User, len(users))
	for i, u := range users {
		castedUsers[i] = u.(*User)
	}
	return castedUsers
}
