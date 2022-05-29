package chick

import "context"

type RouteParams struct {
	Keys, Values []string
}

func (s *RouteParams) Add(k, v string) {
	s.Keys = append(s.Keys, k)
	s.Values = append(s.Values, v)
}

type Context struct {
	Routes Routes

	parentCtx context.Context

	RoutePath   string
	RouteMethod string

	URLParams RouteParams

	routeParams RouteParams

	routePattern string

	RoutePatterns []string

	methodNotAllowed bool
}

func (x *Context) Reset() {
	x.Routes = nil
	x.RoutePath = ""
	x.RouteMethod = ""
	x.RoutePatterns = x.RoutePatterns[:0]
	x.URLParams.Keys = x.URLParams.Keys[:0]
	x.URLParams.Values = x.URLParams.Values[:0]

	x.routePattern = ""
	x.routeParams.Keys = x.routeParams.Keys[:0]
	x.routeParams.Values = x.routeParams.Values[:0]
	x.methodNotAllowed = false
	x.parentCtx = nil
}

func (x *Context) URLParam(key string) string {
	for k := len(x.URLParams.Keys) - 1; k >= 0; k-- {
		if x.URLParams.Keys[k] == key {
			return x.URLParams.Values[k]
		}
	}
	return ""
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "chik context key " + k.name
}
