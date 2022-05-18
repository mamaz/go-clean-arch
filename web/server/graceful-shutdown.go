package server

import (
	"context"
	"fmt"
	"go-clean-arch/infrastructure/config"
	"go-clean-arch/infrastructure/config/secrets"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func startServerWithGracefulShutdown(r *chi.Mux) {
	addr := fmt.Sprintf("0.0.0.0:%d", secrets.APP_PORT)
	server := &http.Server{Addr: addr, Handler: r}

	// Create server context
	serverCtx, cancelServerCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		graceful_timeout := config.CONFIG_SERVER_GRACEFUL_SHUTDOWN_TIMEOUT_S

		shutDownCtx, _ := context.WithTimeout(serverCtx, time.Duration(graceful_timeout)*time.Second)

		go func() {
			<-shutDownCtx.Done()
			if shutDownCtx.Err() == context.DeadlineExceeded {
				log.Fatal().Msg("graceful shutdown timed out.. forcing exit")
			}
		}()

		log.Info().Msgf("Shutting down server with timeout: %v seconds", graceful_timeout)
		err := server.Shutdown(shutDownCtx)
		if err != nil {
			log.Fatal().Msgf("error on shutting down gracefully %+v", err)
		}

		cancelServerCtx()
	}()

	log.Info().Msgf("Starting server in port :%d", secrets.APP_PORT)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("error on starting up server: %+v", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	log.Info().Msgf("Server is shut down!")
}
