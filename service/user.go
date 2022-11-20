package service

import (
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/repository"
)

type service struct {
	repository repository.Repository
}

type Service interface {
	Register(user models.User) (models.User, error)
}

func NewService(repository repository.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Register(user models.User) (models.User, error) {
	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}
