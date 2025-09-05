package user

import "golang-api/core"

func castUsers(users []core.IEntity) []*User {
	castedUsers := make([]*User, len(users))
	for i, u := range users {
		castedUsers[i] = u.(*User)
	}
	return castedUsers
}
