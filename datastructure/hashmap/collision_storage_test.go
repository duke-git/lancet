package datastructure_test

import (
	"fmt"
	"reflect"
	"testing"

	hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

type collisionKey string

func TestHashMapCollisionStorage(t *testing.T) {
	cases := map[string]func(*testing.T){
		"distinct colliding keys": func(t *testing.T) {
			h := hashmap.NewHashMapWithCapacity(0, 1)
			keys := []any{1, "1", collisionKey("1")}
			for i, k := range keys {
				h.Put(k, i+10)
			}
			for i, k := range keys {
				if h.Get(k) != i+10 || !h.Contains(k) {
					t.Fatalf("lost key %T(%v): %v", k, k, h.Get(k))
				}
			}
			if h.Size() != 3 {
				t.Fatalf("size=%d", h.Size())
			}
		},
		"overwrite cardinality": func(t *testing.T) {
			h := hashmap.NewHashMap()
			h.Put("same", 1)
			h.Put("same", 2)
			if h.Size() != 1 || !reflect.DeepEqual(h.Keys(), []any{"same"}) || !reflect.DeepEqual(h.Values(), []any{2}) {
				t.Fatalf("overwrite size/keys/values: %d %v %v", h.Size(), h.Keys(), h.Values())
			}
		},
		"delete absent colliding key": func(t *testing.T) {
			h := hashmap.NewHashMap()
			h.Put(1, "kept")
			h.Delete("1")
			if h.Size() != 1 || h.Get(1) != "kept" {
				t.Fatalf("absent delete removed another key: size=%d get=%v", h.Size(), h.Get(1))
			}
		},
		"delete every chain position": func(t *testing.T) {
			for _, remove := range []any{1, "1", collisionKey("1")} {
				h := hashmap.NewHashMapWithCapacity(0, 4)
				keys := []any{1, "1", collisionKey("1")}
				for i, k := range keys {
					h.Put(k, i)
				}
				h.Delete(remove)
				if h.Contains(remove) || h.Size() != 2 {
					t.Fatalf("deletion failed: %T size=%d", remove, h.Size())
				}
				for i, k := range keys {
					if !reflect.DeepEqual(k, remove) && h.Get(k) != i {
						t.Fatalf("delete lost other key: %T", k)
					}
				}
			}
		},
		"delete preserves bucket indexes": func(t *testing.T) {
			h := hashmap.NewHashMap()
			for i := 0; i < 100; i++ {
				h.Put(fmt.Sprintf("key-%d", i), i)
			}
			for i := 0; i < 100; i += 2 {
				h.Delete(fmt.Sprintf("key-%d", i))
			}
			for i := 1; i < 100; i += 2 {
				if h.Get(fmt.Sprintf("key-%d", i)) != i {
					t.Fatalf("lost untouched bucket key %d", i)
				}
			}
			if h.Size() != 50 {
				t.Fatalf("size=%d", h.Size())
			}
		},
		"resize and enumerate chains": func(t *testing.T) {
			h := hashmap.NewHashMapWithCapacity(0, 1)
			h.Put(1, 1)
			h.Put("1", 2)
			h.Put(collisionKey("1"), 3)
			for i := 2; i <= 64; i++ {
				h.Put(i, i)
			}
			seen := map[string]any{}
			h.Iterate(func(k, v any) { seen[fmt.Sprintf("%T:%v", k, k)] = v })
			if h.Size() != 66 || len(seen) != 66 || len(h.Keys()) != 66 || len(h.Values()) != 66 {
				t.Fatalf("size=%d enumerated=%d", h.Size(), len(seen))
			}
			if seen["string:1"] != 2 || h.Get(collisionKey("1")) != 3 {
				t.Fatal("collision chain lost during growth")
			}
		},
		"filter colliding values": func(t *testing.T) {
			h := hashmap.NewHashMapWithCapacity(0, 4)
			h.Put(1, 1)
			h.Put("1", 2)
			h.Put(collisionKey("1"), 3)
			f := h.FilterByValue(func(v any) bool { return v.(int)%2 == 1 })
			if f == nil || f.Size() != 2 || f.Get(1) != 1 || f.Get(collisionKey("1")) != 3 || f.Contains("1") {
				t.Fatal("filter skipped or corrupted collision chain")
			}
			if h.FilterByValue(func(any) bool { return false }) != nil {
				t.Fatal("empty-filter nil contract changed")
			}
		},
		"zero capacity and deep equal keys": func(t *testing.T) {
			h := hashmap.NewHashMapWithCapacity(0, 0)
			if h.Get("missing") != nil || h.Contains("missing") {
				t.Fatal("empty lookup")
			}
			h.Delete("missing")
			h.Put([]int{1, 2}, "first")
			h.Put([]int{1, 2}, "second")
			if h.Get([]int{1, 2}) != "second" || h.Size() != 1 {
				t.Fatal("deep equal overwrite failed")
			}
			h.Delete([]int{1, 2})
			if h.Size() != 0 {
				t.Fatal("deep equal delete failed")
			}
		},
	}
	for name, fn := range cases {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("public map operation panicked: %v", r)
				}
			}()
			fn(t)
		})
	}
}
