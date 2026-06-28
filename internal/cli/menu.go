package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/srajalnikhra/cli-login-system/internal/session"
)

var reader = bufio.NewReader(os.Stdin)

func ShowMenu() {
	for {

		if session.IsLoggedIn {
			fmt.Println("\n========== CLI Login System ==========")
			fmt.Println("1. Profile")
			fmt.Println("2. Enable MFA")
			fmt.Println("3. Logout")
			fmt.Println("4. Exit")
		} else {
			fmt.Println("\n========== CLI Login System ==========")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("3. Exit")
		}

		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if session.IsLoggedIn {
			switch choice {
			case "1":
				Profile()
			case "2":
				EnableMFA()
			case "3":
				session.Logout()
				fmt.Println("\n✅ Logged Out")
			case "4":
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("Invalid choice")
			}
		} else {
			switch choice {
			case "1":
				Register()
			case "2":
				Login()
			case "3":
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("Invalid choice")
			}
		}
	}
}
