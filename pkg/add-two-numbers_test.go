package pkg

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	h1 := linkedListFromSlice(l1)
	h2 := linkedListFromSlice(l2)
	ret := addTwoNumbers(h1, h2)
	if ret.Val != 7 {
		t.Fail()
	}
	if ret.Next.Val != 0 {
		t.Fail()
	}
	if ret.Next.Next.Val != 8 {
		t.Fail()
	}
}

func TestLinkedListFromSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	// Pointer to the head of the linked list
	head := linkedListFromSlice(list)
	// Check binary tree:
	//
	// [1,2,3,4,5] -> 1->2->3->4->5
	if head.Val != 1 {
		t.Fail()
	}
	if head.Next.Val != 2 {
		t.Fail()
	}
	if head.Next.Next.Val != 3 {
		t.Fail()
	}
	if head.Next.Next.Next.Val != 4 {
		t.Fail()
	}
	if head.Next.Next.Next.Next.Val != 5 {
		t.Fail()
	}
}
