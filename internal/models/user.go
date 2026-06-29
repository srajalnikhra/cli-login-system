package models

import "time"

type User struct {
	ID           int
	Username     string
	PasswordHash string
	MFAEnabled   bool
	TOTPSecret   string
	CreatedAt    time.Time
}
