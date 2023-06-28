package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from home page"))
	if err != nil {
		log.Panicf("error write to response: %v", err)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("server is starting on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
