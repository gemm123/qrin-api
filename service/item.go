package service

import (
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
	"strings"
)

type serviceItem struct {
	repositoryItem repository.RepositoryItem
}

type ServiceItem interface {
	AddItem(item models.Item) (models.Item, error)
	ShowAllItem(cashierID uint) ([]models.Item, error)
	FilterItem(items []models.Item, inputItem string) []models.Item
	ShowDetailItem(itemID, cashierID uint) (models.Item, error)
}

func NewServiceItem(repositoryItem repository.RepositoryItem) *serviceItem {
	return &serviceItem{repositoryItem: repositoryItem}
}

func (s *serviceItem) AddItem(item models.Item) (models.Item, error) {
	newItem, err := s.repositoryItem.CreateItem(item)
	return newItem, err
}

func (s *serviceItem) ShowAllItem(cashierID uint) ([]models.Item, error) {
	items, err := s.repositoryItem.GetAllItem(cashierID)
	return items, err
}

func (s *serviceItem) FilterItem(items []models.Item, inputItem string) []models.Item {
	var filteredItems []models.Item
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(inputItem)) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

func (s *serviceItem) ShowDetailItem(itemID, cashierID uint) (models.Item, error) {
	item, err := s.repositoryItem.GetDetailItem(itemID, cashierID)
	return item, err
}
