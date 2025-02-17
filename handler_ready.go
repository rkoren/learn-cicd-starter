package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	log.Println("Received health check request") // Add this
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("ok"))
}
