package cli

import (
	"fmt"

	"github.com/srajalnikhra/cli-login-system/internal/session"
)

func Help() {
	fmt.Println("\n========== Help ==========")

	if session.IsLoggedIn {
		fmt.Println("1. Profile")
		fmt.Println("2. Enable MFA")
		fmt.Println("3. Disable MFA")
		fmt.Println("4. Logout")
		fmt.Println("5. Exit")
	} else {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
	}

	fmt.Println("==========================")
}
