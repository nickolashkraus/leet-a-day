package main

import (
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	// Array values are comparable if values of the array element type are
	// comparable. Two array values are equal if their corresponding elements are
	// equal.
	//
	// Slice, map, and function values are not comparable. However, as a special
	// case, a slice, map, or function value may be compared to the predeclared
	// identifier nil. Comparison of pointer, channel, and interface values to
	// nil is also allowed and follows from the general rules above.
	//
	// https://go.dev/ref/spec#Comparison_operators
	ret := Decrypt([]int{5, 7, 1, 4}, 3)
	if !reflect.DeepEqual(ret, []int{12, 10, 16, 13}) {
		t.Fail()
	}

	ret = Decrypt([]int{1, 2, 3, 4}, 0)
	if !reflect.DeepEqual(ret, []int{0, 0, 0, 0}) {
		t.Fail()
	}

	ret = Decrypt([]int{2, 4, 9, 3}, -2)
	if !reflect.DeepEqual(ret, []int{12, 5, 6, 13}) {
		t.Fail()
	}
}
