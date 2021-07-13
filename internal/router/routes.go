package router

import (
	"getir-case/internal/handler/fetch"
	"getir-case/internal/handler/inmemory"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	
	router.Handle("/fetch", fetch.NewHandler())

	router.Handle("/in-memory", inmemory.NewHandler())

	return router
}
