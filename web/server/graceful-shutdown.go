package server

import (
	"context"
	"fmt"
	"go-clean-arch/infrastructure/configuration"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func startServerWithGracefulShutdown(r *chi.Mux) {
	addr := fmt.Sprintf("0.0.0.0:%d", configuration.APP_PORT)
	server := &http.Server{Addr: addr, Handler: r}

	// Create server context
	serverCtx, cancelServerCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		graceful_timeout := configuration.CONFIG_SERVER_GRACEFUL_SHUTDOWN_TIMEOUT_S

		shutDownCtx, _ := context.WithTimeout(serverCtx, time.Duration(graceful_timeout)*time.Second)

		go func() {
			<-shutDownCtx.Done()
			if shutDownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit")
			}
		}()

		log.Printf("Shutting down server with timeout: %v seconds", graceful_timeout)
		err := server.Shutdown(shutDownCtx)
		if err != nil {
			log.Fatal("error on shutting down gracefully", err)
		}

		cancelServerCtx()
	}()

	log.Printf("Starting server in port :%d", configuration.APP_PORT)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("error on starting up server", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	log.Printf("Server is shut down!")
}
