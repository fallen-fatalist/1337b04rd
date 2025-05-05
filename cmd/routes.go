package cmd

import (
	httpserver "1337bo4rd/internal/adapter/handler/http"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc()

	// Logging middleware applied
	middlewareAppliedMux := httpserver.RequestLoggingMiddleware(mux)
	middlewareAppliedMux = httpserver.HeadersMiddleware(middlewareAppliedMux)

	return middlewareAppliedMux
}
