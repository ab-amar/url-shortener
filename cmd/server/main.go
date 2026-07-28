package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ab-amar/url-shortener/internal/config"
	"github.com/ab-amar/url-shortener/internal/handler"
	"github.com/ab-amar/url-shortener/internal/metrics"
	"github.com/ab-amar/url-shortener/internal/middleware"
	"github.com/ab-amar/url-shortener/internal/repository"
	"github.com/ab-amar/url-shortener/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	port := conf.Port
	databaseURL := conf.DatabaseURL
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		slog.Error("db failed", "error", err)
		panic(err)
	}
	appMetrics := metrics.New()

	pgRepo := repository.NewPostgresRepository(pool)
	var postgresRepository repository.URLRepository = &pgRepo
	var shortenerService service.URLService = service.ShortenerService{
		URLRepo: postgresRepository,
	}
	var h handler.Handler = handler.New(shortenerService, appMetrics)
	server := createServer(port, h, appMetrics)
	defer pool.Close()
	slog.Info("server starting", "port", port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	slog.Info("shutdown signal received")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
		panic(err)
	}
	slog.Info("server stopped")
}

func createServer(port string, h handler.Handler, appMetrics *metrics.Metrics) http.Server {

	router := chi.NewRouter()
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.AppHeaderMiddleware)
	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.MetricsMiddleware(appMetrics))
	router.Use(middleware.LoggingMiddleware)
	router.Get("/", h.RootHandler)
	router.Get("/{code}", h.CodeHandler)
	router.Get("/health", handler.HealthHandler)
	router.Get("/ready", handler.ReadyHandler)
	router.Post("/shorten", h.ShortenHandler)
	server := http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	return server
}
