package database

import (
	model "bootcamp_es/models"
	"fmt"
	"log"
)

type Get struct {
}

var transaction DBoperation

func (g Get) GetPhoneNumber(username string) string {
	var phone string
	getStmnt := `SELECT phone FROM user_data WHERE username = $1;`
	rows := Db.QueryRow(getStmnt, username)
	rows.Scan(&phone)
	return phone
}

func (g Get) GetUsername(phone string) string {
	var username string
	getStmnt := `SELECT username FROM user_data WHERE phone = $1;`
	rows := Db.QueryRow(getStmnt, phone)
	rows.Scan(&username)
	return username
}

// returns the new id after inserting dummy data into user_data
func (g Get) GetNewAchievementName(user, content string) string {
	var userId, val string
	transaction.StartTransaction()
	getStmnt := `SELECT id FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getStmnt, user)
	if err := row.Scan(&userId); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}

	insertStmnt := `INSERT INTO user_achievements (type,data,user_id) VALUES ($1,$2,$3);`
	_, err := Db.Exec(insertStmnt, content, "sample", userId)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	getStmnt = `SELECT id FROM user_achievements WHERE data = $1`
	row = Db.QueryRow(getStmnt, "sample")
	err = row.Scan(&val)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	return val
}

func (g Get) UserBio(user string) (
	bool, model.UserProfileData) {
	var userData model.UserProfileData
	getBioData := `SELECT id,username,phone,email,bio,team,crew,popularity,created_at FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getBioData, user)
	if row.Err() != nil {
		transaction.RollBackTransaction()
		fmt.Println(row.Err())
		return false, userData
	}
	err := row.Scan(&userData.Id, &userData.UserName, &userData.Phone, &userData.Email, &userData.Bio, &userData.Team, &userData.Crew, &userData.Popularity, &userData.Created_at)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
	}
	return true, userData
}

func (g Get) UserSocial(userData model.UserProfileData) (bool, model.UserProfileData) {

	res, err := Db.Exec(`SELECT * FROM user_social WHERE id = $1;`, userData.Id)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error(), "social check")
		return false, userData
	}
	result, _ := res.RowsAffected()
	if result != 0 {
		getSocialData := `SELECT instagram,whatsapp,discord FROM user_social WHERE Id = $1;`
		row := Db.QueryRow(getSocialData, userData.Id)
		if row.Err() != nil {
			transaction.RollBackTransaction()
			log.Panic(row.Err())
			return false, userData
		}
		if err := row.Scan(&userData.Instagram, &userData.Whatsapp, &userData.Discord); err != nil {
			transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false, userData
		}
	}
	return true, userData
}

func (g Get) UserAchievements(data model.UserProfileData) (bool, model.UserProfileData) {
	var ingame, tourney *string
	res, err := Db.Exec(`SELECT * FROM user_achievements WHERE user_id = $1 AND type = 'INGAME';`, data.Id)
	if err != nil {
		transaction.RollBackTransaction()
		return false, data
	}
	count, _ := res.RowsAffected()
	if count != 0 {
		getIngame := `SELECT data FROM user_achievements WHERE user_id = $1 AND type = 'INGAME';`
		rows, err := Db.Query(getIngame, data.Id)
		if err != nil {
			transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false, data
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&ingame); err != nil {
				transaction.RollBackTransaction()
				fmt.Println(err.Error())
				return false, data
			}
			data.UserAchievements.Ingame = append(data.UserAchievements.Ingame, *ingame)
		}
	}
	res, err = Db.Exec(`SELECT * FROM user_achievements WHERE user_id = $1 AND type = 'TOURNEY';`, data.Id)
	if err != nil {
		transaction.RollBackTransaction()
		return false, data
	}
	count, _ = res.RowsAffected()
	if count != 0 {
		getIngame := `SELECT data FROM user_achievements WHERE user_id = $1 AND type = 'TOURNEY';`
		rows, err := Db.Query(getIngame, data.Id)
		if err != nil {
			transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false, data
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&tourney); err != nil {
				transaction.RollBackTransaction()
				fmt.Println(err.Error())
				return false, data
			}
			data.UserAchievements.Tourney = append(data.UserAchievements.Tourney, *tourney)
		}
	}
	return true, data
}

func (g Get) UserNotification(data model.UserProfileData) (bool, model.UserProfileData) {
	var notification model.Notification
	getNotification := `SELECT team,role,time FROM user_notification WHERE player = $1;`
	rows, err := Db.Query(getNotification, data.UserName)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return false, data
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&notification.Team, &notification.Role, &notification.Time); err != nil {
			transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false, data
		}
		data.UserNotification = append(data.UserNotification, notification)
	}
	transaction.CommitTransaction()
	return true, data
}
