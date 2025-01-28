package api

import (
	"fmt"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

// StartApi starts the API server
//
//	@title			SafeShare API
//	@version		1.0
//	@description	This is the API for SafeShare, a secure text sharing service.
//	@termsOfService	https://github.com/ichwillkeinenaccount/SafeShare
//	@contact.name	SafeShare GitHub
//	@contact.url	https://github.com/ichwillkeinenaccount/SafeShare
//	@license.name	All Rights Reserved
//	@license.url	https://en.wikipedia.org/wiki/All_rights_reserved
//	@server			http://localhost:8080
func StartApi() {
	textHandler := &TextHandler{}

	router := initRoutes(textHandler)
	routerWithMiddlewares := chainMiddlewares(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: routerWithMiddlewares,
	}

	slog.Info(fmt.Sprintf("Server running on port %d", viper.GetInt("server.port")))
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Could not start api", "error", err)
		return
	}
}

func initRoutes(textHandler *TextHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /api/v1/", httpSwagger.Handler(httpSwagger.URL("docs/swagger.yaml")))
	mux.Handle("GET /api/v1/docs/", http.StripPrefix("/api/v1/docs", http.FileServer(http.Dir("internal/api/docs"))))
	mux.HandleFunc("GET /api/v1/text/", textHandler.getAll)
	mux.HandleFunc("POST /api/v1/text/create", textHandler.postText)
	return mux
}
