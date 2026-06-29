package repository

import "github.com/srajalnikhra/cli-login-system/internal/database"

func EnableMFA(username, secret string) error {
	query := `
	UPDATE users
	SET mfa_enabled = true,
	    totp_secret = $1
	WHERE username = $2
	`

	_, err := database.DB.Exec(query, secret, username)
	return err
}
