// Package secrets are values that supposed to be secrets and private and not public consuming
// for example: keys, passwords, ip addresses
//
// To differentiate with config, we place secrets on a different package, so it will be called from outside like this:
// `secrets.POSTGRES_HOST` for instance
// You just need to add more env variables here
package secrets

import "github.com/spf13/viper"

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
