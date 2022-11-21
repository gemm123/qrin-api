package config

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = os.Getenv("JWT_KEY")

type CustomClaims struct {
	Id     uint
	Email  string
	Name   string
	Image  string
	Phone  string
	Budget int
	Otp    int
	Role   string
	jwt.RegisteredClaims
}
