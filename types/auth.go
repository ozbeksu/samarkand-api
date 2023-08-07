package types

import (
	"fmt"
	"github.com/ozbeksu/samarkand-api/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost    = 12
	minTextLength = 3
	minPassLength = 6
)

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s AuthParams) Validate() map[string][]string {
	messages := make(map[string][]string)

	if !utils.IsValidEmail(s.Email) {
		messages["email"] = append(messages["email"], fmt.Sprintf("email is invalid"))
	}
	if len(s.Password) < minPassLength {
		messages["password"] = append(messages["password"], fmt.Sprintf("password length must be at least %d characters", minPassLength))
	}

	return messages
}

func HashPassword(password string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
