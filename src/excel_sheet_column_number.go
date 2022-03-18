package main

import (
	"math"
)

// 171. Excel Sheet Column Number
//
// https://leetcode.com/problems/excel-sheet-column-number
//
// NOTES
//   * Use ASCII with an offset to determine decimal value.
func titleToNumber(columnTitle string) int {

	// A is 65 in decimal
	offset, columnNumber := 64, 0

	for i, c := range columnTitle {
		if i == len(columnTitle)-1 {
			columnNumber += int(c) - offset
		} else {
			columnNumber += (int(c) - offset) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
		}
	}

	return columnNumber
}
