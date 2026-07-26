package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
