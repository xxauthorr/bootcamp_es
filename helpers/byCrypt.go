package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ByCrypt struct{}

func (by ByCrypt) HashPassword(password string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
		return ""
	}
	return string(hashPass)
}
func (by ByCrypt) VerifyPassword(userPassword, givenPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), ([]byte(userPassword)))
	check := true
	// msg := ""
	if err != nil {
		check = false
		return check
	}
	return check
}
