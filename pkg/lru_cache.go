package pkg

import (
	"container/list"
)

// 146. LRU Cache
//
// https://leetcode.com/problems/lru-cache
//
// NOTES
//   * Use a hash map and linked list.
//
//     See the Wikipedia article:
//     https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU
//
//  * The functions get and put must each run in O(1) average time complexity.
//
//    In order to satisfiy this requirement, a combination of a hash map
//    (O(1) time lookups) and a linked list (O(1) time insertions and
//    deletions) is used.
//
//    See also:
//    https://pkg.go.dev/container/list

// Insertions and deletions are O(1) time only if you can instantly access the
// element to be deleted. Since we keep a pointer to the element in a hash map,
// access is in constant time.
type LRUCache struct {
	Capacity int
	Map      map[int]*list.Element
	List     *list.List
}

type Item struct {
	Key int
	Val int
}

// Initialize the LRU cache with positive size capacity.
func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*list.Element),
		List:     list.New(),
	}
}

// Return the value of the key if the key exists, otherwise return -1.
//
// When an element is accessed, move it to the front of the cache (list).
func (this *LRUCache) Get(key int) int {
	elem, ok := this.Map[key]
	if !ok {
		return -1
	}
	val := elem.Value.(*Item).Val
	this.List.MoveToFront(elem)
	return val
}

// Update the value of the key if the key exists. Otherwise, add the key-value
// pair to the cache. If the number of keys exceeds the capacity from this
// operation, evict the least recently used key.
//
// When adding or updating an element, the element is moved to the front of the
// linked list.
func (this *LRUCache) Put(key int, value int) {
	elem, ok := this.Map[key]
	if ok {
		elem.Value.(*Item).Val = value
		this.List.MoveToFront(elem)
	} else {
		elem := this.List.PushFront(&Item{
			Key: key,
			Val: value,
		})
		this.Map[key] = elem
	}
	if this.List.Len() > this.Capacity {
		last := this.List.Back()
		delete(this.Map, last.Value.(*Item).Key)
		this.List.Remove(last)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
