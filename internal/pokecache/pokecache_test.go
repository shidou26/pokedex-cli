package pokecache

import (
	"slices"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(1 * time.Millisecond)

	c.Add("hello", []byte{100, 200, 255})

	bytes, ok := c.Get("hello")
	if !ok || !slices.Equal(bytes, []byte{100, 200, 255}) {
		t.Errorf("bytes don't matched")
	}

	time.Sleep(5 * time.Millisecond)
	bytes, ok = c.Get("hello")
	if ok {
		t.Errorf("bytes don't erased")	
	}

	bytes, ok = c.Get("world")
	if ok {
		t.Errorf("bytes don't exist")
	}

	c.Add("hello", []byte{0, 0, 0})
	
	bytes, ok = c.Get("hello")
	if !ok || slices.Equal(bytes, []byte{100, 200, 255}) {
		t.Errorf("bytes don't overlap")
	}
}
