package database

import (
	helper "bootcamp_es/helpers"
	model "bootcamp_es/models"
	"time"
)

type User struct{}


func (u User) RegisterUser(user model.Signup) error {

	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	*user.Password = helper.HashPassword(*user.Password)

	insertStm := `INSERT INTO user_data (username,phone,password,user_type,popularity,created_at,updated_at,block) VALUES ($1,$2,$3,'USER','0',$4,$5,'false');`
	_, err := Db.Exec(insertStm, *user.UserName, *user.Phone, *user.Password, created_at, updated_at)
	if err != nil {
		return err
	}
	return nil
}
