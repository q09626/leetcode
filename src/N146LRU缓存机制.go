package main

type DLNode struct {
	key, value int
	prev, next *DLNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLNode
	head, tail *DLNode
}

func Constructor1(capacity int) LRUCache {
	l := LRUCache{0, capacity, map[int]*DLNode{}, &DLNode{}, &DLNode{}}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key, value int) {
	if _, ok := l.cache[key]; !ok {
		node := &DLNode{key, value, nil, nil}
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next = node
	node.next.prev = node
}

func (l *LRUCache) moveToHead(node *DLNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *DLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) removeTail() *DLNode {
	removed := l.tail.prev
	l.removeNode(removed)
	return removed
}
