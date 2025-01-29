package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"server/internal/api"
)

func main() {
	initConfig()
	slog.Info(fmt.Sprintf("Starting %s", viper.GetString("app_name")))
	if viper.GetBool("debug") {
		slog.Info("Debug enabled")
	}

	api.StartServer()

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
