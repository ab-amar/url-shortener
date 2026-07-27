package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type fakeURLService struct{}

func (s fakeURLService) Shorten(originalURL string) model.URL {
	return model.URL{
		OriginalURL: originalURL,
		ShortCode:   "abcd1001",
	}
}

func (s fakeURLService) Resolve(code string) (model.URL, bool) {
	return model.URL{
		OriginalURL: "https://example.com",
		ShortCode:   "abcd1001",
	}, true
}
func TestShortenHandler(t *testing.T) {
	target := "/shorten"
	h := New(fakeURLService{})

	tests := []struct {
		name                string
		method              string
		reqBodyString       string
		expectedStatus      int
		expectedContentType string
	}{
		{
			name:                "Success",
			method:              http.MethodPost,
			expectedStatus:      http.StatusOK,
			reqBodyString:       `{"url":"https://example.com"}`,
			expectedContentType: "application/json",
		},
		{
			name:                "Get not allowed",
			method:              http.MethodGet,
			expectedStatus:      http.StatusMethodNotAllowed,
			reqBodyString:       `{"url":"https://example.com"}`,
			expectedContentType: "application/json",
		},
		{
			name:                "Bad url",
			method:              http.MethodPost,
			expectedStatus:      http.StatusBadRequest,
			reqBodyString:       `{"url":"https//example.com"}`,
			expectedContentType: "application/json",
		},
		{
			name:                "Empty url",
			method:              http.MethodPost,
			expectedStatus:      http.StatusBadRequest,
			reqBodyString:       `{"url":""}`,
			expectedContentType: "application/json",
		},
		{
			name:                "Bad req body",
			method:              http.MethodPost,
			expectedStatus:      http.StatusBadRequest,
			reqBodyString:       `{"url":""`,
			expectedContentType: "application/json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.reqBodyString)

			req := httptest.NewRequest(tt.method, target, body)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			h.ShortenHandler(resp, req)

			if resp.Code != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, resp.Code)
			}

			if resp.Header().Get("Content-Type") != tt.expectedContentType {
				t.Fatalf("Did not get Content-Type %s", tt.expectedContentType)
			}

		})
	}
}
func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name                string
		method              string
		expectedStatus      int
		expectedBody        string
		expectedContentType string
	}{
		{
			name:                "Get",
			method:              http.MethodGet,
			expectedStatus:      http.StatusOK,
			expectedBody:        "Ok!",
			expectedContentType: "text/plain",
		},
		{
			name:                "Post",
			method:              http.MethodPost,
			expectedStatus:      http.StatusMethodNotAllowed,
			expectedBody:        "",
			expectedContentType: "",
		},
		{
			name:                "Put",
			method:              http.MethodPut,
			expectedStatus:      http.StatusMethodNotAllowed,
			expectedBody:        "",
			expectedContentType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/health", nil)
			resp := httptest.NewRecorder()
			HealthHandler(resp, req)

			if resp.Code != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, resp.Code)
			}

			if tt.name == "Get" {
				if resp.Body.String() != tt.expectedBody {
					t.Fatalf("Did not get status %s", tt.expectedBody)
				}

				if resp.Header().Get("Content-Type") != tt.expectedContentType {
					t.Fatalf("Did not get Content-Type %s", tt.expectedContentType)
				}
			}

		})
	}
}

func TestReadyHandler(t *testing.T) {
	tests := []struct {
		name                string
		method              string
		expectedStatus      int
		expectedBody        string
		expectedContentType string
	}{
		{
			name:                "Get",
			method:              http.MethodGet,
			expectedStatus:      http.StatusOK,
			expectedBody:        "Ready!",
			expectedContentType: "text/plain",
		},
		{
			name:                "Post",
			method:              http.MethodPost,
			expectedStatus:      http.StatusMethodNotAllowed,
			expectedBody:        "",
			expectedContentType: "",
		},
		{
			name:                "Put",
			method:              http.MethodPut,
			expectedStatus:      http.StatusMethodNotAllowed,
			expectedBody:        "",
			expectedContentType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/ready", nil)
			resp := httptest.NewRecorder()
			ReadyHandler(resp, req)

			if resp.Code != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, resp.Code)
			}

			if tt.name == "Get" {
				if resp.Body.String() != tt.expectedBody {
					t.Fatalf("Did not get status %s", tt.expectedBody)
				}

				if resp.Header().Get("Content-Type") != tt.expectedContentType {
					t.Fatalf("Did not get Content-Type %s", tt.expectedContentType)
				}
			}

		})
	}
}

type fakeNotFoundURLService struct{}

func (s fakeNotFoundURLService) Shorten(originalURL string) model.URL {
	return model.URL{}
}

func (s fakeNotFoundURLService) Resolve(code string) (model.URL, bool) {
	return model.URL{}, false
}
func TestCodeHandler(t *testing.T) {
	target := "/abcd1235"
	h := New(fakeNotFoundURLService{})

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("code", "abcd1235")
	req := httptest.NewRequest(http.MethodGet, target, nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	resp := httptest.NewRecorder()
	h.CodeHandler(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))

}
