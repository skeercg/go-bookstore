package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-bookstore/pkg/model"
	"go-bookstore/pkg/repository"
	"time"
)

var signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")

const (
	salt     = "hjqrhjqw124617ajfhajs"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type AuthServiceService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthServiceService {
	return &AuthServiceService{repo: repo}
}

func (r *AuthServiceService) CreateUser(user model.User) error {
	user.Password = generatePasswordHash(user.Password)
	err := r.repo.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthServiceService) GenerateToken(username, password string) (string, error) {
	user, err := r.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		}})

	return token.SignedString(signingKey)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
