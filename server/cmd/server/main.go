package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log/slog"
	"server/internal/api"
)

func main() {
	configFile := flag.StringP("config", "c", "./config.yaml", "config file (default is ./config.yaml)")
	flag.Parse()

	readConfig(*configFile)

	slog.Info(fmt.Sprintf("Starting %s", viper.GetString("app_name")))
	if viper.GetBool("debug") {
		slog.Info("Debug enabled")
	}
	api.StartServer()
}

// readConfig reads the configuration file.
//
// Parameters:
//
//	configFile: The name and path of the configuration file.
//
// Panics:
//
//	If there is a fatal error reading the configuration file.
func readConfig(configFile string) {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
