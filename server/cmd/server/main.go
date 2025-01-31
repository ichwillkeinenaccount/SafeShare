package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"server/internal/api"
)

func main() {
	slog.Info(fmt.Sprintf("Starting %s", viper.GetString("app_name")))
	if viper.GetBool("debug") {
		slog.Info("Debug enabled")
	}
	readConfig("config", "yaml", ".")
	api.StartServer()
}

// readConfig reads the configuration file.
//
// Parameters:
//
//	configName (string): The name of the configuration file.
//	configType (string): The type of the configuration file.
//	configPath (string): The path to the configuration file.
//
// Panics:
//
//	If there is a fatal error reading the configuration file.
func readConfig(configName string, configType string, configPath string) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
