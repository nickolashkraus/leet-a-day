package pkg

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	ret := maxProfit(prices)
	if ret != 5 {
		t.Fail()
	}
	prices = []int{3, 2, 6, 5, 0, 3}
	ret = maxProfit(prices)
	if ret != 4 {
		t.Fail()
	}
}
