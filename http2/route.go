package http2

import (
	"net/http"
	"string"
)

const (
	prefixMatch iota
	exactMatch
	matchAll
)

// inspired by gorilla/mux
type Route struct {
	Name      string
	matchType int
}

func (route *Route) Match(r *http.Request) bool {
	path := r.URL.Path
	switch route.matchType {
	case matchAll:
		return true
	case prefixMatch:
		return strings.HasPrefix(path, route.Name)
	default:
		return path == route.Name
	}
}

func Compile(string s) *Route {
	s = strings.TrimSpace(s)
	route := &Route{
		Name:      s,
		matchType: exactMatch,
	}
	switch {
	case s == "" || s == "/":
		route.matchType = matchAll
	case strings.HasSuffix(s, "/"):
		route.matchType = prefixMatch
	}
	return route
}
