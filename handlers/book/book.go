package book

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/batazor/go-bookmarks/db"
	"github.com/batazor/go-bookmarks/models"
	"github.com/pressly/chi"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

var log = logrus.New()

// Routes creates a REST router for the todos resource
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", List)    // Get /book - read a list of books
	r.Post("/", Create) // Post /book - create a new book

	r.Route("/:id", func(r chi.Router) {
		r.Get("/", Get)       // GET /book/:id - read a single book by :id
		r.Put("/", Update)    // PUT /book/:id - update a single book by :id
		r.Delete("/", Delete) // DELETE /book/:id - delete a single book by :id
	})

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}
	err := db.Session.DB("books").C(models.CollectionBook).Find(nil).Sort("-updated_on").All(&books)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Get list links"))
	}

	res, err := json.Marshal(&books)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func Create(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Don't create link"))
		return
	}

	var book models.Book
	err = json.Unmarshal(b, &book)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Don't create link"))
		return
	}

	id := bson.NewObjectId()
	book.Id = id
	err = db.Session.DB("books").C(models.CollectionBook).Insert(book)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Don't create link"))
		return
	}

	output, err := json.Marshal(book)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Don't create link"))
		return
	}
	w.Write(output)
}

func Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update links"))
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get link"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update link"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete link"))
}
