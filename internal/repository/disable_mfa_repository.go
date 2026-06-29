package repository

import "github.com/srajalnikhra/cli-login-system/internal/database"

func DisableMFA(username string) error {
	query := `
	UPDATE users
	SET mfa_enabled = false,
	    totp_secret = ''
	WHERE username = $1
	`

	_, err := database.DB.Exec(query, username)
	return err
}
