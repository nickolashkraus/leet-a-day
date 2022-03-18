package main

import (
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	ret := titleToNumber("A")
	if ret != 1 {
		t.Fail()
	}

	ret = titleToNumber("AB")
	if ret != 28 {
		t.Fail()
	}

	ret = titleToNumber("ZY")
	if ret != 701 {
		t.Fail()
	}

	ret = titleToNumber("FXSHRXW")
	if ret != 2147483647 {
		t.Fail()
	}
}
