package service

import (
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
)

type serviceCashier struct {
	repositoryCashier repository.RepositoryCashier
}

type ServiceCashier interface {
	Register(cashier models.Cashier) (models.Cashier, error)
}

func NewServiceCashier(repositoryCashier repository.RepositoryCashier) *serviceCashier {
	return &serviceCashier{repositoryCashier: repositoryCashier}
}

func (s *serviceCashier) Register(cashier models.Cashier) (models.Cashier, error) {
	newCashier, err := s.repositoryCashier.CreateCashier(cashier)
	return newCashier, err
}
