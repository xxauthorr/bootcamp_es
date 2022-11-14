package models

import "mime/multipart"

type TeamReg struct {
	Username  string   `json:"username" validate:"required,min=4,max=12"`
	TeamName  string   `json:"team" validate:"required,min=4,max=20"`
	CoLeader  string   `json:"co_leader" validate:"required,min=4,max=12"`
	Instagram string   `json:"instagram"`
	Discord   string   `json:"discord"`
	Players   []string `json:"members"`
	// Games           *[]string `json:"games" validate:"required"`
	ContentCreators []string `json:"content_creators"`
}

type TeamData struct {
	Visit         bool
	TeamName      string
	Leader        string
	Co_leader     *string
	Bio           *string
	Avatar        *string
	Instagram     *string
	Youtube       *string
	Discord       *string
	Achievements  TeamAchievements
	Notifications []TeamNotification
	Token         Token
}

type TeamAchievements struct {
	Scrims      []string
	Tournaments []string
}

type TeamNotification struct {
	Id      int64
	Player  string
	Request string
	Time    string
}

type TeamAchievementsAdd struct {
	Content  string                `form:"type" validate:"required"`
	TeamName string                `form:"team" validate:"required"`
	Data     *multipart.FileHeader `form:"data" binding:"required"`
}

type TeamAchievementsDel struct {
	TeamName string `json:"username" validate:"required"`
	Data     string `json:"data" validate:"required"`
}

type TeamBioEdit struct {
	TeamName  string                `form:"team" validate:"required"`
	Bio       string                `form:"bio"`
	Instagram string                `form:"instagram"`
	Discord   string                `form:"discord"`
	Youtube   string                `form:"youtube"`
	Avatar    *multipart.FileHeader `form:"avatar" binding:"required"`
}
