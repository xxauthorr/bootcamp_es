package models

import "time"

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
	Popularity string `gorm:"NOT NULL"`
	Created_at time.Time
	Updated_at time.Time
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
	Id     uint64    `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team   string    `gorm:"NOT NULL"`
	Player string    `gorm:"NOT NULL"`
	Role   string    `gorm:"NOT NULL"`
	Time   time.Time `gorm:"NOT NULL"`
}

type User_social struct {
	Id        uint64 `gorm:"NOT NULL;UNIQUE"`
	Instagram string
	Whatsapp  string
	Discord   string
}

type Team_data struct {
	Id        uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team_name string `gorm:"NOT NULL;UNIQUE"`
	Leader    string `gorm:"NOT NULL;UNIQUE"`
	Instagram string
	Discord   string
	Youtube   string
	Co_leader string `gorm:"UNIQUE"`
}

type Team_achievement struct {
	Id      uint64 `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Team_id uint64 `gorm:"NOT NULL"`
	Content string `gorm:"NOT NULL;CHECK:content = 'TOURNAMENT' OR content = 'SCRIMS'"`
	Data    string `gorm:"NOT NULL"`
}

type Team_notification struct{
	Id     uint64    `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Player string    `gorm:"NOT NULL"`
	Role   string    `gorm:"NOT NULL"`
	Time   time.Time `gorm:"NOT NULL"`
}