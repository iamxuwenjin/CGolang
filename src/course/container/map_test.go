package container

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]int)
	mapCreated["a"] = 100
	mapCreated["b"] = 90
	fmt.Println(mapLit, mapCreated)
	for index, value := range mapCreated {
		fmt.Println(index, value)
	}

	// 并发安全map
	var scene sync.Map
	scene.Store("tony", 100)
	scene.Store("alice", 22)
	scene.Store("bob", 99)

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
