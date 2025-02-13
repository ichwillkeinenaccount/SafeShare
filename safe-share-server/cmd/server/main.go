package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"server/internal/api"
)

func main() {
	// Parse command line flags
	configFile := flag.StringP("config", "c", "./config.yaml", "config file (default is ./config.yaml)")
	flag.Parse()

	// Read the configuration file
	readConfig(*configFile)

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), viper.GetString("database.url"))
	if err != nil {
		slog.Error(fmt.Sprintf("Error connecting to database: %s", err))
		os.Exit(1)
	}

	// Close the database connection when the program exits
	defer func(conn *pgx.Conn) {
		err := conn.Close(context.Background())
		if err != nil {
			slog.Error(fmt.Sprintf("Error closing database connection: %s", err))
			os.Exit(1)
		}
	}(conn)

	// Start the server
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
