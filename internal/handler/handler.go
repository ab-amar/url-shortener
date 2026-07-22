package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Message string `json:"message"`
}


func ShortenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w,"method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(req.Body)
	var reqBody shortenRequest
	if err := dec.Decode(&reqBody); err != nil {
		http.Error(w,"Bad request", http.StatusBadRequest)
		return
	}

	respBody := shortenResponse{
		Message: "Will shorten json",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&respBody); err != nil {
		http.Error(w,"Bad response", http.StatusInternalServerError)
		return
	}
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

func RootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Shortens your URL!")
}
