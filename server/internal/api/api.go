package api

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
)

func StartApi() {
	textHandler := &TextHandler{}

	router := initRoutes(textHandler)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: router,
	}

	slog.Info(fmt.Sprintf("Server running on port %d", viper.GetInt("server.port")))
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Could not start api", err)
		return
	}
}

func initRoutes(textHandler *TextHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/text/", textHandler.getAll)
	mux.HandleFunc("POST /api/v1/text/create", textHandler.postText)
	return mux
}
