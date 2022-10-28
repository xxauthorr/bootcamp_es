package models

type Signup struct {
	Phone    *string `json:"phone" validate:"required,min=7,max=14"`
	UserName *string `json:"username" validate:"required,min=6,max=100"`
	Password *string `json:"password" validate:"required,min=6,max=100"`
}
type ForOtp struct {
	Number *string `json:"phone" validate:"required,min=7,max=14"`
	Otp    *string `json:"otp"`
}
