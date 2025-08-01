package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read configuration file: %s", err)
	}
}

const (
	Port    = "port"
	BaseKey = "base_key"
	KDFSalt = "kdf_salt"
)
