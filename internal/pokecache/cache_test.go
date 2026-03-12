package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(500 * time.Millisecond)
	cache.Add("key", []byte("value"))
	val, ok := cache.Get("key")
	if !ok {
		t.Fatalf("Key does not exist in cache")
	}
	if string(val) != "value" {
		t.Errorf("Expected %q, got %q", "value", string(val))
	}
}

func TestGetMiss(t *testing.T) {
	cache := NewCache(500 * time.Millisecond)
	cache.Add("key", []byte("value"))
	_, ok := cache.Get("ball")
	if ok {
		t.Errorf("Key should not exist but does")
	}

}

func TestReap(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("key", []byte("value"))
	time.Sleep(50 * time.Millisecond)
	_, ok := cache.Get("key")
	if ok {
		t.Errorf("Expected reap to cull keys past duration")
	}

}
