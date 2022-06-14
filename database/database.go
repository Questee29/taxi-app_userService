package database

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/Questee29/taxi-app_userService/configs"
	"github.com/pressly/goose"
)

func New() (*sql.DB, error) {
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	//connectionInfo
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName)

	//open connection
	db, err := sql.Open(config.Database.DbDriver, psqlInfo)
	if err != nil {
		return nil, err
	}

	//make migrations

	if err := Migrate(db); err != nil {
		return nil, nil
	}

	//ping

	if err := db.Ping(); err != nil {
		log.Fatalln("error while connect")
		return nil, err
	}
	log.Println("Successfully connected to database!")

	return db, nil
}
func Migrate(db *sql.DB) error {
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	if config.Database.DBReload {
		log.Println("Start reloading database")
		err := goose.DownTo(db, ".", 0)
		if err != nil {
			return err
		}
	}

	log.Println("start migrating database")
	err = goose.Up(db, "migrations")
	if err != nil {
		return err
	}
	return nil
}
