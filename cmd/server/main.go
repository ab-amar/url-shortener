package main

import (
	"net/http"

	"github.com/ab-amar/url-shortener/internal/handler"
	"github.com/ab-amar/url-shortener/internal/config"
)

func main() {
	conf := config.NewConfig()

	port := conf.Port
	server := createServer(port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func createServer(port string) http.Server {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/health", handler.HealthHandler)
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return server
}
