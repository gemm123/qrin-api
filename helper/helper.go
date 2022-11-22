package helper

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateCodeImage() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-0001) + 0001
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
