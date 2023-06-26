package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
