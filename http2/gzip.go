package http2

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipHandler struct {
	handler http.Handler
}

func NewGzipHandler(handler http.Handler) *GzipHandler {
	gzipHandler := &GzipHandler{
		handler: handler,
	}
	return gzipHandler
}

// ResponseWriter is required for Header/WriteHeader methods
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// Writer and ResponseWriter both have Write method, make sure we invoke Write on
// the gzip wrapped writer
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (g *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		g.handler.ServeHTTP(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	g.handler.ServeHTTP(gzr, r)
}

// from https://gist.github.com/the42/1956518
func MakeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r)
	}
}
