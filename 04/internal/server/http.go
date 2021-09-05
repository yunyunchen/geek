package server

import (
	"context"
	"geek/04/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Option func(ctx gin.IRouter)

type HttpServer struct {
	http.Server
	options []Option
}

func NewHttpServer(config *configs.Config) *HttpServer {
	router := gin.Default()

	httpServer := &HttpServer{}
	httpServer.Addr = config.ServerHttpAddr
	httpServer.Handler = router

	return httpServer
}

func (srv *HttpServer) AddRouters(opts ...Option) {
	srv.options = append(srv.options, opts...)
}

func (srv *HttpServer) InitAndStart() error {
	// init routers
	for _, opt := range srv.options {
		opt(srv.Handler.(gin.IRouter))
	}

	// start server
	return srv.ListenAndServe()
}

func (srv *HttpServer) Shutdown(ctx context.Context) error {
	return srv.Shutdown(ctx)
}
