package service

import (
	"go-bookstore/pkg/model"
	"go-bookstore/pkg/repository"
)

type BookstoreService struct {
	repo repository.Bookstore
}

func NewBookstoreService(repo repository.Bookstore) *BookstoreService {
	return &BookstoreService{repo: repo}
}

func (r *BookstoreService) Create(book model.Book) error {
	err := r.repo.Create(book)

	return err
}

func (r *BookstoreService) GetAll(title, sort string) ([]model.Book, error) {
	books, err := r.repo.GetAll(title, sort)

	return books, err
}

func (r *BookstoreService) GetById() error {
	return nil
}

func (r *BookstoreService) Delete() error {
	return nil
}

func (r *BookstoreService) Update() error {
	return nil
}
