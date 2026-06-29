package cli

import (
	"fmt"
	"strings"

	"github.com/srajalnikhra/cli-login-system/internal/auth"
	"github.com/srajalnikhra/cli-login-system/internal/models"
	"github.com/srajalnikhra/cli-login-system/internal/repository"
)

// Register creates a new user account
// and stores it in the database.
func Register() {
	fmt.Println("\n===== Register =====")

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hash, err := auth.HashPassword(password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user := models.User{
		Username:     username,
		PasswordHash: hash,
		MFAEnabled:   false,
		TOTPSecret:   "",
	}

	err = repository.CreateUser(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n✅ Registration Successful")
}
