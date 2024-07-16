package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"backend/config"
	"backend/internal/application"
	"backend/internal/infrastructure/database/postgres"
	"backend/internal/interfaces/http/handlers"
	"backend/internal/interfaces/http/middleware"
	"backend/internal/interfaces/http/router"
	"backend/internal/logging"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger := logging.NewLogger()

	// Connect to database
	db, err := postgres.NewDB(cfg.DatabaseURL)
	if err != nil {
		logger.ErrorLog.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	refreshTokenRepo := postgres.NewRefreshTokenRepository(db)

	// Initialize services
	userService := application.NewUserService(userRepo)
	authService := application.NewAuthService(userRepo, refreshTokenRepo, cfg)



	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService, logger)
	authHandler := handlers.NewAuthHandler(authService, logger)

	// Set up router
	r := router.NewRouter(userHandler, authHandler)

	// Add middleware
	r.Use(middleware.ErrorMiddleware(logger))


	r.HandleFunc("/api/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/users", middleware.AuthMiddleware(userHandler.ListUsers)).Methods("GET")
	r.HandleFunc("/api/users", middleware.AuthMiddleware(userHandler.CreateUser)).Methods("POST")
	r.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(userHandler.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/api/users/search", middleware.AuthMiddleware(userHandler.SearchUsers)).Methods("GET")

	// Start server with graceful shutdown
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		logger.InfoLog.Printf("Server starting on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorLog.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.InfoLog.Println("Shutting down server...")

	if err := server.Close(); err != nil {
		logger.ErrorLog.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.InfoLog.Println("Server exiting")
}