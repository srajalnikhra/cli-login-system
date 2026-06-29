package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf(
		"host=%s port=5432 user=postgres password=password dbname=cli_login_db sslmode=disable",
		host,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	for i := 1; i <= 10; i++ {
		err = db.Ping()
		if err == nil {
			DB = db
			fmt.Println("✅ Connected to PostgreSQL successfully!")
			return nil
		}

		fmt.Printf("Waiting for PostgreSQL... (%d/10)\n", i)
		time.Sleep(2 * time.Second)
	}

	return err
}