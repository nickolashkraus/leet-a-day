package pkg

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	ret := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(ret, []int{0, 1}) {
		t.Fail()
	}

	ret = twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(ret, []int{1, 2}) {
		t.Fail()
	}

	ret = twoSum([]int{3, 3}, 6)
	if !reflect.DeepEqual(ret, []int{0, 1}) {
		t.Fail()
	}
}
