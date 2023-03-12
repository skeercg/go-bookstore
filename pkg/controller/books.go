package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-bookstore/pkg/model"
	"log"
	"net/http"
	"strconv"
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
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	book, err := c.services.Bookstore.GetById(id)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func (c *Controller) deleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Bookstore.Delete(id)
	if err != nil {
		return
	}
}

func (c *Controller) createBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Bookstore.Create(book)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func (c *Controller) updateBookById(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Bookstore.Update(book, id)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}
