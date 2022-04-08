package pkg

// 2. Add Two Numbers
//
// https://leetcode.com/problems/add-two-numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, prev *ListNode
	n1, n2, carry := l1, l2, 0
	for n1 != nil || n2 != nil {
		v1, v2 := 0, 0
		if n1 != nil {
			v1, n1 = n1.Val, n1.Next
		}
		if n2 != nil {
			v2, n2 = n2.Val, n2.Next
		}
		val := v1 + v2 + carry
		if val >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		curr := &ListNode{
			Val:  val % 10,
			Next: nil,
		}
		if prev != nil {
			prev.Next = curr
		} else {
			head = curr
		}
		prev = curr
	}
	if carry > 0 {
		prev.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head
}

// Given a slice representing a linked list, generate the linked list and
// return a pointer to its head.
func linkedListFromSlice(slice []int) *ListNode {
	var head, prev *ListNode
	for i := 0; i < len(slice); i++ {
		curr := &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		if prev != nil {
			prev.Next = curr
		} else {
			head = curr
		}
		prev = curr
	}

	return head
}
