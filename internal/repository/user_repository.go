package repository

import (
	"github.com/srajalnikhra/cli-login-system/internal/database"
	"github.com/srajalnikhra/cli-login-system/internal/models"
)

func CreateUser(user models.User) error {
	query := `
	INSERT INTO users (username, password_hash, mfa_enabled, totp_secret)
	VALUES ($1, $2, $3, $4)
	`

	_, err := database.DB.Exec(
		query,
		user.Username,
		user.PasswordHash,
		user.MFAEnabled,
		user.TOTPSecret,
	)

	return err
}

func FindUserByUsername(username string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, username, password_hash, mfa_enabled, totp_secret, created_at
	FROM users
	WHERE username = $1
	`

	err := database.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.MFAEnabled,
		&user.TOTPSecret,
		&user.CreatedAt,
	)

	return user, err
}
