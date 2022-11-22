package service

import (
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
)

type serviceOrder struct {
	repositoryOrder repository.RepositoryOrder
}

type ServiceOrder interface {
	AddOrder(order models.Order) (models.Order, error)
	AddDetailOrder(detailOrder models.DetailOrder) (models.DetailOrder, error)
}

func NewServiceOrder(repositoryOrder repository.RepositoryOrder) *serviceOrder {
	return &serviceOrder{repositoryOrder: repositoryOrder}
}

func (s *serviceOrder) AddOrder(order models.Order) (models.Order, error) {
	newOrder, err := s.repositoryOrder.CreateOrder(order)
	return newOrder, err
}

func (s *serviceOrder) AddDetailOrder(detailOrder models.DetailOrder) (models.DetailOrder, error) {
	newDetailOrder, err := s.repositoryOrder.CreateDetailOrder(detailOrder)
	return newDetailOrder, err
}
