package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
