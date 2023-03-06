package service

import "go-bookstore/pkg/repository"

type AuthServiceService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthServiceService {
	return &AuthServiceService{repo: repo}
}

func (r *AuthServiceService) CreateUser() (string, error) {
	return "", nil
}

func (r *AuthServiceService) GetUser() (string, error) {
	return "", nil
}
