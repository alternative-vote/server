package generated

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

var middleware = []func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc){}

type Handler func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc)

func AddMiddleware(handler Handler, args ...string) {
	uriPattern := getUriPattern(args)
	methods := getMethods(args)
	matchyHandler := func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		match, vars := isMatch(uriPattern, methods, req)
		if match {
			req = contextSet(req, varsKey, vars)
			handler(res, req, next)
			return
		}
		next(res, req)
	}

	middleware = append(middleware, matchyHandler)
}

func getUriPattern(args []string) string {
	if len(args) < 1 {
		return "/{uri:.*}" //this is how you match all routes with gorilla mux
	}
	return args[0]
}

func getMethods(args []string) []string {
	if len(args) < 2 {
		return []string{"GET", "PUT", "POST", "DELETE", "HEAD", "TRACE", "CONNECT", "PATCH", "COPY", "LINK", "UNLINK", "PURGE", "LOCK", "UNLOCK", "PROPFIND", "VIEW"}
	}
	return args[1:]

}

func isMatch(pattern string, methods []string, req *http.Request) (bool, map[string]string) {
	rm := new(mux.RouteMatch)
	route := new(mux.Route)
	route = route.Path(pattern)
	if len(methods) > 0 {
		route = route.Methods(methods...)
	}
	rm.Route = route
	match := route.Match(req, rm)
	return match, rm.Vars
}

//The following mostly coppied from gorilla/mux (https://github.com/gorilla/mux/blob/master/mux.go)
//gorilla mux does not provide an API to pull params from a uri and set it to
//the context outside of setting up a handler, so we're manually doing it
type contextKey int

const varsKey contextKey = iota

// Vars returns the route variables for the current request, if any.
var Vars = func(r *http.Request) map[string]string {
	if rv := contextGet(r, varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

func contextGet(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

func contextSet(r *http.Request, key, val interface{}) *http.Request {
	if val == nil {
		return r
	}

	return r.WithContext(context.WithValue(r.Context(), key, val))
}
