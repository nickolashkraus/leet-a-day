package pkg

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	k, nums := 3, []int{4, 5, 8, 2}
	kthLargest := Constructor(k, nums)
	ret := kthLargest.Add(3) // return 4
	if ret != 4 {
		t.Fail()
	}
	ret = kthLargest.Add(5) // return 5
	if ret != 5 {
		t.Fail()
	}
	ret = kthLargest.Add(10) // return 5
	if ret != 5 {
		t.Fail()
	}
	ret = kthLargest.Add(9) // return 8
	if ret != 8 {
		t.Fail()
	}
	ret = kthLargest.Add(4) // return 8
	if ret != 8 {
		t.Fail()
	}
}
