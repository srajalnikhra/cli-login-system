package cli

import (
	"fmt"

	"github.com/srajalnikhra/cli-login-system/internal/mfa"
	"github.com/srajalnikhra/cli-login-system/internal/repository"
	"github.com/srajalnikhra/cli-login-system/internal/session"
)

func EnableMFA() {
	if session.CurrentUser.MFAEnabled {
		fmt.Println("\n✅ MFA is already enabled.")
		return
	}

	key, err := mfa.GenerateSecret(session.CurrentUser.Username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nOpen this URL in any QR generator or scan it directly:")
	fmt.Println(key.URL())

	err = repository.EnableMFA(session.CurrentUser.Username, key.Secret())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println(key.URL())

	session.CurrentUser.MFAEnabled = true
	session.CurrentUser.TOTPSecret = key.Secret()

	fmt.Println("\n✅ MFA Enabled Successfully")
}
