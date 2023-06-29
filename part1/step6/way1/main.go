package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("Hello from home page"))
	if err != nil {
		log.Panicf("error write to response: %v", err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("get/view snippet here..."))
	if err != nil {
		log.Panicf("error write to response: %v", err)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// Add Allow header to let user know which request methods are supported for a particular URL.
		w.Header().Set("Allow", "POST")
		// If the request method is not POST, it will send 405 status code and a "method not allowed" in response body
		w.WriteHeader(405)

		_, err := w.Write([]byte("Method Not Allowed"))
		if err != nil {
			log.Panicf("error write to response: %v", err)
		}
		return
	}
	_, err := w.Write([]byte("create snippet here..."))
	if err != nil {
		log.Panicf("error write to response: %v", err)
	}
}

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
