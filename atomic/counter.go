package atomic

import (
	"sync/atomic"
)

type counter32 int32

func NewCounter32(init int32) *counter32 {
	counter := counter32(init)
	return &counter
}

func (c *counter32) IncrementAndGet() int32 {
	address := (*int32)(c)
	swapped := false
	var oldValue int32
	var newValue int32
	for swapped != true {
		oldValue = atomic.LoadInt32(address)
		newValue = oldValue + 1
		swapped = atomic.CompareAndSwapInt32(address, oldValue, newValue)
	}
	return newValue
}

func (c *counter32) Get() int32 {
	return int32(*c)
}
