package http2

import (
	"fmt"
	"mime"
	"testing"
)

func TestMime(t *testing.T) {
	jsType := mime.TypeByExtension(".js")
	fmt.Println(jsType)
}
