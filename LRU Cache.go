//146. LRU Cache
//https://leetcode.com/problems/lru-cache/
type LRUCache struct {
	Cap  int        //will be compared with the map's len
	Head *CacheNode // pointer to the first node of the list
	Tail *CacheNode // pointer to the last node of the list (the one we'll be replacing when the list is full)
	Mp   map[int]*CacheNode
}

type CacheNode struct {
	// we'll be using these nodes for our linked list, they'll have a pointer to the next and previous node in the list
	Next  *CacheNode
	Prev  *CacheNode
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	// create node to use as our head
	head := CacheNode{
		Prev:  nil,
		Next:  nil,
		Key:   0,
		Value: 0,
	}
	// same for the tail
	tail := CacheNode{
		Prev:  nil,
		Next:  nil,
		Key:   0,
		Value: 0,
	}
	// point the head's next to the tail, and the tail's prev to the head
	head.Next = &tail
	tail.Prev = &head
	// return a LRUCache struct using the head and tail nodes
	return LRUCache{
		Head: &head,
		Tail: &tail,
		Mp:   make(map[int]*CacheNode),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// check if they key is in the map, if so, we remove and insert the node in the linked list again to make it our head
	if n, ok := this.Mp[key]; ok {
		this.remove(n)
		this.insert(n)
		return n.Value
	}
	// else we return -1
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if the key is in the map we remove it
	if _, ok := this.Mp[key]; ok {
		this.remove(this.Mp[key])
	}
	// we check if the len of the map is at the cap, in which case we remove the tail's prev (tail is our made up node)
	if len(this.Mp) == this.Cap {
		this.remove(this.Tail.Prev)
	}
	// then we insert a new node using the helper method
	this.insert(&CacheNode{
		Prev:  nil,
		Next:  nil,
		Key:   key,
		Value: value,
	})
}

func (this *LRUCache) remove(node *CacheNode) {
	// helper method, removes the key from the map, and points node.Prev.Next to node.Next, and node.Next.Prev to node.Prev
	// making the node before and after it point to each other
	delete(this.Mp, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) insert(node *CacheNode) {
	// helper method to insert nodes
	// we add the key and the node to our map
	this.Mp[node.Key] = node
	// then we make a pointer to the head.Next
	next := this.Head.Next
	// then we point the head.Next to our new node
	this.Head.Next = node
	// we point the new node.Prev to the head
	node.Prev = this.Head
	// finally, we point the previous next to node.Next, and next.Prev to the new inserted node
	next.Prev = node
	node.Next = next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */