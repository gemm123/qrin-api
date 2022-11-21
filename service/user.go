package service

import (
	"gemm123/qrin-api/config"
	"gemm123/qrin-api/helper"
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	repository repository.Repository
}

type Service interface {
	Register(user models.User) (models.User, error)
	CheckEmail(email string) error
	CheckPassword(email, password string) (bool, error)
	GetUser(email string) (models.User, error)
	GenerateToken(id uint, email, name, image, phone, role string, budget, otp int) (string, error)
	GetUserByID(id uint) (models.User, error)
}

func NewService(repository repository.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Register(user models.User) (models.User, error) {
	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}

func (s *service) GetUser(email string) (models.User, error) {
	user, err := s.repository.GetUser(email)
	return user, err
}

func (s *service) GetUserByID(id uint) (models.User, error) {
	user, err := s.repository.GetUserByID(id)
	return user, err
}

func (s *service) CheckEmail(email string) error {
	err := s.repository.FindEmail(email)
	return err
}

func (s *service) CheckPassword(email, password string) (bool, error) {
	passwordHas, err := s.repository.GetPassByEmail(email)
	if err != nil {
		return false, err
	}

	ok := helper.CheckPasswordHash(password, passwordHas)
	return ok, err
}

func (s *service) GenerateToken(id uint, email, name, image, phone, role string, budget, otp int) (string, error) {
	claims := config.CustomClaims{
		Id:     id,
		Email:  email,
		Name:   name,
		Image:  image,
		Phone:  phone,
		Budget: budget,
		Otp:    otp,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
			Issuer:    "qrin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JwtKey))
	return signedToken, err
}
