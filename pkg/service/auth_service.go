package service

import "go-bookstore/pkg/repository"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) CreateUser() (string, error) {
	return "", nil
}

func (r *AuthService) GetUser() (string, error) {
	return "", nil
}
