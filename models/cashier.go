package models

import "gorm.io/gorm"

type Cashier struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Name     string `json:"name" gorm:"not null" binding:"required"`
	Image    string `json:"image"`
	Phone    string `json:"phone" gorm:"not null" binding:"required"`
	Shop     string `json:"shop" gorm:"not null" binding:"required"`
	Address  string `json:"address" gorm:"not null" binding:"required"`
	Items    []Item
}

type InputLoginCashier struct {
	Email    string `json:"email" gorm:"-:migration" binding:"required"`
	Password string `json:"password" gorm:"-:migration" binding:"required"`
}
