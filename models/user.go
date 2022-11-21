package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Name     string `json:"name" gorm:"not null" binding:"required"`
	Image    string `json:"image"`
	Phone    string `json:"phone" gorm:"not null" binding:"required"`
	Budget   int    `json:"budget" gorm:"default:0"`
}

type InputLoginUser struct {
	Email    string `json:"email" gorm:"-:migration" binding:"required"`
	Password string `json:"password" gorm:"-:migration" binding:"required"`
}
