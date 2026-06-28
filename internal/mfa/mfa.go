package mfa

import (
	"fmt"

	"github.com/srajalnikhra/cli-login-system/internal/repository"
	"github.com/srajalnikhra/cli-login-system/internal/session"
)

func EnableMFA() {
	key, err := GenerateSecret(session.CurrentUser.Username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nScan this QR URL using Google Authenticator:")
	fmt.Println(key.URL())

	err = repository.EnableMFA(session.CurrentUser.Username, key.Secret())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	session.CurrentUser.MFAEnabled = true
	session.CurrentUser.TOTPSecret = key.Secret()

	fmt.Println("\n✅ MFA Enabled Successfully")
}