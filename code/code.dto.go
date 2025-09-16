package code

type VerifyUserDto struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type RequestCodeDto struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordDto struct {
	Email    string `json:"email" validate:"required,email"`
	Code     string `json:"code" validate:"required"`
	Password string `json:"password" validate:"required"`
}
