package http2

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestData struct {
	A string `json:"a,omitempty"`
}

func Test(t *testing.T) {
	handler := &TemplateFileHandler{
		RootDir: "testdata",
	}
	myData := &TestData{
		A: "val",
	}
	handler.DataFunc = func(*http.Request) interface{} {
		return myData
	}
	respWriter := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://example.com/", nil)
	handler.ServeHTTP(respWriter, request)
	respBody := respWriter.Body.String()

	if !strings.Contains(respBody, "val") {
		t.Fatal("should contain val")
	}
	contentType := respWriter.Header().Get(HttpHeaderContentType)
	fmt.Println(contentType)

	fmt.Println(respBody)
}
