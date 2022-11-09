package models

import "mime/multipart"

type TeamReg struct {
	Username  string   `json:"username" validate:"required,min=4,max=12"`
	TeamName  string   `json:"team" validate:"required,min=4,max=20"`
	CoLeader  string   `json:"co_leader" validate:"required,min=4,max=12"`
	Instagram string   `json:"instagram"`
	Discord   string   `json:"discord"`
	Players   []string `json:"players"`
	// Games           *[]string `json:"games" validate:"required"`
	ContentCreators []string `json:"content_creators"`
}

type TeamData struct {
	TeamName  string
	Leader    string
	Co_leader string
	Instagram string
	Youtube   string
	Discord   string
}

type TeamAchievementsAdd struct {
	Content  string                `form:"type" validate:"required"`
	TeamName string                `form:"team" validate:"required,min=4,max=12"`
	Data     *multipart.FileHeader `form:"data" binding:"required"`
}
