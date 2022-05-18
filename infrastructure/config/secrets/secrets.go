package secrets

import "github.com/spf13/viper"

// Secrets are values that supposed to be secrets and private and not public consuming
// for example: keys, passwords, ip addresses

// You just need to add more env variables here
var POSTGRES_HOST string
var POSTGRES_PORT int
var POSTGRES_USERNAME string
var POSTGRES_PASSWORD string
var POSTGRES_DATABASE string
var APP_PORT int

func Reload() {
	APP_PORT = viper.GetInt("APP_PORT")
	POSTGRES_HOST = viper.GetString("POSTGRES_HOST")
	POSTGRES_PORT = viper.GetInt("POSTGRES_PORT")
	POSTGRES_USERNAME = viper.GetString("POSTGRES_USERNAME")
	POSTGRES_PASSWORD = viper.GetString("POSTGRES_PASSWORD")
	POSTGRES_DATABASE = viper.GetString("POSTGRES_DATABASE")
}
