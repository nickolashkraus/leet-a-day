package pkg

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	ret := findDuplicate(nums)
	if ret != 2 {
		t.Fail()
	}
	nums = []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}
	ret = findDuplicate(nums)
	if ret != 9 {
		t.Fail()
	}
	// Index 0 is the start of the cycle.
	nums = []int{3, 1, 3, 4, 2}
	ret = findDuplicate(nums)
	if ret != 3 {
		t.Fail()
	}
}
