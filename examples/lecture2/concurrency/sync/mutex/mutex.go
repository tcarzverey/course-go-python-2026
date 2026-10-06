package main

import "sync"

type NotSafeMap interface {
	Get(string) (string, bool)
	Set(string, string)
}

type Cache struct {
	m  NotSafeMap
	mu sync.Mutex
}

func (c Cache) Get(key string) string {
	c.mu.Lock()
	res, _ := c.m.Get(key)
	c.mu.Unlock()

	return res
}

func main() {
	c := Cache{}
	_ = c.Get("abc")
}
