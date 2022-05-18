package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
)

// This test is just for showing how to use logging with serilog

func TestSimpleLog(t *testing.T) {
	SetupWithEnv("local")

	log.Info().Msg("This is simple log")
}

func TestLogWithDuration(t *testing.T) {
	SetupWithEnv("local")

	log.Info().Dur("duration", time.Duration(200*time.Millisecond)).Msg("This is simple log with duration")
}

func TestErrorLogWith(t *testing.T) {
	SetupWithEnv("local")

	err := fmt.Errorf("some random error")

	log.Error().Err(err).Msg("an error has occurred")
}
