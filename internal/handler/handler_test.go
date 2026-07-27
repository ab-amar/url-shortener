package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ab-amar/url-shortener/internal/model"
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
			name:                "Get sot allowed",
			method:              http.MethodGet,
			expectedStatus:      http.StatusMethodNotAllowed,
			reqBodyString:       `{"url":"https://example.com"}`,
			expectedContentType: "",
		},
		{
			name:                "Bad url",
			method:              http.MethodPost,
			expectedStatus:      http.StatusBadRequest,
			reqBodyString:       `{"url":"https//example.com"}`,
			expectedContentType: "",
		},
		{
			name:                "Empty url",
			method:              http.MethodPost,
			expectedStatus:      http.StatusBadRequest,
			reqBodyString:       `{"url":""}`,
			expectedContentType: "",
		},
		{
			name:                "Bad req body",
			method:              http.MethodPost,
			expectedStatus:      http.StatusBadRequest,
			reqBodyString:       `{"url":""`,
			expectedContentType: "",
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

			if resp.Code == http.StatusOK {
				if resp.Header().Get("Content-Type") != tt.expectedContentType {
					t.Fatalf("Did not get Content-Type %s", tt.expectedContentType)
				}
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
