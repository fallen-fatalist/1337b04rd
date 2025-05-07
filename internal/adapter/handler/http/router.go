package httpserver

import (
	"net/http"
)

func NewRouter(
	postHandler PostHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", postHandler.HandleCatalog)
	mux.HandleFunc("/archive", postHandler.HandleArchive)

	// Logging middleware applied
	middlewareAppliedMux := RequestLoggingMiddleware(mux)
	middlewareAppliedMux = HeadersMiddleware(middlewareAppliedMux)

	return middlewareAppliedMux
}
