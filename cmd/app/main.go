package main

import (
	"go-pet-shop/internal/config"
	"go-pet-shop/internal/handlers"
	"go-pet-shop/internal/lib/logger"
	"go-pet-shop/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := config.MustLoad()

	// Settings logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting the project...", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
	log.Error("error messages are enabled")

	// Settings and started database
	storage, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Init router
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)     // Хороший middleware для логирования
	router.Use(middleware.Recoverer)     // Перехватывает паники и возвращает 500
	router.Use(middleware.URLFormat)     // Для красивых URL при подключении к обработчикам
	router.Use(logger.CustomLogger(log)) // Логирует все исходящие запросы

	// Handlers
	router.Get("/health", handlers.StatusHandler)
	router.Get("/products", handlers.GetAllProducts(log, storage))
	router.Post("/products", handlers.CreateProduct(log, storage))
	router.Delete("/products/{id}", handlers.DeleteProduct(log, storage))
	router.Put("/products/{id}", handlers.UpdateProduct(log, storage))

	// Оборачиваем роутер в middleware
	handler := logger.LoggingMiddleware(log, router)

	// Settings and started server
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	log.Info("Starting server on", slog.String("address", cfg.Address))

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("Server error: ", slog.String("err", err.Error()))
	}
}
