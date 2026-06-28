package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	connStr := "host=localhost port=5432 user=postgres password=password dbname=cli_login_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DB = db

	fmt.Println("✅ Connected to PostgreSQL successfully!")

	return nil
}