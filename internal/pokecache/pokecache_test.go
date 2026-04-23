package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "fakeurl.com",
			val: []byte("example value"),
		},
		{
			key: "newfake.com",
			val: []byte("hello"),
		},
	}

	cache := NewCache(1 * time.Second)

	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d_%s", i, c.key), func(t *testing.T) {
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("url %q is not cached", c.key)
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("url %q | Expected: %q | Got Value: %q", c.key, string(c.val), string(val))
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "fakeurl.com",
			val: []byte("example value"),
		},
		{
			key: "newfake.com",
			val: []byte("hello"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case: %d. URL: %q", i, c.key), func(t *testing.T) {
			cache := NewCache(5 * time.Millisecond)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("url %q is not cached", c.key)
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("url %q | Expected: %q | Got Value: %q", c.key, string(c.val), string(val))
				return
			}

			time.Sleep(10 * time.Millisecond)

			_, ok = cache.Get(c.key)
			if ok {
				t.Errorf("url %q is cached", c.key)
				return
			}
		})
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		key string
		ok  bool
	}{
		{
			key: "fakeurl.com",
			ok:  false,
		},
		{
			key: "newfake.com",
			ok:  false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case: %d | url: %q", i, c.key), func(t *testing.T) {
			cache := NewCache(5 * time.Millisecond)

			_, ok := cache.Get(c.key)

			if ok != c.ok {
				t.Errorf("url: %q should not hold a value", c.key)
				return
			}
		})
	}
}
