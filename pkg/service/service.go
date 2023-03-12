package service

import (
	"go-bookstore/pkg/model"
	"go-bookstore/pkg/repository"
)

type AuthService interface {
	CreateUser(user model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Bookstore interface {
	Create(book model.Book) error
	GetAll(title, sort string) ([]model.Book, error)
	GetById(id int) (model.Book, error)
	Delete(id int) error
	Update(book model.Book, id int) error
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
