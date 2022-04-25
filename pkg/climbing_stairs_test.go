package pkg

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	ret := climbStairs(2)
	if ret != 2 {
		t.Fail()
	}
	ret = climbStairs(3)
	if ret != 3 {
		t.Fail()
	}
	ret = climbStairs(45)
	if ret != 1836311903 {
		t.Fail()
	}
}
