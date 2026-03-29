package lru

import "testing"

func TestCacheEvictsLeastRecentlyUsed(t *testing.T) {
	cache := NewCache[int, string](2)

	cache.Set(1, "one")
	cache.Set(2, "two")

	if v, ok := cache.Get(1); !ok || v != "one" {
		t.Fatalf("expected to get key 1 = one, got %q, ok=%v", v, ok)
	}

	cache.Set(3, "three")

	if _, ok := cache.Get(2); ok {
		t.Fatalf("expected key 2 to be evicted as least recently used")
	}

	if v, ok := cache.Get(1); !ok || v != "one" {
		t.Fatalf("expected key 1 to remain in cache, got %q, ok=%v", v, ok)
	}

	if v, ok := cache.Get(3); !ok || v != "three" {
		t.Fatalf("expected key 3 in cache, got %q, ok=%v", v, ok)
	}
}

func TestCacheEvictionOnCapacityOverflow(t *testing.T) {
	cache := NewCache[int, string](3)

	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "three")

	if _, ok := cache.Get(1); !ok {
		t.Fatalf("expected key 1 to exist")
	}
	if _, ok := cache.Get(3); !ok {
		t.Fatalf("expected key 3 to exist")
	}

	cache.Set(4, "four")

	if _, ok := cache.Get(2); ok {
		t.Fatalf("expected key 2 to be evicted as least recently used after overflow")
	}

	if v, ok := cache.Get(1); !ok || v != "one" {
		t.Fatalf("expected key 1 to remain after overflow, got %q, ok=%v", v, ok)
	}

	if v, ok := cache.Get(3); !ok || v != "three" {
		t.Fatalf("expected key 3 to remain after overflow, got %q, ok=%v", v, ok)
	}

	if v, ok := cache.Get(4); !ok || v != "four" {
		t.Fatalf("expected key 4 to be in cache, got %q, ok=%v", v, ok)
	}
}
