package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Get(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()
	HealthHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected staus %d, got %d", http.StatusOK, resp.Code)
	}

	if resp.Body.String() != "Ok!" {
		t.Fatalf("Did not get staus %s", "Ok!")
	}

	if resp.Header().Get("Content-Type") != "text/plain" {
		t.Fatalf("Did not get Content-Type %s", "text/plain")
	}
}

func TestHealthHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	resp := httptest.NewRecorder()
	HealthHandler(resp, req)

	if resp.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected staus %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}
}
