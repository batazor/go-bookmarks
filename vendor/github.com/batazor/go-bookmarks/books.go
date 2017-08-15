package main

import (
	"github.com/pressly/chi"
	"net/http"
)

type booksResource struct{}

// Routes creates a REST router for the todos resource
func (rs booksResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // Get /books - read a list of books
	r.Post("/", rs.Create) // Post /books - create a new and book persist it
	r.Put("/", rs.Put)

	r.Route("/:id", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /books/:id - read a single book by :id
		r.Put("/", rs.Update)    // PUT /books/:id - update a single book by :id
		r.Delete("/", rs.Delete) // DELETE /books/:id - delete a single book by :id
	})

	return r
}

func (rs booksResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get list books"))
}

func (rs booksResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create books"))
}

func (rs booksResource) Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update books"))
}

func (rs booksResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get book"))
}

func (rs booksResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update book"))
}

func (rs booksResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete book"))
}
