package main

import (
	"log"

	"github.com/srajalnikhra/cli-login-system/internal/cli"
	"github.com/srajalnikhra/cli-login-system/internal/database"
)

func main() {

	// Establish database connection.
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create required database tables if they do not exist.
	err = database.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// Start the CLI application.
	cli.ShowMenu()
}
