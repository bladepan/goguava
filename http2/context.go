package http2

import (
	"log"
	"net/http"
)

type Context interface {
	HandleNext(http.ResponseWriter, *http.Request)
	// by convension 'route' is the matched route pattern for the context,
	// 'error' is error occured in the request
	Param(interface{}) interface{}
	SetParam(interface{}, interface{}) interface{}
	Route() Route
}

type ContextHandlerFunc func(*Context, http.ResponseWriter, *http.Request)

func NoopContextHandlerFunc(context *Context, w http.ResponseWriter, r *http.Request) {
	context.HandleNext(w, r)
}

type DefaultContext struct {
	route    Route
	params   map[interface{}]interface{}
	handlers []ContextHandlerFunc
	index    int
}

func (d *DefaultContext) HandleNext(w ResponseWriter, r *Request) {
	if index < len(d.handlers) {
		handler := d.handlers[index]
		d.index++
		handler(d, w, r)
	} else {
		//do nothing
	}
}

func (d *DefaultContext) Param(key interface{}) interface{} {
	return d.params[key]
}

func (d *DefaultContext) SetParam(key interface{}, val interface{}) interface{} {
	d.params[key] = val
}

func (d *DefaultContext) Route() Route {

}

type ContextMux interface {
	Middleware(ContextHandlerFunc)
	HandlerFunc(string, ContextHandlerFunc)
}

type DefaultContextMux struct {
	middlewares []ContextHandlerFunc
	// the most specific route is in front
	routes     []Route
	handlerMap map[Route]ContextHandlerFunc
}

func (c *ContextMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := new(DefaultContext)
	var matched Route
	for _, route := range c.routes {
		if route.Match(r) {
			matched = route
			break
		}
	}
	if matched == nil {
		//TODO redirect to error page
		errorMsg := "cannot find handler for request " + r.RequestURI
		log.Println(errorMsg)

		http.Error(w, errorMsg, http.StatusNotFound)
		return
	}
	context.route = matched
	context.handlers = make([]ContextHandlerFunc, len(c.middlewares)+1)
	for _, middleware := range c.middlewares {
		context.handlers = append(context.handlers, middleware)
	}
	handler := handlerMap[matched]
	context.handlers = append(context.handlers, handler)
	context.HandleNext(w, r)
}

func NewDefaultContextMux() *DefaultContextMux {
	result := &DefaultContextMux{
		routes:      make([]Route, 0, 16),
		middlewares: make([]ContextHandlerFunc, 0, 16),
		handlerMap:  make(map[Route]ContextHandlerFunc),
	}
	return result
}

func (d *DefaultContextMux) Middleware(handler ContextHandlerFunc) {

}

func (d *DefaultContextMux) HandleFunc(path string, handler ContextHandlerFunc) {

}
