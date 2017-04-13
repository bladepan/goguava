package http2

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

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

// inspired by https://gist.github.com/the42/1956518
func GzipHandler(w http.ResponseWriter, r *http.Request, inner http.HandlerFunc) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		inner(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	inner(gzr, r)
}
