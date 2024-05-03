package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kytnacode/simple-crud/internal/handlers"
)

func TestGetPosts_returnOkResponseCode(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := http.StatusOK

	r, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	handlers.GetPosts(rr, r)

	// Assert
	if rr.Code != expected {
		t.Errorf("non 200 response code: expected %v actual %v", expected, rr.Code)
	}
}

