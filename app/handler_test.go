package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidZipcode(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather?cep=123", nil)
	req.Header.Set("x-api-key", "test-api-key")
	w := httptest.NewRecorder()

	WeatherHandler(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", w.Code)
	}
}
