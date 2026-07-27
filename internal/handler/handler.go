package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/ab-amar/url-shortener/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	URLService service.URLService
}

func New(urlService service.URLService) Handler {
	return Handler{
		URLService: urlService,
	}
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Message  string    `json:"message"`
	URLModel model.URL `json:"urlModel"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h Handler) ShortenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}

	dec := json.NewDecoder(req.Body)
	var reqBody shortenRequest
	if err := dec.Decode(&reqBody); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "bad request"})
		return
	}
	urlString := strings.TrimSpace(reqBody.URL)
	if urlString == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "bad request"})
		return
	}

	if parsedUrl, err := url.Parse(urlString); err != nil || parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "bad request"})
		return
	}

	urlModel := h.URLService.Shorten(urlString)
	respBody := shortenResponse{
		Message:  "Will shorten json",
		URLModel: urlModel,
	}
	writeJSON(w, http.StatusOK, &respBody)
}

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_ = ctx
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok!")
}

func (h Handler) RootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Shortens your URL!")
}

func (h Handler) CodeHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}
	code := chi.URLParam(req, "code")
	url, isFound := h.URLService.Resolve(code)
	if isFound == false {
		writeJSON(w, http.StatusNotFound, errorResponse{Error: "not found"})
		return
	}
	http.Redirect(w, req, url.OriginalURL, http.StatusFound)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
