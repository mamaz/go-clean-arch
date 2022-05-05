package configuration

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

// SetUpConfig will set config from .env if it's exist
// otherwise it will set from system's ENV variables
// filename should be and env file: .env or .env.* file
// dirpath should be in this format: /some/dirpath
func SetUpConfig(dirpath string, filename string) {
	fileExist := isFileExist(dirpath + "/" + filename)

	if fileExist {
		viper.SetConfigFile(filename)
		viper.AddConfigPath(dirpath)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading config file: %+v", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	// reload config
	reloadConfig()
}

func isFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
