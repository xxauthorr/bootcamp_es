package database

import (
	"bootcamp_es/models"
	"fmt"
	"log"
)

type Admin struct {
}

func (a Admin) GetUserCount() string {
	var count string
	getStmnt := `SELECT count(id) FROM user_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}
func (a Admin) GetTeamCount() string {
	var count string
	getStmnt := `SELECT count(id) FROM team_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}
func (a Admin) GetTournamentCount() string {
	var count string
	getStmnt := `SELECT count(id) FROM tournament_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}

func (a Admin) GetTournamentCout() string {
	var count string
	getStmnt := `SELECT count(id) FROM tournament_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}
func (a Admin) GetSerachData(req models.Search) (string, interface{}) {
	var count int
	if req.Entity == "user" {
		getcount := `SELECT count(id) from user_data WHERE username = $1;`
		res := Db.QueryRow(getcount, req.Value)
		if res.Scan(&count); count == 0 {
			return "no data", req
		}
		var data models.AdminUserData
		stmnt := `SELECT username,phone,email,team,user_type,popularity,created_at,updated_at,block,avatar FROM user_data WHERE username = $1;`
		rows := Db.QueryRow(stmnt, req.Value)
		if rows.Err() != nil {
			fmt.Println(rows.Err().Error())
			return rows.Err().Error(), req
		}
		err := rows.Scan(&data.UserName, &data.Phone, &data.Email, &data.Team, &data.User_type, &data.Popularity, &data.Created_at, &data.Updated_at, &data.Block, &data.Avatar)
		if err != nil {
			fmt.Println(rows.Err().Error())
			req.Value = rows.Err().Error()
			return rows.Err().Error(), req
		}
		return "", data
	}
	if req.Entity == "team" {
		getcount := `SELECT count(id) from team_data WHERE team_name = $1;`
		res := Db.QueryRow(getcount, req.Value)
		if res.Scan(&count); count == 0 {
			return "no data", req
		}
		var data models.AdminTeamData
		stmnt := `SELECT team_name,leader,instagram,discord,youtube,avatar,co_leader,created_at FROM team_data WHERE team_name = $1;`
		rows := Db.QueryRow(stmnt, req.Value)
		if rows.Err() != nil {
			fmt.Println(rows.Err().Error())
			return rows.Err().Error(), req
		}
		err := rows.Scan(&data.Team_name, &data.Leader, &data.Instagram, &data.Discord, &data.YouTube, &data.Avatar, &data.Co_leader, &data.Created_at)
		if err != nil {
			fmt.Println(rows.Err().Error())
			return rows.Err().Error(), req
		}
		return "", data
	}
	if req.Entity == "tournament" {
		getcount := `SELECT count(id) from tournament_data WHERE tournament_name = $1;`
		res := Db.QueryRow(getcount, req.Value)
		if res.Scan(&count); count == 0 {
			return "no data", req
		}
		var data models.AdminTournamentData
		stmnt := `SELECT owner,game,tournament_name,prize_pool,no_of_slots,registration_ends,t_start,t_end,registration_link,live_stream,discord,created_at FROM tournament_data WHERE tournament_name=$1;`
		rows := Db.QueryRow(stmnt, req.Value)
		if rows.Err() != nil {
			fmt.Println(rows.Err().Error())
			return rows.Err().Error(), req
		}
		err := rows.Scan(&data.Owner, &data.Game, &data.Name, &data.Prize_pool, &data.Slots, &data.Reg_end, &data.T_start, &data.T_end, &data.Reg_link, &data.Live_Stream, &data.Discord, &data.Created_at)
		if err != nil {
			fmt.Println(rows.Err().Error())
			return rows.Err().Error(), req
		}
		return "", data
	}
	return "invalid", req
}

func (a Admin) GetUsersList(page int64) interface{} {
	limit := 10 * page
	offset := limit - 10
	var (
		data models.AdminUserData
		list models.UserDataList
	)
	stmnt := `SELECT username,phone,email,team,user_type,popularity,created_at,updated_at,block,avatar FROM user_data ORDER BY id ASC OFFSET $1 LIMIT 10;`
	rows, err := Db.Query(stmnt, offset)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&data.UserName, &data.Phone, &data.Email, &data.Team, &data.User_type, &data.Popularity, &data.Created_at, &data.Updated_at, &data.Block, &data.Avatar)
		if err != nil {
			fmt.Println(err.Error())
			return data
		}
		list.Data = append(list.Data, data)
	}
	return list
}
func (a Admin) GetTeamList(page int64) interface{} {
	limit := 10 * page
	offset := limit - 10
	var (
		data models.AdminTeamData
		list models.TeamDataList
	)
	stmnt := `SELECT team_name,leader,instagram,discord,youtube,avatar,co_leader,created_at FROM team_data ORDER BY id ASC OFFSET $1 LIMIT 10;`
	rows, err := Db.Query(stmnt, offset)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&data.Team_name, &data.Leader, &data.Instagram, &data.Discord, &data.YouTube, &data.Avatar, &data.Co_leader, &data.Created_at)
		if err != nil {
			fmt.Println(err.Error())
			return data
		}
		list.Data = append(list.Data, data)
	}
	return list
}

func (a Admin) GetTournamentList(page int64) interface{} {
	limit := 10 * page
	offset := limit - 10
	var (
		data models.AdminTournamentData
		list models.TornamentDataList
	)
	stmnt := `SELECT owner,game,tournament_name,prize_pool,no_of_slots,registration_ends,t_start,t_end,registration_link,live_stream,discord,created_at FROM tournament_data ORDER BY id asc OFFSET $1 LIMIT 10;`
	rows, err := Db.Query(stmnt, offset)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&data.Owner, &data.Game, &data.Name, &data.Prize_pool, &data.Slots, &data.Reg_end, &data.T_start, &data.T_end, &data.Reg_link, &data.Live_Stream, &data.Discord, &data.Created_at)
		if err != nil {
			fmt.Println(err.Error())
			return data
		}
		list.Data = append(list.Data, data)
	}
	return list
}

func (a Admin) MakeAdmin(action string, user string) bool {
	var stmnt string
	if action == "true" {
		stmnt = `UPDATE user_data SET user_type = 'ADMIN' WHERE username = $1;`
	} else {
		stmnt = `UPDATE user_data SET user_type = 'USER' WHERE username = $1;`
	}
	_, err := Db.Exec(stmnt, user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (a Admin) Block(action string, user string) bool {
	var stmnt string
	if action == "true" {
		stmnt = `UPDATE user_data SET block = 'true' WHERE username = $1;`
	} else {
		stmnt = `UPDATE user_data SET block = 'false' WHERE username = $1;`
	}
	_, err := Db.Exec(stmnt, user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (a Admin) DeleteTournament(tourney string) bool {
	stmnt := `DELETE FROM tournament_data WHERE tournament_name = $1;`
	_, err := Db.Exec(stmnt, tourney)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (a Admin) DeleteTeam(team string) bool {
	stmnt := `DELETE FROM team_data WHERE team_name = $1;`
	_, err := Db.Exec(stmnt, team)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	stmnt = `UPDATE user_data SET team = null WHERE team = $1;`
	_, err = Db.Exec(stmnt, team)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	stmnt = `DELETE FROM user_notifications WHERE team = $1;`
	_, err = Db.Exec(stmnt, team)
	if err != nil {

		fmt.Println(err.Error())
		return false
	}
	stmnt = `DELETE FROM team_notifications WHERE team = $1;`
	_, err = Db.Exec(stmnt, team)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
