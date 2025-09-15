package code

type VerifyUserDto struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type RequestCodeDto struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordDto struct {
	Email       string `json:"email" validate:"required,email"`
	Code        string `json:"code" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
