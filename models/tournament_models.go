package models

type Tournament_registration_data struct {
	Id                int64    `json:"id"`
	User              string   `json:"user" validate:"required"`
	Game              string   `form:"game" validate:"required"`
	Tournament_name   string   `form:"name" validate:"required,min=5"`
	Prize_pool        int16    `form:"prizepool" validate:"required"`
	No_of_slots       int16    `form:"slots" validate:"required"`
	Registration_ends string   `form:"reg_ends" validate:"required"`
	T_start           string   `form:"start" validate:"required"`
	T_end             string   `form:"end" validate:"required"`
	Registration_link string   `form:"reg_link"`
	Live_stream       string   `form:"live" validate:"required"`
	Discord           string   `form:"discord"`
	Banner            string   `json:"banner"`
	Prize_pool_banner string   `json:"pp_banner"`
	Road_map          string   `json:"road_map"`
	Manager           string   `form:"manager"`
	Sponsers          []string `form:"sponsers"`
	Streamers         []string `form:"streamers"`
}

type Tournament_registration_File struct {
	Banner            string `form:"banner"`
	Prize_pool_banner string `form:"prize"`
	Road_map          string `form:"road_map"`
}


type Tournament_fetch_data struct{
	Id                int64    `json:"id"`
	Owner              string   `json:"user" validate:"required"`
	Game              string   `form:"game" validate:"required"`
	Tournament_name   string   `form:"name" validate:"required,min=5"`
	Prize_pool        int16    `form:"prizepool" validate:"required"`
	No_of_slots       int16    `form:"slots" validate:"required"`
	Registration_ends string   `form:"reg_ends" validate:"required"`
	T_start           string   `form:"start" validate:"required"`
	T_end             string   `form:"end" validate:"required"`
	Registration_link string   `form:"reg_link"`
	Live_stream       string   `form:"live" validate:"required"`
	Discord           *string   `form:"discord"`
	Banner            *string   `json:"banner"`
	Prize_pool_banner *string   `json:"pp_banner"`
	Road_map          *string   `json:"road_map"`
	Manager           *string   `form:"manager"`
}