package service

import (
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
)

type serviceItem struct {
	repositoryItem repository.RepositoryItem
}

type ServiceItem interface {
	AddItem(item models.Item) (models.Item, error)
}

func NewServiceItem(repositoryItem repository.RepositoryItem) *serviceItem {
	return &serviceItem{repositoryItem: repositoryItem}
}

func (s *serviceItem) AddItem(item models.Item) (models.Item, error) {
	newItem, err := s.repositoryItem.CreateItem(item)
	return newItem, err
}
