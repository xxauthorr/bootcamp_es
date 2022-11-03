package models

import "mime/multipart"

type UserProfileData struct {
	UserName   string
	Phone      string
	Email      string
	Bio        string
	Team       string
	Crew       string
	Popularity string
	Created_at string
	Instagram  string
	Discord    string
	Whatsapp   string
}

type UserProfileEdit struct {
	UserName string                `json:"username" validate:"required,min=4,max=12"`
	Bio      string                `json:"user_bio" validate:"max=130"`
	Crew     string                `json:"crew" validate:"max=20"`
	Role     string                `json:"role" validate:"max=20"`
	Avatar   *multipart.FileHeader `json:"avatar"`
}

type UserAchievements struct {
	UserName    string                `json:"username" validate:"required,min=4,max=12"`
	Achievement *multipart.FileHeader `json:"achievement" binding:"required"`
}

type UserSocialEdit struct {
	UserName  string `json:"username" validate:"required,min=4,max=12"`
	Instagram string `json:"social"`
}
