package models

type ForOtp struct {
	Number *string `json:"phone" validate:"required,min=7,max=14"`
	Otp    string  `json:"otp" validate:"required,min=6,max=6"`
}

type SignupForm struct {
	Phone    string `json:"phone" validate:"required,min=7,max=14"`
	UserName string `json:"username" validate:"required,min=4,max=12"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginForm struct {
	UserName string `json:"username" validate:"required,min=4,max=14"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type ForgetPassword struct {
	Username *string `json:"username" validate:"required,min=4,max=12"`
	Phone    *string `json:"phone" validate:"required,min=7,max=14"`
	Password string  `json:"password" validate:"required,min=6,max=100"`
}
