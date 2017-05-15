package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

func main() {
	// Get configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}

	// if no config is found, use the default(s)
	viper.SetDefault("server.port", "4070")

	PORT := viper.GetString("server.port")

	// routing
	http.HandleFunc("/", HelloWorld)

	// start HTTP-server
	http.ListenAndServe(":"+PORT, nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!!!")
}
