package http2

import "net/http"
import "context"

// MiddleWare define a chain of handler functions
type MiddleWare struct {
	handlers []HandleFunc
}

//HandleFunc return true to indicate invoke next handleFunc in the chain
type HandleFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, bool)

// HandleFunc implements http.HandleFunc
func (mw *MiddleWare) HandleFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	next := false
	for _, h := range mw.handlers {
		ctx, next = h(ctx, w, r)
		if !next {
			return
		}
	}
}

// NewMiddleWare create a new middleware
func NewMiddleWare(handlers ...HandleFunc) *MiddleWare {
	w := &MiddleWare{
		handlers: handlers,
	}
	return w
}

// NewHandleFunc create a default implementation of HandleFunc
func NewHandleFunc(handleFunc func(w http.ResponseWriter, r *http.Request)) HandleFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, bool) {
		handleFunc(w, r)
		return ctx, true
	}
}
