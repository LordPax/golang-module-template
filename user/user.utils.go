package user

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

var secureFields = []string{"id", "username", "firstname", "lastname", "email", "profile", "roles", "created_at", "updated_at"}
