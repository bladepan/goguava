package http2

import (
	"log"
	"net/http"
)

type ChainableHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request, http.HandlerFunc)
}

func NoopHandlerFunc(http.ResponseWriter, *http.Request) {
}

func WrapHandlerFunc(c ChainableHandler, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.ServeHTTP(w, r, f)
	}
}

func AsHandlerFunc(c ChainableHandler) http.HandlerFunc {
	return WrapHandlerFunc(c, NoopHandlerFunc)
}

type simpleChainableHandler struct {
	impl http.HandlerFunc
}

func (h *simpleChainableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, f http.HandlerFunc) {
	h.impl(w, r)
	f(w, r)
}

// convert a handlerFunc to a chainable handler
func AsChainableHandler(f http.HandlerFunc) ChainableHandler {
	handler := &simpleChainableHandler{
		impl: f,
	}
	return handler
}

type handlerChain struct {
	index    int
	handlers []ChainableHandler
}

func (h *handlerChain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerLen := len(h.handlers)
	switch {
	case h.index == handlerLen-1:
		h.handlers[h.index].ServeHTTP(w, r, NoopHandlerFunc)
	default:
		currentHandler := h.handlers[h.index]
		h.index++
		currentHandler.ServeHTTP(w, r, h.ServeHTTP)
	}
}

func ChainHandlers(handlers ...ChainableHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chain := &handlerChain{
			index:    0,
			handlers: handlers,
		}
		chain.ServeHTTP(w, r)
	}
}
