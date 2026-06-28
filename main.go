package main

import (
	"log"

	"github.com/srajalnikhra/cli-login-system/internal/cli"
	"github.com/srajalnikhra/cli-login-system/internal/database"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	cli.ShowMenu()
}
