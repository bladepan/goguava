package atomic

import (
	"sync/atomic"
)

type atomicInt32 int32

func NewAtomicInt32(init int32) *atomicInt32 {
	i := atomicInt32(init)
	return &i
}

//Atomically increments by one the current value.
//Returns: the updated value
func (i *atomicInt32) IncrementAndGet() int32 {
	return i.AddAndGet(1)
}

//Atomically increments by delta the current value.
//Returns: the updated value
func (i *atomicInt32) AddAndGet(delta int32) int32 {
	address := (*int32)(i)
	swapped := false
	var oldValue int32
	var newValue int32
	for swapped != true {
		oldValue = atomic.LoadInt32(address)
		newValue = oldValue + delta
		swapped = atomic.CompareAndSwapInt32(address, oldValue, newValue)
	}
	return newValue
}

//Atomically sets the value to update if the current value equals to expect
//Returns: true if successful
func (i *atomicInt32) CompareAndSwap(expect int32, update int32) (swapped bool) {
	address := (*int32)(i)
	swapped = atomic.CompareAndSwapInt32(address, expect, update)
	return
}

func (c *atomicInt32) Get() int32 {
	address := (*int32)(c)
	return atomic.LoadInt32(address)
}
