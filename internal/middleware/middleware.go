package middleware

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func AppHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("X-App-Name", "url-shortener")
		next.ServeHTTP(w, req)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
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
			"request_id", GetRequestID(req),
		)
	})
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

type contextKey string

const requestIDKey contextKey = "request_id"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqID := strconv.FormatInt(time.Now().UnixNano(), 10)
		w.Header().Set("X-Request-ID", reqID)
		ctx := context.WithValue(req.Context(), requestIDKey, reqID)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func GetRequestID(req *http.Request) string {
	value := req.Context().Value(requestIDKey)
	reqID, ok := value.(string)

	if !ok {
		return ""
	}
	return reqID
}
