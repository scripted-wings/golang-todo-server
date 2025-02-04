package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRouter() *gin.Engine {
	ginMode := viper.GetString("server.gin_mode")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	trustedProxies := viper.GetStringSlice("server.trusted_proxies")
	if len(trustedProxies) > 0 {
		router.SetTrustedProxies(trustedProxies)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("cors.allow_origins"),
		AllowMethods:     viper.GetStringSlice("cors.allow_methods"),
		AllowHeaders:     viper.GetStringSlice("cors.allow_headers"),
		AllowCredentials: viper.GetBool("cors.allow_credentials"),
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	return router
}
