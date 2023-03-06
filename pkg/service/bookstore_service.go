package service

import "go-bookstore/pkg/repository"

type BookstoreService struct {
	repo repository.Bookstore
}

func NewBookstoreService(repo repository.Bookstore) *BookstoreService {
	return &BookstoreService{repo: repo}
}

func (r *BookstoreService) Create() error {
	return nil
}

func (r *BookstoreService) GetAll() error {
	return nil
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
