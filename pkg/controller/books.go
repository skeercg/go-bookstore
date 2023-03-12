package controller

import (
	"encoding/json"
	"go-bookstore/pkg/model"
	"log"
	"net/http"
)

type searchParams struct {
	Title string `json:"title" binding:"required"`
	Sort  string `json:"sort" binding:"required"`
}

func (c *Controller) getBooks(w http.ResponseWriter, r *http.Request) {
	var params searchParams
	err := json.NewDecoder(r.Body).Decode(&params)

	books, err := c.services.Bookstore.GetAll(params.Title, params.Sort)
	if err != nil {
		log.Print(err)
	}

	err = json.NewEncoder(w).Encode(books)
}

func (c *Controller) getBookById(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.GetById()
	if err != nil {
		return
	}
}

func (c *Controller) deleteBookById(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.Delete()
	if err != nil {
		return
	}
}

func (c *Controller) createBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Print(err)
	}

	err = c.services.Bookstore.Create(book)

	if err != nil {
		log.Print(err)
	}
}

func (c *Controller) updateBookById(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.Update()
	if err != nil {
		return
	}
}
