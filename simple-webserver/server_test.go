package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHealthz(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthz)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	body := rr.Body.String()
	expectedBody := `{"You good?":"Cannot be better"}`
	if status != http.StatusOK || body != expectedBody {
		t.Errorf("Expected %v but got %v", body, expectedBody)
	}

}
