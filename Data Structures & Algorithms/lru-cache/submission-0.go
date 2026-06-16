type LRUCache struct {
	cache map[int]*Node
	capacity int
	left, right *Node
}

type Node struct {
	key, val int
	next, prev *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache {
		cache : make(map[int]*Node, 0),
		capacity : capacity,
		left : &Node{},
		right : &Node{},
	}
	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache)insert(node *Node) {
	prev, next := this.right.prev, this.right
	prev.next = node
	node.next = next
	next.prev = node
	node.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key];ok{
		// mark this as more recently used.
		this.remove(node)
		this.insert(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key];ok{
		this.remove(node)
		delete(this.cache, key)
	}
	n := &Node{key:key, val:value}
	this.cache[key]=n
	this.insert(n)

	if len(this.cache)>this.capacity{
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}
