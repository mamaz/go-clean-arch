package server

import (
	"os"

	"github.com/rs/zerolog/log"
)

func PrintEnvVariables() {
	log.Debug().Msg("ENV Variables: ")

	for _, value := range os.Environ() {
		log.Debug().Msg(value)
	}
}
