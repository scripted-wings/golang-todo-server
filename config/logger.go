package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	ginMode := viper.GetString("server.gin_mode")

	if ginMode == "release" {
		Logger, err = zap.NewProduction()
	} else {
		Logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("Failed to initialize logger")
	}

	zap.ReplaceGlobals(Logger)
}
