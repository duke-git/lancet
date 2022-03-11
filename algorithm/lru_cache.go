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

// LRUCache lru cache
type LRUCache[K comparable, V any] struct {
	cache    map[K]*lruNode[K, V]
	head     *lruNode[K, V]
	tail     *lruNode[K, V]
	capacity int
	length   int
}

// NewLRUCache return a LRUCache pointer
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		cache:    make(map[K]*lruNode[K, V]),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		length:   0,
	}
}

// Get value of key from lru cache
func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	var value V

	node, ok := l.cache[key]
	if node == nil || !ok {
		return value, false
	}

	l.moveToHead(node)

	return node.value, true
}

// Put value of key into lru cache
func (l *LRUCache[K, V]) Put(key K, value V) {
	node, ok := l.cache[key]
	if node == nil || !ok {
		newNode := newLruNode(key, value)
		l.cache[key] = newNode
		l.addNode(newNode)

		l.length++
		if l.length > l.capacity {
			tail := l.popTail()
			delete(l.cache, tail.key)
			l.length--
		}
	}

	node.value = value
	l.moveToHead(node)
}

func (l *LRUCache[K, V]) addNode(node *lruNode[K, V]) {
	node.pre = l.head
	node.next = l.head.next

	l.head.next.pre = node
	l.head.next = node
}

func (l *LRUCache[K, V]) deleteNode(node *lruNode[K, V]) {
	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
}

func (l *LRUCache[K, V]) moveToHead(node *lruNode[K, V]) {
	l.deleteNode(node)
	l.addNode(node)
}

func (l *LRUCache[K, V]) popTail() *lruNode[K, V] {
	node := l.tail.pre
	l.deleteNode(node)
	return node
}
