package main

import (
	"net/http"
	"os"
	"os/signal"
	"context"
	"time"
	"github.com/ab-amar/url-shortener/internal/handler"
	"github.com/ab-amar/url-shortener/internal/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	port := conf.Port
	server := createServer(port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func createServer(port string) http.Server {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/shorten", handler.ShortenHandler)
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return server
}
