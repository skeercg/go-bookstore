package service

import "go-bookstore/pkg/repository"

type AuthService interface {
	CreateUser() (string, error)
	GetUser() (string, error)
}

type Bookstore interface {
	Create() error
	GetAll() error
	GetById() error
	Delete() error
	Update() error
}

type Service struct {
	AuthService
	Bookstore
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repos.Authorization),
		Bookstore:   NewBookstoreService(repos.Bookstore),
	}
}
