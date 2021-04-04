package services

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomPassword() string {

	password := base64.StdEncoding.EncodeToString([]byte(random(9)))

	return (string(password))
}

func random(number int) string {
	b := make([]byte, number)

	_, err := rand.Read(b)

	if err != nil {
		panic("cannot generate password")
	}

	return string(b)
}
