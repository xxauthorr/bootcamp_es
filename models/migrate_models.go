package models

type User_data struct {
	Id         uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Username   string `gorm:"NOT NULL;UNIQUE"`
	Phone      string `gorm:"NOT NULL;UNIQUE"`
	Email      string
	Password   string `gorm:"NOT NULL"`
	Bio        string
	Team       string
	User_type  string `gorm:"NOT NULL;CHECK:user_type = 'USER' OR user_type = 'ADMIN'"`
	Crew       string
	Role       string
	Popularity int64 `gorm:"NOT NULL"`
	Created_at string
	Updated_at string
	Block      bool `gorm:"NOT NULL"`
	Avatar     string
}

type User_achievement struct {
	Id      uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	User_id uint64 `gorm:"NOT NULL"`
	Content string `gorm:"NOT NULL;CHECK:content = 'TOURNAMENT' OR content = 'INGAME'"`
	Data    string `gorm:"NOT NULL"`
}

type User_notification struct {
	Id     uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team   string `gorm:"NOT NULL"`
	Player string `gorm:"NOT NULL"`
	Role   string `gorm:"NOT NULL"`
	Time   string `gorm:"NOT NULL"`
}

type User_social struct {
	Id        uint64 `gorm:"NOT NULL;UNIQUE"`
	Instagram string
	Whatsapp  string
	Discord   string
}

type Team_data struct {
	Id         uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team_name  string `gorm:"NOT NULL;UNIQUE"`
	Leader     string `gorm:"NOT NULL;UNIQUE"`
	Bio        string
	Instagram  string
	Discord    string
	Youtube    string
	Avatar     string
	Co_leader  string `gorm:"UNIQUE"`
	Created_at string `gorm:"NOT NULL"`
}

type Team_achievement struct {
	Id      uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team_id uint64 `gorm:"NOT NULL"`
	Content string `gorm:"NOT NULL;CHECK:content = 'TOURNAMENT' OR content = 'SCRIMS'"`
	Data    string `gorm:"NOT NULL"`
}

type Team_notification struct {
	Id      uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team    string `gorm:"NOT NULL"`
	Player  string `gorm:"NOT NULL"`
	Request string `gorm:"NOT NULL"`
	Time    string `gorm:"NOT NULL"`
}

type User_popularity struct {
	Id      uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Provide string `gorm:"NOT NULL"`
	Consume string `gorm:"NOT NULL"`
}

type Tournament_data struct {
	Id                int64  `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Owner             string `gorm:"NOT NULL;UNIQUE"`
	Game              string `gorm:"NOT NULL"`
	Tournament_name   string `gorm:"NOT NULL;UNIQUE"`
	Prize_pool        int32  `gorm:"NOT NULL"`
	No_of_slots       int32  `gorm:"NOT NULL"`
	Registration_ends string `gorm:"NOT NULL"`
	T_start           string `gorm:"NOT NULL"`
	T_end             string `gorm:"NOT NULL"`
	Registration_link string `gorm:"NOT NULL"`
	Live_stream       bool   `gorm:"NOT NULL"`
	Discord           string
	Banner            string
	Prize_pool_banner string
	Road_map          string
	Manager           string
	Sponser           string
	Streamer          string
	Created_at        string `gorm:"NOT NULL"`
}
