package web

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
)

var (
	corsAllowHeaders = []string{"*"}
	corsAllowMethods = []string{
		fasthttp.MethodHead,
		fasthttp.MethodGet,
		fasthttp.MethodPost,
		fasthttp.MethodPut,
		fasthttp.MethodPatch,
		fasthttp.MethodDelete,
	}
	corsAllowOrigin = []string{"*", "http://localhost:3000"}
)

type fastHTTPServer struct {
	server *fasthttp.Server
	router *router.Router
	port   int
}

func newFastHTTPServer(router *router.Router, port int) fastHTTPServer {
	f := fastHTTPServer{router: router, port: port}

	// Set CORS
	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   corsAllowOrigin,
		AllowedHeaders:   corsAllowHeaders,
		AllowedMethods:   corsAllowMethods,
		AllowCredentials: false,
	})

	f.server = &fasthttp.Server{
		Handler:              withCors.CorsMiddleware(fasthttp.CompressHandler(f.router.Handler)),
		ReadTimeout:          5 * time.Second,
		WriteTimeout:         5 * time.Second,
		MaxConnsPerIP:        50000,
		MaxRequestsPerConn:   50000,
		MaxKeepaliveDuration: 5 * time.Second,
	}

	return f
}

// Shutdown
func (f fastHTTPServer) Shutdown(ctx context.Context) {
	f.server.Shutdown()
}

// Listen
// Do not use *FastHTTPServer
func (f fastHTTPServer) Listen() {
	log.Printf("Web server started on : %v\n", f.port)
	f.server.ListenAndServe(fmt.Sprintf(":%v", f.port))
}
