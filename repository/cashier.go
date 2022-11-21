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
	GetCashier(email string) (models.Cashier, error)
	GetCashierByID(id uint) (models.Cashier, error)
	FindEmailCashier(email string) error
	GetPassCashierByEmail(email string) (string, error)
}

func NewRepositoryCashier(DB *gorm.DB) *repositoryCashier {
	return &repositoryCashier{DB: DB}
}

func (r *repositoryCashier) CreateCashier(cashier models.Cashier) (models.Cashier, error) {
	err := r.DB.Create(&cashier).Error
	return cashier, err
}

func (r *repositoryCashier) GetCashier(email string) (models.Cashier, error) {
	var cashier models.Cashier
	err := r.DB.Where("email = ?", email).First(&cashier).Error
	return cashier, err
}

func (r *repositoryCashier) GetCashierByID(id uint) (models.Cashier, error) {
	var cashier models.Cashier
	err := r.DB.Where("id = ?", id).First(&cashier).Error
	return cashier, err
}

func (r *repositoryCashier) FindEmailCashier(email string) error {
	var cashier models.Cashier
	err := r.DB.Where("email = ?", email).First(&cashier).Error
	return err
}

func (r *repositoryCashier) GetPassCashierByEmail(email string) (string, error) {
	var cashier models.Cashier
	err := r.DB.Where("email = ?", email).First(&cashier).Error
	return cashier.Password, err
}
