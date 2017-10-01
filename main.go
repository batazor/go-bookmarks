package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/batazor/go-bookmarks/db"
	"github.com/batazor/go-bookmarks/handlers/book"
	"github.com/batazor/go-bookmarks/utils"
	"github.com/go-chi/render"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"net/http"
	"github.com/batazor/go-bookmarks/handlers/prometheus"
)

var log = logrus.New()

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	// Connect to MongoDB
	db.Connect()

	// Init Prometheus
	go prometheus.Init()
}

func main() {

	// Get configuration ======================================================
	PORT := utils.Getenv("PORT", "4070")

	// Routes ==================================================================
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", HelloWorld)
	r.Mount("/book", book.Routes())

	// start HTTP-server
	log.Info("Run services on port " + PORT)
	http.ListenAndServe(":"+PORT, r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}
