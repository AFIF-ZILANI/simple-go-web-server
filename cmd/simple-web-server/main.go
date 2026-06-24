package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AFIF-ZILANI/simple-web-server/pkg/config"
)

func main() {

	cfg := config.MustLoadConfig()

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
		fmt.Fprintf(w, "Hello, World! This is a simple web server.")
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("Server is running on %s\n", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			fmt.Printf("Failed to start server: %s\n", err.Error())
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
