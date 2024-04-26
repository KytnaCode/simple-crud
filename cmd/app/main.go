package main

import (
	"log"
	"net/http"

	"github.com/kytnacode/simple-crud/internal/handlers"
)

func main() {
	log.Println("Starting server")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.Ping)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
