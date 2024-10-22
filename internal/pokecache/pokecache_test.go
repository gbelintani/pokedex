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
			key: "https://test.com",
			val: []byte("test"),
		},
		{
			key: "https://test2.com",
			val: []byte("test2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(5 * time.Second)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("should have find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("should have same value")
				return
			}
		})
	}

}

func TestReap(t *testing.T) {
	interval := 2 * time.Second
	waitTime := interval + 2*time.Second

	cache := NewCache(interval)
	cache.Add("test", []byte("testtt"))
	_, ok := cache.Get("test")
	if !ok {
		t.Errorf("should have find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("test")
	if ok {
		t.Errorf("should have NOT find key")
		return
	}

}
