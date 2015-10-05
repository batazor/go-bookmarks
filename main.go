package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output, err := json.Marshal("Hello, wolrd")

		if err != nil {
			fmt.Println("Something went wrong!")
		}

		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}
