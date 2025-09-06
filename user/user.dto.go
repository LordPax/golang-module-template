package user

type CreateUserDto struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type UpdateUserDto struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// type SanitizedUser struct {
// 	ID        string    `json:"id"`
// 	Username  string    `json:"username"`
// 	Firstname string    `json:"firstname"`
// 	Lastname  string    `json:"lastname"`
// 	Email     string    `json:"email"`
// 	Roles     []string  `json:"roles"`
// 	Verified  bool      `json:"verified"`
// 	Profile   string    `json:"profile"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }
