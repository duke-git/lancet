package algorithm

import (
	"strconv"
	"testing"
)

func TestLRUCacheReuseAfterDeletingLastEntry(t *testing.T) {
	cache := NewLRUCache[int, string](1)
	cache.Put(1, "one")
	if !cache.Delete(1) || cache.Len() != 0 {
		t.Fatal("deleting the sole entry must empty the cache")
	}
	cache.Put(2, "two")
	cache.Put(3, "three")
	if cache.Len() != 1 {
		t.Fatalf("Len() = %d, want capacity 1", cache.Len())
	}
	if _, ok := cache.Get(2); ok {
		t.Fatal("the least recently used replacement entry was not evicted")
	}
	if value, ok := cache.Get(3); !ok || value != "three" {
		t.Fatalf("Get(3) = %q, %v", value, ok)
	}
}

func TestLRUCacheEmptyAndReuse(t *testing.T) {
	for _, capacity := range []int{0, 1, 2, 4} {
		t.Run(strconv.Itoa(capacity), func(t *testing.T) {
			defer func() {
				if value := recover(); value != nil {
					t.Errorf("cache reuse panicked: %v", value)
				}
			}()
			cache := NewLRUCache[int, int](capacity)
			for round := 0; round < 4; round++ {
				for i := 0; i < capacity+2; i++ {
					cache.Put(i, round*10+i)
					if cache.Len() > capacity {
						t.Fatalf("Len() = %d exceeds capacity %d", cache.Len(), capacity)
					}
				}
				for i := 0; i < capacity+2; i++ {
					cache.Delete(i)
				}
				if cache.Len() != 0 {
					t.Fatalf("Len() = %d after deleting every key", cache.Len())
				}
			}
		})
	}
}

func TestLRUCacheEndpointDeletionKeepsRecency(t *testing.T) {
	cache := NewLRUCache[int, int](3)
	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	cache.Delete(1)
	cache.Delete(3)
	cache.Put(4, 40)
	cache.Put(5, 50)
	cache.Get(2)
	cache.Put(6, 60)
	if cache.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", cache.Len())
	}
	if _, ok := cache.Get(4); ok {
		t.Fatal("Get(2) should make 4 the least recently used entry")
	}
	for key, want := range map[int]int{2: 20, 5: 50, 6: 60} {
		if got, ok := cache.Get(key); !ok || got != want {
			t.Fatalf("Get(%d) = %d, %v; want %d", key, got, ok, want)
		}
	}
}
