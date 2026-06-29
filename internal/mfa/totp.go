package mfa

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateSecret(username string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      "CLI Login System",
		AccountName: username,
	})
}

func VerifyOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}
