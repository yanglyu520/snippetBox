package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

// for /snippet/view?id=1 path
func snippetView(w http.ResponseWriter, r *http.Request) {
	// retrieve the value of the id parameter from the URL query string, which we can do using the r.URL.Query().Get() method
	// this always return a string value for a parameter, or the empty string "" if no matching parameter exists
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Get view for snippet with id of %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
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
