package handlers

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{ \"message\": \"ok\" }")
}
