package api

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
	"server/internal/middleware"
)

// StartServer starts the API server
//
//	@title			SafeShare API
//	@version		1.0
//	@description	This is the API for SafeShare, a secure text sharing service.
//	@termsOfService	https://github.com/ichwillkeinenaccount/SafeShare
//	@contact.name	SafeShare GitHub
//	@contact.url	https://github.com/ichwillkeinenaccount/SafeShare
//	@license.name	All Rights Reserved
//	@license.url	https://en.wikipedia.org/wiki/All_rights_reserved
func StartServer() {
	router := loadRoutes()
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Authentication,
		middleware.AllowCors,
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: stack(router),
	}

	slog.Info(fmt.Sprintf("Server running on port %d", viper.GetInt("server.port")))
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Could not start api", "error", err)
		return
	}
}

func loadRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("GET /docs/swagger.yaml", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/api/docs/swagger.yaml")
	}))
	router.Handle("GET /docs/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/api/docs/swagger.json")
	}))

	textHandler := &TextHandler{}
	router.HandleFunc("GET /text", textHandler.getAll)
	router.HandleFunc("POST /text/create", textHandler.postText)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))
	return v1
}
