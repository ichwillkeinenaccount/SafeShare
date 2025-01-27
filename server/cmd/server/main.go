package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
)

func main() {
	initConfig()

	appName := viper.GetString("app_name")
	debug := viper.GetBool("debug")
	port := viper.GetInt("server.port")
	slog.Info(fmt.Sprintf("%s started", appName))
	slog.Info(fmt.Sprintf("Debug enabled: %t", debug))
	slog.Info(fmt.Sprintf("Server running in port: %d", port))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
