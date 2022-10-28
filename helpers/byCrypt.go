package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
		return ""
	}
	return string(hashPass)
}
func verifyPassword(userPassword, givenPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), ([]byte(userPassword)))
	check := true
	// msg := ""
	if err != nil {
		fmt.Println("error at verify password :", err)
		check = false
		return check, err
	}
	return check, err
}
