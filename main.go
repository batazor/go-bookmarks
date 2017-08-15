package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/spf13/viper"
	"net/http"
	"github.com/batazor/go-bookmarks/db"
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
}

func main() {

	// Get configuration ======================================================
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Panic("No configuration file loaded - using defaults")
	}

	// if no config is found, use the default(s)
	viper.SetDefault("server.port", "4070")
	PORT := viper.GetString("server.port")

	// Routes ==================================================================
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", HelloWorld)
	r.Mount("/book", linksResource{}.Routes())

	// start HTTP-server
	log.Info("Run services on port " + PORT)
	http.ListenAndServe(":"+PORT, r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}
