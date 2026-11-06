package main

import "fmt"

// ---

type EvictionAlgorithm interface {
	evict(c *Cache)
}

// ---

type FIFO struct {
}

func (l *FIFO) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}

// ---

type LRU struct {
}

func (l *LRU) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

// ---

type LFU struct {
}

func (l *LFU) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

// ---

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgorithm
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgorithm) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgorithm) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// ---

func main() {
	cache := initCache(&LFU{})
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	cache.setEvictionAlgo(&LRU{})
	cache.add("d", "4")

	cache.setEvictionAlgo(&FIFO{})
	cache.add("e", "5")
}
