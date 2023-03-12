package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-bookstore/pkg/model"
	"go-bookstore/pkg/repository"
	"time"
)

var signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type tokenClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (r *AuthorizationService) CreateUser(user model.User) error {
	user.Password = generatePasswordHash(user.Password)
	err := r.repo.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationService) GenerateToken(username, password string) (string, error) {
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

func (r *AuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return signingKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
