package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
