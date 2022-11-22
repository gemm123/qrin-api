package repository

import (
	"gemm123/qrin-api/models"

	"gorm.io/gorm"
)

type repositoryItem struct {
	DB *gorm.DB
}

type RepositoryItem interface {
	CreateItem(item models.Item) (models.Item, error)
	GetAllItem(cashierID uint) ([]models.Item, error)
}

func NewRepositoryItem(DB *gorm.DB) *repositoryItem {
	return &repositoryItem{DB: DB}
}

func (r *repositoryItem) CreateItem(item models.Item) (models.Item, error) {
	err := r.DB.Create(&item).Error
	return item, err
}

func (r *repositoryItem) GetAllItem(cashierID uint) ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Where("cashier_id = ?", cashierID).Find(&items).Error
	return items, err
}
