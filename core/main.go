package main

import (
	"log"
	"net/http"

	"debugr/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/trace", handler.TraceHandler)

	log.Println("[debugr] Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
