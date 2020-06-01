package config

import (
	"github.com/spf13/viper"
)

// Config - application config struct
type Config struct {
	Port    string `dotenv:"PORT"`
	GoEnv   string `dotenv:"GO_ENV"`
	Version string `dotenv:"VERSION"`
}

// NewConfig - creates the application config struct
func NewConfig() *Config {
	var config Config

	// Setting defaults
	if viper.Get("GO_ENV") == nil {
		viper.SetDefault("GO_ENV", "local")
	}

	if viper.Get("VERSION") == nil {
		viper.SetDefault("VERSION", "local")
	}

	viper.Unmarshal(&config)

	return &config
}
