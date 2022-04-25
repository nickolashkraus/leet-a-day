package pkg

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	// Input
	//
	//   ["LRUCache","put","put","get","put","get","put","get","get","get"]
	//   [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//
	// Output
	//
	//   [null,null,null,1,null,-1,null,-1,3,4]
	cache := LRUCacheConstructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	ret := cache.Get(1)
	if ret != 1 {
		t.Fail()
	}
	cache.Put(3, 3)
	ret = cache.Get(2)
	if ret != -1 {
		t.Fail()
	}
	cache.Put(4, 4)
	ret = cache.Get(1)
	if ret != -1 {
		t.Fail()
	}
	ret = cache.Get(3)
	if ret != 3 {
		t.Fail()
	}
	ret = cache.Get(4)
	if ret != 4 {
		t.Fail()
	}
	// Input
	//
	//   ["LRUCache","put","put","get","put","put","get"]
	//   [[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
	//
	// Output
	//
	//   [null,null,null,2,null,null,-1]
	cache = LRUCacheConstructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	ret = cache.Get(2)
	if ret != 2 {
		t.Fail()
	}
	cache.Put(1, 1)
	cache.Put(4, 1)
	ret = cache.Get(2)
	if ret != -1 {
		t.Fail()
	}
}
