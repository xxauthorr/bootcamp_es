package models

type AdminResult struct {
	Admin string      `json:"user"`
	Data  interface{} `json:"data"`
}

type Entities struct {
	Users       string `json:"users"`
	Teams       string `json:"teams"`
	Tournaments string `json:"tournaments"`
	// Scrims        string
	// Recruitments  string
}

type AdminUserData struct {
	UserName   *string `json:"username"`
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	Team       *string `json:"team"`
	User_type  *string `json:"user_type"`
	Popularity *string `json:"popularity"`
	Created_at *string `json:"created_at"`
	Updated_at *string `json:"uploaded_at"`
	Block      *string `json:"block"`
	Avatar     *string `json:"avatar"`
}

type UserDataList struct {
	Data []AdminUserData `json:"users"`
}
type TeamDataList struct {
	Data []AdminTeamData `json:"users"`
}
type TornamentDataList struct {
	Data []AdminTournamentData `json:"users"`
}

type AdminTeamData struct {
	Team_name  *string
	Leader     *string
	Bio        *string
	Instagram  *string
	Discord    *string
	YouTube    *string
	Avatar     *string
	Co_leader  *string
	Created_at *string
}
type AdminTournamentData struct {
	Owner       *string
	Game        *string
	Name        *string
	Prize_pool  *int32
	Slots       *int
	Reg_end     *string
	T_start     *string
	T_end       *string
	Reg_link    *string
	Live_Stream *string
	Discord     *string
	Created_at  *string
}

type UpdateUserType struct {
	User   string `json:"user" validate:"required"`
	Action string `json:"action" validate:"required"`
}
