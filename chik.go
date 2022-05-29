package chick

import "net/http"

type Routes interface {
	Routes() []Route
	Middlewares() Middlewares
	Match(rctx *Context, method, path string) bool
}

type Middlewares []func(http.Handler) http.Handler
