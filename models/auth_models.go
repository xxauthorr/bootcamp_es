package models

type SignupForm struct {
	Phone    string `json:"phone" validate:"required"`
	UserName string `json:"username" validate:"required,min=4,max=12"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Otp      string `json:"otp" validate:"required,min=6,max=6"`
}

type LoginForm struct {
	UserName string `json:"username" validate:"required,min=4,max=14"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type ForgetPassword struct {
	Username string `json:"username" validate:"required,min=4,max=12"`
	Phone    string `json:"phone" validate:"required,min=7,max=14"`
}
type ChangePassword struct {
	Otp      string `json:"otp" validate:"required,min=6,max=6"`
	Phone    string `json:"phone" validate:"required,min=7,max=14"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type Search struct {
	Entity string `json:"entity" validate:"required"`
	Value  string `json:"value" validate:"required,max=25"`
}

type HomeData struct {
	User        string       `json:"user"`
	Banner      string       `json:"banner"`
	Top_Players []TopPlayers `json:"players"`
	Top_teams   []TopTeams   `json:"teams"`
	Youtube     interface{}  `json:"youtube_data"`
}

type TopPlayers struct {
	Player     string  `json:"name"`
	Avatar     *string `json:"avatar"`
	Team       *string `json:"team"`
	Popularity string  `json:"popularity"`
}

type TopTeams struct {
	Team       string  `json:"team"`
	Avatar     *string `json:"avatar"`
	Leader     string  `json:"leader"`
	Popularity string  `json:"popularity"`
}
