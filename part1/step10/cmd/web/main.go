package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server that serves files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Use mux.Handle() function to register the file server as the handler for all URL paths that start with "/static/" for matching paths, we strip the "/static" prefix before the request reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("server is starting on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
