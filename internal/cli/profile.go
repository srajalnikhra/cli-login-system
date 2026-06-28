package cli

import (
	"fmt"

	"github.com/srajalnikhra/cli-login-system/internal/session"
)

func Profile() {
	fmt.Println("\n===== Profile =====")
	fmt.Println("Username:", session.CurrentUser.Username)
	fmt.Println("MFA Enabled:", session.CurrentUser.MFAEnabled)

	if session.CurrentUser.MFAEnabled {
		fmt.Println("Status: MFA Already Enabled")
	} else {
		fmt.Println("Status: MFA Disabled")
	}

	fmt.Println("===================")
}