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

type UserAchievements struct {
	UserName    string `json:"username" validate:"required,min=4,max=12"`
	Achievement string `json:"achievement" binding:"required"`
}

type UserBioEdit struct {
	UserName string                `form:"username" validate:"required,min=4,max=12"`
	Bio      string                `form:"user_bio" validate:"max=130"`
	Crew     string                `form:"crew" validate:"max=20"`
	Role     string                `form:"role" validate:"max=20"`
	Avatar   *multipart.FileHeader `form:"avatar" binding:"required"`
}

type UserSocialEdit struct {
	UserName  string `json:"username" validate:"required,min=4,max=12"`
	Instagram string `json:"social"`
}

type UserAchievementsEdit struct {
	Content     string                `form:"type" validate:"required"`
	UserName    string                `form:"username" validate:"required,min=4,max=12"`
	Achievement *multipart.FileHeader `form:"data" binding:"required"`
}
