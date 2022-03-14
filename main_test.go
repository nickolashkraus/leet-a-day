package main

import "testing"

func TestTitleToNumber(t *testing.T) {
	ret, err := TitleToNumber("A")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 1 {
		t.Fail()
	}

	ret, err = TitleToNumber("AB")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 28 {
		t.Fail()
	}

	ret, err = TitleToNumber("ZY")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 701 {
		t.Fail()
	}

	ret, err = TitleToNumber("FXSHRXW")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 2147483647 {
		t.Fail()
	}
}

func TestNumDifferentIntegers(t *testing.T) {
	ret, err := NumDifferentIntegers("a123bc34d8ef34")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 3 {
		t.Fail()
	}

	ret, err = NumDifferentIntegers("leet1234code234")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 2 {
		t.Fail()
	}

	ret, err = NumDifferentIntegers("a1b01c001")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 1 {
		t.Fail()
	}

	ret, err = NumDifferentIntegers("2393706880236110407059624696967828762752651982730115221690437821508229419410771541532394006597463715513741725852432559057224478815116557380260390432211227579663571046845842281704281749571110076974264971989893607137140456254346955633455446057823738757323149856858154529105301197388177242583658641529908583934918768953462557716z97438020429952944646288084173334701047574188936201324845149110176716130267041674438237608038734431519439828191344238609567530399189316846359766256507371240530620697102864238792350289978450509162697068948604722646739174590530336510475061521094503850598453536706982695212493902968251702853203929616930291257062173c79487281900662343830648295410")
	if err != nil {
		t.Logf("An error occurred: %s", err)
		t.Fail()
	} else if ret != 3 {
		t.Fail()
	}
}
