package server

import (
	"log"
	"os"
)

func PrintEnvVariables() {
	log.Println("ENV Variables: ")

	for _, value := range os.Environ() {
		log.Println(value)
	}
}
