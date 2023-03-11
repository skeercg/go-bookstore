package controller

import (
	"github.com/gorilla/mux"
	"go-bookstore/pkg/service"
)

type Controller struct {
	services *service.Service
}

func NewController(services *service.Service) *Controller {
	return &Controller{services: services}
}

func (c *Controller) InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/sign-in", c.signIn).Methods("POST")

	auth.HandleFunc("/sign-up", c.signUp).Methods("POST")

	books := router.PathPrefix("/books").Subrouter()

	books.HandleFunc("", c.getBooks).Methods("GET")

	books.HandleFunc("/{id}", c.getBookById).Methods("GET")

	books.HandleFunc("/{id}", c.deleteBookById).Methods("DELETE")

	books.HandleFunc("", c.createBook).Methods("POST")

	books.HandleFunc("/{id}", c.updateBookById).Methods("PUT")

	return router
}
