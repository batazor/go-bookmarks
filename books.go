package main

import (
	"github.com/pressly/chi"
	"net/http"
)

type linksResource struct{
	isbn string
	title string
	author string
	price float32
}

// Routes creates a REST router for the todos resource
func (rs linksResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // Get /book - read a list of books
	r.Post("/", rs.Create)  // Post /book - create a new book

	r.Route("/:id", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /book/:id - read a single book by :id
		r.Put("/", rs.Update)    // PUT /book/:id - update a single book by :id
		r.Delete("/", rs.Delete) // DELETE /book/:id - delete a single book by :id
	})

	return r
}

func (rs linksResource) List(w http.ResponseWriter, r *http.Request) {
	mongo := mgoSession.Clone()

	w.Write([]byte("Get list links"))

	defer mongo.Close()
}

func (rs linksResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create link"))
}

func (rs linksResource) Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update links"))
}

func (rs linksResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get link"))
}

func (rs linksResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update link"))
}

func (rs linksResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete link"))
}
