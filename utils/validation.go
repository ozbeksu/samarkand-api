package utils

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func IsValidPassword(encrypted, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password)) == nil
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
