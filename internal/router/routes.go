package router

import (
	"getir-case/internal/handler"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes() *httprouter.Router {
	router := httprouter.New()
	router.POST("/fetch", handler.Fetch)
	router.POST("/in-memory", handler.InMemory)

	return router
}
