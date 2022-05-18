package config

import (
	"github.com/spf13/viper"
)

// Configuration is a collection of values that can be set to configure the behaviour of the app
// For example: rate limiting values, timeout values, feature flags, etc

var CONFIG_SERVER_GRACEFUL_SHUTDOWN_TIMEOUT_S int

// this will reload config either from file or from system's ENV
// see: infrastructure/configuration/setup.go
func reloadConfig() {
	CONFIG_SERVER_GRACEFUL_SHUTDOWN_TIMEOUT_S = viper.GetInt("CONFIG_SERVER_GRACEFUL_SHUTDOWN_TIMEOUT_S")
}
