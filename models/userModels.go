package models

import "mime/multipart"

type UserProfileData struct {
	Liked             bool `json:"liked"`
	Id                string
	UserName          string
	Phone             string
	Popularity        string
	Created_at        string
	Email             *string
	Bio               *string
	Team              *string
	Crew              *string
	Instagram         *string
	Discord           *string
	Whatsapp          *string
	Avatar            *string
	UserNotifications []Notification
	UserAchievements  UserAchievements
}

type UserData struct {
	Liked            bool   `json:"liked"`
	UserName         string `json:"username"`
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
	Avatar           *string
	UserAchievements UserAchievements
}

type UnAutResult struct {
	User          UserData
	Authorization Token
}
type AuthResult struct {
	User          UserProfileData
	Authorization Token
}

type UserPopularityUpdate struct {
	From   string `json:"from" validate:"required"`
	To     string `json:"to" validate:"required"`
	Action string `json:"action" validate:"required"`
}

type Notification struct {
	Id   string
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
	Bio      string                `form:"bio" validate:"max=130"`
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

type UserNotification struct {
	Id     string `json:"id" validate:"required"`
	Action string `json:"action" validate:"required"`
}

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}
