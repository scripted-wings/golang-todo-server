package server

import (
	"github.com/scripted-wings/todo-server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func StartServer() {
	router := SetupRouter()

	port := viper.GetString("server.port")
	config.Logger.Info("Starting server...", zap.String("port", port))

	if err := router.Run(":" + port); err != nil {
		config.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
