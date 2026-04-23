package pokecache

import (
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

	for _, c := range cases {
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

	}
}
