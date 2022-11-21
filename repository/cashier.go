package repository

import (
	"gemm123/qrin-api/models"

	"gorm.io/gorm"
)

type repositoryCashier struct {
	DB *gorm.DB
}

type RepositoryCashier interface {
	CreateCashier(cashier models.Cashier) (models.Cashier, error)
}

func NewRepositoryCashier(DB *gorm.DB) *repositoryCashier {
	return &repositoryCashier{DB: DB}
}

func (r *repositoryCashier) CreateCashier(cashier models.Cashier) (models.Cashier, error) {
	err := r.DB.Create(&cashier).Error
	return cashier, err
}
