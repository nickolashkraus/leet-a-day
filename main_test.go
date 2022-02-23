package main

import "testing"

func TestTitleToNumber(t *testing.T) {
	ret, err := TitleToNumber("A")
	if err != nil {
		t.Log("An error occurred.")
		t.Fail()
	} else if ret != 1 {
		t.Fail()
	}

	ret, err = TitleToNumber("AB")
	if err != nil {
		t.Log("An error occurred.")
		t.Fail()
	} else if ret != 28 {
		t.Fail()
	}

	ret, err = TitleToNumber("ZY")
	if err != nil {
		t.Log("An error occurred.")
		t.Fail()
	} else if ret != 701 {
		t.Fail()
	}

	ret, err = TitleToNumber("FXSHRXW")
	if err != nil {
		t.Log("An error occurred.")
		t.Fail()
	} else if ret != 2147483647 {
		t.Fail()
	}
}
