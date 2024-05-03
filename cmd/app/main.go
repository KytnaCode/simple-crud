package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kytnacode/simple-crud/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.Ping)

	handlers.RegisterPostRoutes(mux)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
