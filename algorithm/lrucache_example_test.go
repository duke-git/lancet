package algorithm

import "fmt"

func ExampleLRUCache_Put() {
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	result1, ok1 := cache.Get(1)
	result2, ok2 := cache.Get(2)
	result3, ok3 := cache.Get(3)

	fmt.Println(result1, ok1)
	fmt.Println(result2, ok2)
	fmt.Println(result3, ok3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleLRUCache_Get() {
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	result1, ok1 := cache.Get(1)
	result2, ok2 := cache.Get(2)
	result3, ok3 := cache.Get(3)

	fmt.Println(result1, ok1)
	fmt.Println(result2, ok2)
	fmt.Println(result3, ok3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleLRUCache_Delete() {
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	result1, ok1 := cache.Get(1)

	ok2 := cache.Delete(2)

	_, ok3 := cache.Get(2)

	fmt.Println(result1, ok1)
	fmt.Println(ok2)
	fmt.Println(ok3)

	// Output:
	// 1 true
	// true
	// false
}

func ExampleLRUCache_Len() {
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	result := cache.Len()

	fmt.Println(result)

	// Output:
	// 2
}
