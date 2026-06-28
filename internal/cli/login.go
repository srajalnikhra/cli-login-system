package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/srajalnikhra/cli-login-system/internal/auth"
	"github.com/srajalnikhra/cli-login-system/internal/mfa"
	"github.com/srajalnikhra/cli-login-system/internal/repository"
	"github.com/srajalnikhra/cli-login-system/internal/session"
)

func Login() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Login =====")

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user, err := repository.FindUserByUsername(username)
	if err != nil {
		fmt.Println("\n❌ Invalid username or password")
		return
	}


	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		fmt.Println("\n❌ Invalid username or password")
		return
	}

	if user.MFAEnabled {
		fmt.Print("Enter OTP: ")

		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)

		if !mfa.VerifyOTP(user.TOTPSecret, code) {
			fmt.Println("\n❌ Invalid OTP")
			return
		}
	}

	session.Login(user)
	fmt.Println("\n✅ Login Successful")
}
