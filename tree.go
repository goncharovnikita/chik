package chick

import "net/http"

type Route struct {
	SubRoutes Routes
	Handlers  map[string]http.Handler
	Pattern   string
}
