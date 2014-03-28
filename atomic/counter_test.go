package atomic

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := NewCounter32(42)
	count := counter.IncrementAndGet()
	if count != 43 {
		t.Errorf("wrong value %d", count)
	}

	count = counter.Get()
	if count != 43 {
		t.Errorf("wrong value %d", count)
	}
}
