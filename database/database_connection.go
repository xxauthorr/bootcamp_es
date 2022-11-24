package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	host    string
	port    string
	user    string
	pass    string
	dbName  string
	sslMode string
}

var Db *sql.DB

// var gormDb *gorm.DB

func ConnectDb() error {
	// loads env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		return err
	}

	configure := &config{
		host:    os.Getenv("DB_HOST"),
		port:    os.Getenv("DB_PORT"),
		user:    os.Getenv("DB_USER"),
		pass:    os.Getenv("DB_PASS"),
		dbName:  os.Getenv("DB_NAME"),
		sslMode: os.Getenv("DB_SSLMODE"),
	}

	initial_conf := fmt.Sprintf("host= %s port= %s user= %s password= %s dbname= postgres sslmode=%s",
		configure.host,
		configure.port,
		configure.user,
		configure.pass,
		configure.sslMode)

	Db, err = sql.Open(configure.user, initial_conf)
	if err != nil {
		log.Fatal("Error connecting to database - ", err)
		return err
	}
	// tada
	_, err = Db.Exec(`CREATE DATABASE ` + configure.dbName + `;`)
	if err != nil {
		if res := strings.Contains(err.Error(), configure.dbName); !res {
			return err
		}
		return nil
	}

	primary_conf := fmt.Sprintf("host= %s port= %s user= %s password= %s dbname= %s sslmode=%s",
		configure.host,
		configure.port,
		configure.user,
		configure.pass,
		configure.dbName,
		configure.sslMode)
	Db, err = sql.Open(configure.user, primary_conf)
	if err != nil {
		log.Fatal("Error connecting to database - ", err)
		return err
	}
	gormDb, err := gorm.Open(postgres.Open(primary_conf), &gorm.Config{})
	if err != nil {
		log.Fatal("Error executing gorm  - ", err)
		return err
	}
	AutoMigrateTables(gormDb)
	return nil
}
