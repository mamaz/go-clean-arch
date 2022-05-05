package configuration

import (
	"github.com/spf13/viper"
)

// You just need to add more env variables here
var POSTGRES_HOST string
var POSTGRES_PORT int
var POSTGRES_USERNAME string
var POSTGRES_PASSWORD string
var POSTGRES_DATABASE string

// this will reload config either from file or from system's ENV
// see: infrastructure/configuration/setup.go
func reloadConfig() {
	POSTGRES_HOST = viper.GetString("POSTGRES_HOST")
	POSTGRES_PORT = viper.GetInt("POSTGRES_PORT")
	POSTGRES_USERNAME = viper.GetString("POSTGRES_USERNAME")
	POSTGRES_PASSWORD = viper.GetString("POSTGRES_PASSWORD")
	POSTGRES_DATABASE = viper.GetString("POSTGRES_DATABASE")
}
