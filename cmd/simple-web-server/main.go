package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AFIF-ZILANI/simple-web-server/pkg/config"
	"github.com/AFIF-ZILANI/simple-web-server/pkg/http/handlers/student"
)

func main() {

	cfg := config.MustLoadConfig()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server is running on", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			slog.Error("Failed to start server", "error", err)
		}
	}()

	<-done

	slog.Info("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to gracefully shutdown the server", "error", err)
	} else {
		slog.Info("Server stopped gracefully")
	}

	slog.Info("Server shutdown successfully")
}
