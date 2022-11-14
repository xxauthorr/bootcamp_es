package models

type Tournament_registration_data struct {
	Id                int64    `json:"id"`
	User              string   `json:"user" validate:"required"`
	Game              string   `form:"game" validate:"required"`
	Tournament_name   string   `form:"name" validate:"required,min=5"`
	Prize_pool        string    `form:"pizepool"`
	No_of_slots       string    `form:"slots" validate:"required"`
	Registration_ends string   `form:"reg_ends" validate:"required"`
	T_start           string   `form:"start" validate:"required"`
	T_end             string   `form:"end" validate:"required"`
	Registration_link string   `form:"reg_link"`
	Live_stream       string     `form:"live"`
	Discord           string   `form:"discord"`
	Banner            string   `form:"banner"`
	Prize_pool_banner string   `form:"pp_banner"`
	Road_map          string   `form:"road_map"`
	Manager           string   `form:"manager"`
	Sponsers          []string `form:"sponsers"`
	Streamers         []string `form:"streamers"`
}

type Tournament_registration_File struct {
	Banner            string `form:"banner"`
	Prize_pool_banner string `form:"prize"`
	Road_map          string `form:"road_map"`
}
