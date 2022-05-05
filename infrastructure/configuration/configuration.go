package configuration

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
)
