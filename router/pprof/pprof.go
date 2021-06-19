package pprof

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/pprof"
)

func Register(e *echo.Echo) {
	g := e.Group("/debug/pprof")
	g.GET("/", handler(pprof.Index))
	g.GET("/heap", handler(pprof.Handler("heap").ServeHTTP))
	g.GET("/goroutine", handler(pprof.Handler("goroutine").ServeHTTP))
	g.GET("/block", handler(pprof.Handler("block").ServeHTTP))
	g.GET("/mutex", handler(pprof.Handler("mutex").ServeHTTP))
	g.GET("/allocs", handler(pprof.Handler("allocs").ServeHTTP))
	g.GET("/threadcreate", handler(pprof.Handler("threadcreate").ServeHTTP))
	g.GET("/cmdline", handler(pprof.Cmdline))
	g.GET("/profile", handler(pprof.Profile))
	g.GET("/symbol", handler(pprof.Symbol))
	g.GET("/trace", handler(pprof.Trace))
}

func handler(h http.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		h.ServeHTTP(context.Response().Writer, context.Request())
		return nil
	}
}
