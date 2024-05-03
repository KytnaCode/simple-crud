package handlers

import (
	"fmt"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getting posts...")
}

func RegisterPostRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /posts", GetPosts)
}
