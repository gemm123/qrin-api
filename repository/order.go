package repository

import (
	"gemm123/qrin-api/models"

	"gorm.io/gorm"
)

type repositoryOrder struct {
	DB *gorm.DB
}

type RepositoryOrder interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAllOrder(userID uint) ([]models.Order, error)
	CreateDetailOrder(detailOrder models.DetailOrder) (models.DetailOrder, error)
}

func NewRepositoryOrder(DB *gorm.DB) *repositoryOrder {
	return &repositoryOrder{DB: DB}
}

func (r *repositoryOrder) CreateOrder(order models.Order) (models.Order, error) {
	err := r.DB.Create(&order).Error
	return order, err
}

func (r *repositoryOrder) GetAllOrder(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("DetailOrders").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *repositoryOrder) CreateDetailOrder(detailOrder models.DetailOrder) (models.DetailOrder, error) {
	err := r.DB.Create(&detailOrder).Error
	return detailOrder, err
}
