package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name      string `form:"name" json:"name" binding:"required" gorm:"not null"`
	Image     string `json:"image" gorm:"not null"`
	Price     int64  `form:"price" json:"price" binding:"required" gorm:"not null"`
	CashierID uint
}
