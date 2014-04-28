package atomic2

import (
	"testing"
)

func NewInt32(t *testing.T) {
	counter := NewAtomicInt32(42)
	count := counter.IncrementAndGet()
	if count != 43 {
		t.Errorf("wrong value %d", count)
	}

	count = counter.Get()
	if count != 43 {
		t.Errorf("wrong value %d", count)
	}

	swapped := counter.CompareAndSwap(42, 77)
	if swapped {
		t.Error("should be false")
	}

	swapped = counter.CompareAndSwap(43, 77)

	if !swapped {
		t.Error("should be true")
	}

	count = counter.Get()
	if count != 77 {
		t.Errorf("wrong value %d", count)
	}
}
