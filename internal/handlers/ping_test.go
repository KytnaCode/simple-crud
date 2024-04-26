package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kytnacode/simple-crud/internal/handlers"
)

func TestPing_returnCodeOk(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := 200

	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	handlers.Ping(rr, r)

	// Assert
	if rr.Code != expected {
		t.Errorf("invalid response code: expected %v actual %v", expected, rr.Code)
	}
}

func TestPing_returnExpectedResponse(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := []byte("{ \"message\": \"ok\" }")

	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	handlers.Ping(rr, r)

	// Assert
	if !bytes.Equal(rr.Body.Bytes(), expected) {
		t.Errorf(
			"invalid response: expected \"%v\" actual \"%v\"",
			string(expected),
			rr.Body.String(),
		)
	}
}

func TestPing_returnAJSON(t *testing.T) {
	t.Parallel()

	// Arrange
	expected := "application/json"

	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	handlers.Ping(rr, r)

	if rr.Header().Get("Content-Type") != expected {
		t.Errorf(
			"invalid Content-Type: expected %v actual %v",
			expected,
			rr.Header().Get("Content-Type"),
		)
	}
}
