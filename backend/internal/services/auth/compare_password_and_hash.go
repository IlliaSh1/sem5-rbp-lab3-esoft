package services_auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (service *AuthService) ComparePasswordAndHash(password string, password_hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	if err != nil {
		return fmt.Errorf("failed to compare password and hash: %w", err)
	}

	return nil
}
