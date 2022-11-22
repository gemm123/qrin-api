package models

import "time"

type Order struct {
	ID           string    `json:"id" gorm:"not null;unique"`
	TotalPrice   int64     `json:"totalPrice" gorm:"not null"`
	Date         time.Time `json:"date" gorm:"not null"`
	Payment      string    `json:"payment" gorm:"not null"`
	Address      string
	Shop         string
	CashierID    uint
	UserID       uint
	DetailOrders []DetailOrder
}

type InputOrder struct {
	ID         string `json:"id" gorm:"-:migration" binding:"required"`
	CashierID  uint   `json:"cashierId" gorm:"-:migration" binding:"required"`
	Items      []Item `json:"items" gorm:"-:migration" binding:"required"`
	TotalPrice int64  `json:"totalPrice" gorm:"-:migration" binding:"required"`
	Date       string `json:"date" gorm:"-:migration" binding:"required"`
	Payment    string `json:"payment" gorm:"-:migration" binding:"required"`
}
