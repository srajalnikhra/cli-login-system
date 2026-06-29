package cli

import (
	"fmt"

	"github.com/srajalnikhra/cli-login-system/internal/repository"
	"github.com/srajalnikhra/cli-login-system/internal/session"
)

func DisableMFA() {
	if !session.CurrentUser.MFAEnabled {
		fmt.Println("\n❌ MFA is already disabled.")
		return
	}

	err := repository.DisableMFA(session.CurrentUser.Username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	session.CurrentUser.MFAEnabled = false
	session.CurrentUser.TOTPSecret = ""

	fmt.Println("\n✅ MFA Disabled Successfully")
}
