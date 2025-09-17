package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
)

var DB *sql.DB

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ConnectDb() error {
	password := os.Getenv("PSQL")
	connstr := fmt.Sprintf("postgres://postgres:%s@localhost:5432/sih25_database?sslmode=disable", password)
	var err error

	DB, err = sql.Open("postgres", connstr)
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
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return fmt.Errorf("email already exists")
			}
		}
		return err
	}
	return nil
}

func GetUserByEmail(email string, password string) error {
	var user User
	err := DB.QueryRow("SELECT email,password FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return fmt.Errorf("User Not Authorised")
	}
	return nil
}

func Close() error {
	return DB.Close()
}
