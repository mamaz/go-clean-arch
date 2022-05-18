package logger

import (
	"go-clean-arch/infrastructure/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const LOCAL = "local"
const DEVELOPMENT = "development"
const STAGING = "staging"
const PROD = "production"

func SetupWithEnv(environment string) {
	var _logger zerolog.Logger

	switch environment {
	case LOCAL:
		// on local, for debugging sake, the log format is set to plaintext
		_logger = log.
			With().
			Stack().
			Caller().
			Logger().
			Output(zerolog.ConsoleWriter{Out: os.Stdout}).
			Level(zerolog.DebugLevel)
	default:
		// otherwise is set to JSON on other environment
		_logger = log.
			With().
			Stack().
			Caller().
			Logger().
			Output(os.Stdout).
			Level(zerolog.InfoLevel)
	}

	log.Logger = _logger
}

func Disable() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func Enable() {
	// setting up env var with specific config will re-enable logging
	SetupWithEnv(config.CONFIG_ENV)
}
