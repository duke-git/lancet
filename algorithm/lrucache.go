package algorithm

type lruNode[K comparable, V any] struct {
	key   K
	value V
	pre   *lruNode[K, V]
	next  *lruNode[K, V]
}

// newLruNode return a lruNode pointer
func newLruNode[K comparable, V any](key K, value V) *lruNode[K, V] {
	return &lruNode[K, V]{
		key:   key,
		value: value,
		pre:   nil,
		next:  nil,
	}
}

// LRUCache lru cache (thread unsafe)
type LRUCache[K comparable, V any] struct {
	cache    map[K]*lruNode[K, V]
	head     *lruNode[K, V]
	tail     *lruNode[K, V]
	capacity int
	length   int
}

// NewLRUCache creates a LRUCache pointer instance.
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		cache:    make(map[K]*lruNode[K, V], capacity),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		length:   0,
	}
}

// Get value of key from lru cache.
// Play: https://go.dev/play/p/iUynEfOP8G0
func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	var value V

	node, ok := l.cache[key]
	if ok {
		l.moveToHead(node)
		return node.value, true
	}

	return value, false
}

// Put value of key into lru cache.
// Play: https://go.dev/play/p/iUynEfOP8G0
func (l *LRUCache[K, V]) Put(key K, value V) {
	node, ok := l.cache[key]
	if !ok {
		newNode := newLruNode(key, value)
		l.cache[key] = newNode
		l.addNode(newNode)

		if len(l.cache) > l.capacity {
			oldKey := l.deleteNode(l.head)
			delete(l.cache, oldKey)
		}
	} else {
		node.value = value
		l.moveToHead(node)
	}
	l.length = len(l.cache)
}

// Delete item from lru cache.
func (l *LRUCache[K, V]) Delete(key K) bool {
	node, ok := l.cache[key]
	if ok {
		key := l.deleteNode(node)
		delete(l.cache, key)
		return true
	}

	return false
}

// Len returns the number of items in the cache.
func (l *LRUCache[K, V]) Len() int {
	return l.length
}

func (l *LRUCache[K, V]) addNode(node *lruNode[K, V]) {
	if l.tail != nil {
		l.tail.next = node
		node.pre = l.tail
		node.next = nil
	}
	l.tail = node
	if l.head == nil {
		l.head = node
	}
}

func (l *LRUCache[K, V]) deleteNode(node *lruNode[K, V]) K {
	if node == l.tail {
		l.tail = l.tail.pre
	} else if node == l.head {
		l.head = l.head.next
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.key
}

func (l *LRUCache[K, V]) moveToHead(node *lruNode[K, V]) {
	if l.tail == node {
		return
	}
	l.deleteNode(node)
	l.addNode(node)
}
