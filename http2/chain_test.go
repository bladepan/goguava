package http2

import (
	"net/http"
	"testing"
)

type dumbHttpHandler struct {
	count int
}

func (d *dumbHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.count++
}

func TestChain(t *testing.T) {
	dumbHandler1 := &dumbHttpHandler{}
	chainableDumbHandler1 := AsChainableHandler(dumbHandler1.ServeHTTP)
	chained := ChainHandlers(chainableDumbHandler1, chainableDumbHandler1, chainableDumbHandler1)
	chained(nil, nil)
	if dumbHandler1.count != 3 {
		t.Fatalf("count should be 3 %d \n", dumbHandler1.count)
	}
	//reset count
	dumbHandler1.count = 0
	dumbHandler2 := &dumbHttpHandler{}
	chainableDumbHandler2 := AsChainableHandler(dumbHandler2.ServeHTTP)
	chained = ChainHandlers(chainableDumbHandler1, chainableDumbHandler2)
	chained(nil, nil)
	if dumbHandler2.count != 1 {
		t.Fatalf("count should be 1 %d \n", dumbHandler2.count)
	}
}
