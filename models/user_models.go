package models

import "mime/multipart"

type UserProfileData struct {
	Visit             bool             `json:"visit"`
	Liked             bool             `json:"liked"`
	Id                string           `json:"id"`
	UserName          string           `json:"username"`
	Phone             string           `json:"phone"`
	Popularity        int64            `json:"popularity"`
	Created_at        string           `json:"created_at"`
	Email             *string          `json:"email"`
	Bio               *string          `json:"bio"`
	Team              *string          `json:"team"`
	Crew              *string          `json:"crew"`
	Instagram         *string          `json:"instagram"`
	Discord           *string          `json:"discord"`
	Whatsapp          *string          `json:"whatsapp"`
	Avatar            *string          `json:"avatar"`
	UserNotifications []U_Notification `json:"user_notification"`
	UserAchievements  UserAchievements `json:"user_achievements"`
}

type UserData struct {
	Liked            bool             `json:"liked"`
	UserName         string           `json:"username"`
	Phone            string           `json:"phone"`
	Popularity       int64            `json:"popularity"`
	Created_at       string           `json:"created_at"`
	Email            *string          `json:"email"`
	Bio              *string          `json:"bio"`
	Team             *string          `json:"team"`
	Crew             *string          `json:"crew"`
	Instagram        *string          `json:"instagram"`
	Discord          *string          `json:"discord"`
	Whatsapp         *string          `json:"whatsapp"`
	Avatar           *string          `json:"avatar"`
	UserAchievements UserAchievements `json:"user_achievements"`
}

type AuthResult struct {
	User          string      `json:"user"`
	Data          interface{} `json:"data"`
	Authorization Token       `json:"auth"`
}

type UserPopularityUpdate struct {
	From   string `json:"from" validate:"required"`
	To     string `json:"to" validate:"required"`
	Action string `json:"action" validate:"required"`
}

type U_Notification struct {
	Id   string `json:"id"`
	Time string `json:"time"`
	Team string `json:"team"`
	Role string `json:"role"`
}

type UserAchievements struct {
	Ingame  []string `json:"ingame"`
	Tourney []string `json:"tounament"`
}

type UserBioEdit struct {
	UserName string `form:"username" validate:"required,min=4,max=12"`
	Bio      string `form:"bio" validate:"max=130"`
	Crew     string `form:"crew" validate:"max=20"`
	Role     string `form:"role" validate:"max=20"`
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

type Notification struct {
	Id     string `json:"id" validate:"required"`
	Action string `json:"action" validate:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}
