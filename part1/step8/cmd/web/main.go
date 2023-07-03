package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("server is starting on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
