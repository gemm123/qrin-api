package repository

import (
	"gemm123/qrin-api/models"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateUser(user models.User) (models.User, error)
	FindEmail(email string) error
	GetPassByEmail(email string) (string, error)
	GetUser(email string) (models.User, error)
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB: DB}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *repository) GetUser(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *repository) FindEmail(email string) error {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return err
}

func (r *repository) GetPassByEmail(email string) (string, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user.Password, err
}
