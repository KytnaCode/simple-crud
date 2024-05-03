package handlers

import (
	"fmt"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getting posts...")
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getting a post by id...")
}

func RegisterPostRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /posts", GetPosts)
	mux.HandleFunc("GET /posts/{id}", GetPostById)
}
