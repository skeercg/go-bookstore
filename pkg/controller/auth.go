package controller

import (
	"encoding/json"
	"go-bookstore/pkg/model"
	"log"
	"net/http"
)

func (c *Controller) signUp(w http.ResponseWriter, r *http.Request) {

	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Print(err)
	}

	err = c.services.AuthService.CreateUser(u)
	if err != nil {
		return
	}

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *Controller) signIn(w http.ResponseWriter, r *http.Request) {
	var u signInInput

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Print(err)
	}

	token, err := c.services.AuthService.GenerateToken(u.Username, u.Password)
	if err != nil {
		log.Print(err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})

	if err != nil {
		log.Print(err)
	}
}
