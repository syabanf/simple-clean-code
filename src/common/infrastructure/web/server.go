package web

import (
	"context"

	"github.com/fasthttp/router"
)

type Server interface {
	Listen()
	Shutdown(context.Context)
}

// NewWebServer
func NewWebServer(router *router.Router, port int) (Server, error) {
	return newFastHTTPServer(router, port), nil
}

// NewRouter
func NewRouter() *router.Router {
	router := router.New()
	return router
}
