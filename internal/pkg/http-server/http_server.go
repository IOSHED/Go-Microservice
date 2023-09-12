package http_server

import (
	"micriservice/internal/config"
)

type HttpServer struct {
	Cfg *config.Config
}

type IHttpServer[T any, H any] interface {
	HttpReqest[T, H]
	UsingMiddleware[T]
	Runer
}

type Runer interface {
	Run()
}

type HttpReqest[T any, H any] interface {
	Post(path string, handler H, m ...T)
}

type UsingMiddleware[T any] interface {
	UseMiddleware(m ...T)
}
