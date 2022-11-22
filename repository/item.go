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
	GetDetailItem(itemID, cashierID uint) (models.Item, error)
	GetDetailItemByName(name string, cashierID uint) (models.Item, error)
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

func (r *repositoryItem) GetDetailItem(itemID, cashierID uint) (models.Item, error) {
	var item models.Item
	err := r.DB.Where("id = ? AND cashier_id = ?", itemID, cashierID).First(&item).Error
	return item, err
}

func (r *repositoryItem) GetDetailItemByName(name string, cashierID uint) (models.Item, error) {
	var item models.Item
	err := r.DB.Where("name = ? AND cashier_id = ?", name, cashierID).First(&item).Error
	return item, err
}
