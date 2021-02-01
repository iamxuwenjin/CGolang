package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type SingleGroup struct {
	mu   sync.Mutex
	maps map[string]*call
}

func (g *SingleGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.maps == nil {
		g.maps = make(map[string]*call)
	}
	c, ok := g.maps[key]
	if ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	calling := new(call)
	calling.wg.Add(1)
	g.maps[key] = calling
	g.mu.Unlock()

	calling.val, calling.err = fn()
	calling.wg.Done()

	g.mu.Lock()
	delete(g.maps, key)
	g.mu.Unlock()

	return calling.val, calling.err
}
