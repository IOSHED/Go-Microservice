package http_server_echo

import (
	"micriservice/internal/config"
	http_server "micriservice/internal/pkg/http-server"

	"github.com/labstack/echo/v4"
)

type HttpServerEcho struct {
	Echo *echo.Echo
	http_server.HttpServer
}

func New(cfg *config.Config) *HttpServerEcho {
	e := echo.New()

	e.Server.ReadTimeout = cfg.Timeout
	e.Server.WriteTimeout = cfg.Timeout
	e.Server.IdleTimeout = cfg.IdleTimeout

	return &HttpServerEcho{
		Echo:       e,
		HttpServer: http_server.HttpServer{Cfg: cfg},
	}
}

func (h *HttpServerEcho) Post(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	h.Echo.POST(path, handler, m...)
}

func (h *HttpServerEcho) UseMiddleware(m ...echo.MiddlewareFunc) {
	h.Echo.Use(m...)
}

func (h *HttpServerEcho) Run() {
	h.Echo.Start(h.Cfg.Address)
}
