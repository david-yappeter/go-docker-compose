package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	prefix := os.Getenv("PREFIX")

	router := mux.NewRouter()

	subRouter := router.PathPrefix(prefix).Subrouter()

	subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	subRouter.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Myself")
	})

	log.Printf("Listen and Serve at http://localhost:%s\n", port)
	log.Println(http.ListenAndServe(":"+port, router))
}
