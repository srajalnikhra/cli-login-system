package cli

import (
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/srajalnikhra/cli-login-system/internal/mfa"
	"github.com/srajalnikhra/cli-login-system/internal/repository"
	"github.com/srajalnikhra/cli-login-system/internal/session"
)

// EnableMFA generates a TOTP secret,
// displays a QR code, and enables MFA.
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

	fmt.Println()
	fmt.Println("Scan this QR Code with Google Authenticator:")
	fmt.Println()

	qr, err := qrcode.New(key.URL(), qrcode.Medium)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bitmap := qr.Bitmap()

	for _, row := range bitmap {
		for _, cell := range row {
			if cell {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	err = repository.EnableMFA(session.CurrentUser.Username, key.Secret())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	session.CurrentUser.MFAEnabled = true
	session.CurrentUser.TOTPSecret = key.Secret()

	fmt.Println("\n✅ MFA Enabled Successfully")
}
