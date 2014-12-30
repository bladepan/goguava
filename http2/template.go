package http2

import (
	"html/template"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

type TemplateFileHandler struct {
	RootDir  string
	DataFunc func(*http.Request) interface{}
}

// we must buffer everything before calling template functions
func (h *TemplateFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templateFileName := r.URL.Path
	if templateFileName == "" || strings.HasSuffix(templateFileName, "/") {
		templateFileName += "index.html"
	}
	log.Println("fetch " + templateFileName)
	realPath := path.Join(h.RootDir, templateFileName)
	// caching
	templ, err := template.ParseFiles(realPath)
	// should have the option of specify an error page
	if err != nil {
		log.Printf("error when parse template %#v \n", err)
		http.Error(w, "some thing wrong", http.StatusInternalServerError)
		return
	}
	data := h.DataFunc(r)
	contentType := mime.TypeByExtension(filepath.Ext(templateFileName))
	// should use some constant or sniff the type
	w.Header().Set("Content-Type", contentType)
	templ.Execute(w, &data)
}
