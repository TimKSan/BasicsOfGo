package main

import (
	"fmt"
	//"strings"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработать конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{
		data: make(map[string]string),
	}
}

type cacheImpl struct {
	data map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	val, ok := c.data[k]
	return val, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.data[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "111", "test": "222"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)
	db.cache.Set("привет", "hi")
	db.cache.Set("один", "one")
	fmt.Println(db.Get("привет"))
	fmt.Println(db.Get("один"))
}

