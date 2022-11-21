package service

import (
	"gemm123/qrin-api/config"
	"gemm123/qrin-api/helper"
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type serviceCashier struct {
	repositoryCashier repository.RepositoryCashier
}

type ServiceCashier interface {
	Register(cashier models.Cashier) (models.Cashier, error)
	GetCashierByID(id uint) (models.Cashier, error)
	CheckEmailCashier(email string) error
	CheckPassword(email, password string) (bool, error)
	GenerateToken(id uint) (string, error)
	GetCashier(email string) (models.Cashier, error)
}

func NewServiceCashier(repositoryCashier repository.RepositoryCashier) *serviceCashier {
	return &serviceCashier{repositoryCashier: repositoryCashier}
}

func (s *serviceCashier) Register(cashier models.Cashier) (models.Cashier, error) {
	newCashier, err := s.repositoryCashier.CreateCashier(cashier)
	return newCashier, err
}

func (s *serviceCashier) GetCashier(email string) (models.Cashier, error) {
	cashier, err := s.repositoryCashier.GetCashier(email)
	return cashier, err
}

func (s *serviceCashier) GetCashierByID(id uint) (models.Cashier, error) {
	cashier, err := s.repositoryCashier.GetCashierByID(id)
	return cashier, err
}

func (s *serviceCashier) CheckEmailCashier(email string) error {
	err := s.repositoryCashier.FindEmailCashier(email)
	return err
}

func (s *serviceCashier) CheckPassword(email, password string) (bool, error) {
	passwordHas, err := s.repositoryCashier.GetPassCashierByEmail(email)
	if err != nil {
		return false, err
	}

	ok := helper.CheckPasswordHash(password, passwordHas)
	return ok, err
}

func (s *serviceCashier) GenerateToken(id uint) (string, error) {
	claims := config.CustomClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
			Issuer:    "qrin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JwtKey))
	return signedToken, err
}
