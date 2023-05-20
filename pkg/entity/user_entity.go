package entity

import (
	"html"
	"strings"
	// "golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	Username    string
	Password    string
	Nickname    string
	Picture_url string
}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func CheckPassword(input, password string) bool {
	return input == password
}
