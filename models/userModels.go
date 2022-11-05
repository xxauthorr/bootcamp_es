package models

import "mime/multipart"

type UserProfileData struct {
	Id               string
	UserName         string
	Phone            string
	Popularity       string
	Created_at       string
	Email            *string
	Bio              *string
	Team             *string
	Crew             *string
	Instagram        *string
	Discord          *string
	Whatsapp         *string
	UserNotification []Notification
	UserAchievements UserAchievements
}
type Notification struct {
	Time string
	Team string
	Role string
}

type UserAchievements struct {
	Ingame  []string
	Tourney []string
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
	Instagram string `json:"instagram"`
	Whatsapp  string `json:"whatsapp"`
	Discord   string `json:"discord"`
}

type UserAchievementsAdd struct {
	Content  string                `form:"type" validate:"required"`
	UserName string                `form:"username" validate:"required,min=4,max=12"`
	Data     *multipart.FileHeader `form:"data" binding:"required"`
}

type UserAchievementsDel struct {
	UserName string `json:"username" validate:"required,min=4,max=12"`
	Data     string `json:"data" validate:"required"`
}
