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

func TestGetPostById_returnOkResponse(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := http.StatusOK

	r, err := http.NewRequest("GET", "/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	handlers.GetPostById(rr, r)

	// Assert
	if rr.Code != expected {
		t.Errorf("non 200 response code: expected %v actual %v", expected, rr.Code)
	}
}

func TestCreatePost_returnCreatedResponse(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := http.StatusCreated

	r, err := http.NewRequest("POST", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	handlers.CreatePost(rr, r)

	// Assert
	if rr.Code != expected {
		t.Errorf("non 200 response code: expected %v actual %v", expected, rr.Code)
	}
}

func TestUpdatePost_returnOkResponse(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := http.StatusOK

	r, err := http.NewRequest("PUT", "/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	handlers.UpdatePost(rr, r)

	// Assert
	if rr.Code != expected {
		t.Errorf("non 200 response code: expected %v actual %v", expected, rr.Code)
	}
}

