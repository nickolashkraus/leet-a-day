package pkg

import (
	"container/heap"
)

// 703. Kth Largest Element in a Stream
//
// https://leetcode.com/problems/kth-largest-element-in-a-stream
//
// NOTES
//   * Use a binary heap (min heap).
//
//     See the Wikipedia article:
//     https://en.wikipedia.org/wiki/Binary_heap
//
//     See also:
//     https://pkg.go.dev/container/heap

// An IntHeap is a min heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	K       int
	IntHeap *IntHeap
}

// Initialize KthLargest object with the integer k (kth largest element) and
// a stream of integers nums.
func KthLargestConstructor(k int, nums []int) KthLargest {
	h := make(IntHeap, len(nums))
	for i, elem := range nums {
		h[i] = elem
	}
	heap.Init(&h)
	for h.Len() > k {
		_ = heap.Pop(&h)
	}
	return KthLargest{
		K:       k,
		IntHeap: &h,
	}
}

// Append the integer val to the stream and return the element representing the
// kth largest element in the stream.
func (this *KthLargest) Add(val int) int {
	k, h := this.K, this.IntHeap
	heap.Push(h, val)
	if h.Len() > k {
		_ = heap.Pop(h)
	}
	return (*h)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
