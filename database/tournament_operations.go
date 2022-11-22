package database

import (
	"github.com/xxauthorr/bootcamp_es/models"
	"fmt"
	"time"
)

type Tournament struct {
}

func (db Tournament) RegisterTournament(data models.Tournament_registration_data) (bool, int64) {
	var id int64
	created_at := time.Now().Format("2-01-2006 3:04:05 PM")
	transaction.StartTransaction()
	stmnt := `INSERT INTO tournament_data (owner,tournament_name,game,prize_pool,no_of_slots,registration_ends,t_start,t_end,registration_link,live_stream,discord,manager,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id;`

	rows := Db.QueryRow(stmnt, data.User, data.Tournament_name, data.Game, data.Prize_pool, data.No_of_slots, data.Registration_ends, data.T_start, data.T_end, data.Registration_link, data.Live_stream, data.Discord, data.Manager, created_at)

	if rows.Err() != nil {
		fmt.Println(rows.Err().Error())
		return false, 0
	}
	if err := rows.Scan(&id); err != nil {
		fmt.Println("error in returning")
		fmt.Println(err.Error())
		return false, 0
	}
	return true, id
}

func (db Tournament) InsertFile(data models.Tournament_registration_data) bool {
	stmnt := `UPDATE tournament_data SET banner=$1,prize_pool_banner=$2,road_map=$3 WHERE id = $4;`
	_, err := Db.Exec(stmnt, data.Banner, data.Prize_pool_banner, data.Road_map, data.Id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	transaction.CommitTransaction()
	return true
}

func (db Tournament) GetTournamentData(tour string) models.Tournament_fetch_data {
	var data models.Tournament_fetch_data
	stmnt := `SELECT id,owner,game,tournament_name,prize_pool,no_of_slots,registration_ends,t_start,t_end,registration_link,live_stream,discord,banner,prize_pool_banner,road_map FROM tournament_data WHERE tournament_name = $1;`
	row := Db.QueryRow(stmnt, tour)
	if row.Err() != nil {
		fmt.Println(row.Err().Error())
		return data
	}
	err := row.Scan(&data.Id, &data.Owner, &data.Game, &data.Tournament_name, &data.Prize_pool, &data.No_of_slots, &data.Registration_ends, &data.T_start, &data.T_end, &data.Registration_link, &data.Live_stream, &data.Discord, &data.Banner, &data.Prize_pool_banner, &data.Road_map)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	return data
}

func (db Tournament) UpdateFile(fileName, file string, id int64) bool {
	var stmnt string
	if file == "banner" {
		stmnt = `UPDATE tournament_data SET banner = $1 WHERE id = $2;`
	} else if file == "prize" {
		stmnt = `UPDATE tournament_data SET prize_pool_banner = $1 WHERE id = $2;`
	} else {
		stmnt = `UPDATE tournament_data SET road_map = $1 WHERE id = $2;`
	}
	_, err := Db.Exec(stmnt, fileName, id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (db Tournament) UpdateTournament(data models.EditTournamentData) bool {
	stmnt := `UPDATE tournament_data SET no_of_slots = $1,manager= $2,registration_link = $3,registration_ends = $4 WHERE tournament_name = $5;`
	_, err := Db.Exec(stmnt, data.Slots, data.Manager, data.Reg_Link, data.Reg_ends, data.Name)
	if err != nil {
		fmt.Println(err.Error())

		return false
	}
	return true
}
