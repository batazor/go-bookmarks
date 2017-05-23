package main

import (
	"github.com/spf13/viper"
	http "net/http"
	"github.com/pressly/chi"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func main() {

	// Get configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Panic("No configuration file loaded - using defaults")
	}

	// if no config is found, use the default(s)
	viper.SetDefault("server.port", "4070")

	PORT := viper.GetString("server.port")

	// Routes
	r := chi.NewRouter()
	r.Get("/", HelloWorld)

	// start HTTP-server
	http.ListenAndServe(":"+PORT, r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Info("GET")
	w.Write([]byte("Hello World!!!"))
}