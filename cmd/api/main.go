package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"file_storage/internal/config"
	"file_storage/internal/handler"
	"file_storage/internal/infrastructure/storage"
	filerepository "file_storage/internal/repository/file"
	"file_storage/internal/router"
	fileservice "file_storage/internal/service/file"
	"file_storage/pkg/database"

	"gorm.io/gorm/logger"
)

func main() {
	cfg := config.Load()

	db, err := database.New(&database.Config{
		Path:        cfg.DatabaseURL,
		LogLevel:    logger.Info,
		AutoMigrate: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	s3Storage, err := storage.NewS3Storage(cfg.S3Bucket, cfg.S3Region)
	if err != nil {
		log.Fatalf("Failed to initialize S3 storage: %v", err)
	}

	repo := filerepository.New(db.DB)
	svc := fileservice.NewService(repo, s3Storage, cfg)
	h := handler.NewHandler(svc)
	r := router.SetupRouter(h)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Starting server on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	sqlDB, err := db.SqlDB()
	if err == nil {
		sqlDB.Close()
	}

	log.Println("Server exited gracefully")
}
