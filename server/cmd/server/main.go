package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"server/internal/api"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "server",
		Short: "A simple server",
		Long:  "A simple server that demonstrates how to use Cobra and Viper",
		Run: func(cmd *cobra.Command, args []string) {
			configFile := flag.String("config", "config.yaml", "Path to the configuration file")
			flag.Parse()
			slog.Info(fmt.Sprintf("configFile: %s", *configFile))

			readConfig(*configFile)
			slog.Info(fmt.Sprintf("Starting %s", viper.GetString("app_name")))
			if viper.GetBool("debug") {
				slog.Info("Debug enabled")
			}
			api.StartServer()
		},
	}

	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")

	err := rootCmd.Execute()
	if err != nil {
		slog.Error("error executing cobra command", "error", err)
		return
	}
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
