package main

import (
	"log"
	"net/http"

	"github.com/kytnacode/simple-crud/internal/handlers"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
  
	log.Println("Starting server")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.Ping)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
