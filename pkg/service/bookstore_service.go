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

func (r *BookstoreService) GetById(id int) (model.Book, error) {
	book, err := r.repo.GetById(id)

	return book, err
}

func (r *BookstoreService) Delete(id int) error {
	err := r.repo.Delete(id)

	return err
}

func (r *BookstoreService) Update(book model.Book, id int) error {
	err := r.repo.Update(book, id)

	return err
}
