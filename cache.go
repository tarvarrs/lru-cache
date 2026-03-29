// Package lru provides a thread‑safe LRU (Least Recently Used) cache implementation with generics. It supports O(1) Get, Set and Clear operations.
package lru

import (
	"sync"
)

// Cache is a thread‑safe LRU cache that stores key‑value pairs up to a fixed capacity.
// When the cache is full, the least recently used entry is evicted automatically.
// The zero value is not usable; use NewCache to create a cache.
type Cache[K comparable, V any] struct {
	mu       sync.Mutex
	start    *node[K, V]
	end      *node[K, V]
	capacity int
	items    map[K]*node[K, V]
}

// Node represents a node in the internal doubly linked list used by the cache.
// It holds a key and a value, and pointers to its neighbours.
type node[K comparable, V any] struct {
	next  *node[K, V]
	prev  *node[K, V]
	key   K
	value V
}

// NewCache creates a new LRU cache with the given capacity.
// The capacity must be positive; if capacity is less than or equal to zero,
// the cache will never store any items.
func NewCache[K comparable, V any](cap int) *Cache[K, V] {
	startNode := &node[K, V]{}
	endNode := &node[K, V]{}
	startNode.prev = endNode
	endNode.next = startNode
	return &Cache[K, V]{
		start:    startNode,
		end:      endNode,
		capacity: cap,
		items:    make(map[K]*node[K, V], cap),
	}
}

// Set adds or updates a key‑value pair in the cache.
// If the key already exists, its value is updated and the entry becomes the most recently used.
// If the cache is at capacity, the least recently used entry is evicted before inserting the new one.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.items[key]
	if ok {
		if item.next != c.start {
			c.moveToFront(item)
			item.value = value
		}
	} else {
		newNode := &node[K, V]{}
		newNode.key = key
		newNode.value = value
		if len(c.items) == c.capacity {
			tmp := c.end.next
			c.end.next.next.prev = c.end
			c.end.next = c.end.next.next
			delete(c.items, tmp.key)
		}
		newNode.next = c.start
		newNode.prev = c.start.prev
		c.start.prev.next = newNode
		c.start.prev = newNode
		c.items[key] = newNode
	}
}

// Get retrieves the value for a given key.
// It returns the value and true if the key exists; otherwise it returns the zero value of V and false.
// When the key exists, the entry is marked as the most recently used.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.items[key]
	if ok {
		if item.next != c.start {
			c.moveToFront(item)
		}
		return item.value, true
	} else {
		var zero V
		return zero, false
	}
}

// Clear removes all entries from the cache, leaving it empty.
// The capacity remains unchanged.
func (c *Cache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.start.prev = c.end
	c.end.next = c.start
	clear(c.items)
}

// moveToFront moves an existing node to the front (most recent position) of the linked list.
// It assumes the mutex is already held by the caller.
func (c *Cache[K, V]) moveToFront(item *node[K, V]) {
	item.prev.next = item.next
	item.next.prev = item.prev
	item.prev = c.start.prev
	item.next = c.start
	c.start.prev.next = item
	c.start.prev = item
}
