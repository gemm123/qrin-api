package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Image    string `json:"image"`
	Phone    string `json:"phone" gorm:"not null"`
	Budget   int    `json:"budget" gorm:"default:0"`
	OTP      string `json:"otp"`
	Role     string `json:"role" gorm:"not null"`
}
