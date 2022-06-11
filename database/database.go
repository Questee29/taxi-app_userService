package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "user"
	dbname   = "taxi_db"
)

func New() *sql.DB {
	//connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//ping

	if err := db.Ping(); err != nil {
		fmt.Println("error while connect")
		panic(err)
	}
	fmt.Println("Successfully connected to database!")

	return db
}
