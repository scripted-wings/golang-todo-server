package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort string `mapstructure:"server_port"`
}

var Config AppConfig

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
