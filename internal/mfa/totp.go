package mfa

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// GenerateSecret creates a new TOTP secret
// for the specified user.
func GenerateSecret(username string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      "CLI Login System",
		AccountName: username,
	})
}

// VerifyOTP validates the OTP entered
// by the user.
func VerifyOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}
