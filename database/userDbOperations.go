package database

import (
	model "bootcamp_es/models"
	bycrypt "bootcamp_es/services/byCrypt"
	"fmt"
	"log"
	"time"
)

type User struct {
	helper bycrypt.ByCrypt
}

func (u User) InsertUser(user model.SignupForm) error {

	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Password = u.helper.HashPassword(user.Password)

	insertStm := `INSERT INTO user_data (username,phone,password,user_type,popularity,created_at,updated_at,block) VALUES ($1,$2,$3,'USER','0',$4,$5,'false');`
	_, err := Db.Exec(insertStm, user.UserName, user.Phone, user.Password, created_at, updated_at)
	if err != nil {
		return err
	}
	return nil
}

func (u User) ChangePass(data, password string) error {
	password = u.helper.HashPassword(password)
	if string(data[0]) == "+" {
		updateStmt := `UPDATE user_data SET password = $1 WHERE phone = $2;`
		_, err := Db.Exec(updateStmt, password, data)
		if err != nil {
			return err
		}
		return nil
	}
	updateStmt := `UPDATE user_data SET password = $1 WHERE username = $2;`
	_, err := Db.Exec(updateStmt, password, data)
	if err != nil {
		return err
	}
	return nil
}

func (u User) UpdateBio(data model.UserBioEdit, avatar string) error {
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateStmnt := "UPDATE user_data SET bio = $1,crew = $2,role = $3,avatar = $4,updated_at = $5 WHERE username = $6;"
	_, err := Db.Exec(updateStmnt, data.Bio, data.Crew, data.Role, avatar, updated_at, data.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (u User) UserProfileData(user string) (bool, model.UserProfileData) {
	var (
		id                                                   string
		email, bio, team, crew, instagram, discord, whatsapp *string
	)
	var userData = model.UserProfileData{}
	getBioData := `SELECT id,username,phone,email,bio,team,crew,popularity,created_at FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getBioData, user)
	if row.Err() != nil {
		log.Panic(row.Err())
		return false, userData
	}
	row.Scan(&id, &userData.UserName, &userData.Phone, &email, &bio, &team, &crew, &userData.Popularity, &userData.Created_at)
	// null check
	null := ""
	if email == nil {
		email = &null
	}
	if bio == nil {
		bio = &null
	}
	if team == nil {
		team = &null
	}
	if crew == nil {
		crew = &null
	}
	userData = model.UserProfileData{
		Email: *email,
		Bio:   *bio,
		Team:  *team,
		Crew:  *crew,
	}
	res, err := Db.Exec(`SELECT * FROM user_social WHERE id = $1;`, id)
	if err != nil {
		log.Panic(err)
		return false, userData
	}
	result, _ := res.RowsAffected()
	if result != 0 {
		getSocialData := `SELECT instagram,whatsapp,discord FROM user_social WHERE Id = $1;`
		row = Db.QueryRow(getSocialData, id)
		if row.Err() != nil {
			log.Panic(row.Err())
			return false, userData
		}
		row.Scan(&instagram, &whatsapp, &discord)
		//null check
		null := ""
		if instagram == nil {
			instagram = &null
		}
		if whatsapp == nil {
			whatsapp = &null
		}
		if discord == nil {
			discord = &null
		}
		userData = model.UserProfileData{
			Instagram: *instagram,
			Whatsapp:  *whatsapp,
			Discord:   *discord,
		}
	}

	return true, userData
}

func (u User) InsertAchievements(user, content, location string) bool {
	var id int64
	getStmnt := `SELECT id FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getStmnt, user)
	if err := row.Scan(&id); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return true
	}

	insertStmnt := `INSERT INTO user_achievements(id,type,data) VALUES ($1,$2,$3)`
	_, err := Db.Exec(insertStmnt, id, content, location)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println("error here '")
		fmt.Println(err.Error())
		return false
	}
	transaction.CommitTransaction()
	return true
}
