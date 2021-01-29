package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("A", String("apple"))
	v, ok := lru.Get("A")
	if !ok || string(v.(String)) != "apple" {
		t.Fatal("cache hit A=apple failed")
	}
	if _, ok := lru.Get("B"); ok {
		t.Fatal("B in cache")
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	caps := len(k1 + k2 + v1 + v2)
	lru := New(int64(caps), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get(k1); ok || len(lru.cache) != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}

}

func TestCache_OnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)
	lru.Add("A", String("apple"))
	lru.Add("B", String("banana"))
	lru.Add("C", String("Cat"))
	lru.Add("D", String("Dog"))

	expect := []string{"A", "B"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
