package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDb(connstr string) error {
	var err error

	DB, err := sql.Open("postgres", connstr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL database")
	return nil
}

func AddUser(email string, password string) error {
	_, err := DB.Exec("INSERT INTO users (email,password) VALUES ($1,$2)", email, password)
	return err
}

func Close() error {
	return DB.Close()
}
