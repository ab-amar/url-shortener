package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

type errorResponse struct {
	Error string `json:"error"`
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("panic", "panic", r)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorResponse{Error: "internal server error"})
			}
		}()
		next.ServeHTTP(w, req)
	})
}

func AppHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("X-App-Name", "url-shortener")
		next.ServeHTTP(w, req)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, req)
		slog.Info("request completed",
			"method", req.Method,
			"path", req.URL.Path,
			"status", recorder.statusCode,
		)
	})
}
