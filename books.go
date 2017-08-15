package main

import (
	"github.com/pressly/chi"
	"net/http"
)

type linksResource struct{}

// Routes creates a REST router for the todos resource
func (rs linksResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // Get /links - read a list of links
	r.Post("/", rs.Create)  // Post /link - create a new and link persist it

	r.Route("/:id", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /link/:id - read a single link by :id
		r.Put("/", rs.Update)    // PUT /links/:id - update a single link by :id
		r.Delete("/", rs.Delete) // DELETE /links/:id - delete a single link by :id
	})

	return r
}

func (rs linksResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get list links"))
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
