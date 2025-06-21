package main

import (
	"log"
	"net/http"
)

func main() {
	router, err := NewRouter()
	if err != nil {
		log.Fatalf("Failed to create router: %v", err)
	}
	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
