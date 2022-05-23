package server

import (
	"os"

	"github.com/rs/zerolog/log"
)

// PrintEnvVariables print the ENV variables on debug mode only
func PrintEnvVariables() {
	log.Debug().Msg("ENV Variables: ")

	for _, value := range os.Environ() {
		log.Debug().Msg(value)
	}
}
