package main

import (
	"net/http"

	"github.com/ab-amar/url-shortener/internal/handler"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthHandler)
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
