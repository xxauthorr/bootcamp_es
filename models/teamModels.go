package models


type TeamReg struct {
	Username        *string   `json:"username" validate:"required,min=4,max=12"`
	TeamName        *string   `json:"team" validate:"required,min=4,max=12"`
	CoLeader        *string   `json:"co_leader" validate:"required,min=4,max=12"`
	Instagram       *string   `json:"instagram"`
	Discord         *string   `json:"discord"`
	Players         *[]string `json:"players"`
	// Games           *[4]string `json:"games" validate:"required"`
	ContentCreators *[]string `json:"content_creators"`
}
