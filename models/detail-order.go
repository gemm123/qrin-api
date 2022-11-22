package models

import "gorm.io/gorm"

type DetailOrder struct {
	gorm.Model
	Quantity int64
	Price    int64
	OrderID  string
	ItemID   uint
	Item     Item
}
