package router

import (
	"getir-case/internal/handler/fetch"
	"getir-case/internal/handler/inmemory"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes() *httprouter.Router {
	router := httprouter.New()

	router.POST("/fetch", fetch.Post)

	router.POST("/in-memory", inmemory.Post)
	router.GET("/in-memory", inmemory.Get)

	return router
}
