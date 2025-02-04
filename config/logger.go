package config

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction(
		zap.Fields(zap.String("type", "todo-server")),
		zap.AddCaller(),
	)
	if err != nil {
		panic("Failed to initialize logger")
	}
}
