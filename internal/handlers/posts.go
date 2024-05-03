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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "creating post...")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updating post")
}

func PatchPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "patching post...")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "deleting post...")
}

func Default(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET,POST")
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func DefaultWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET,PUT,PATCH,DELETE")
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func RegisterPostRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /posts", GetPosts)
	mux.HandleFunc("GET /posts/{id}", GetPostById)
	mux.HandleFunc("POST /posts", CreatePost)
	mux.HandleFunc("PUT /posts/{id}", UpdatePost)
	mux.HandleFunc("PATCH /posts/{id}", PatchPost)
	mux.HandleFunc("DELETE /posts/{id}", DeletePost)
	mux.HandleFunc("/posts", Default)
	mux.HandleFunc("/posts/{id}", DefaultWithId)
}
