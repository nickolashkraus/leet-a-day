package main

import (
	"errors"
	"math"
	"strings"
)

// 171. Excel Sheet Column Number
//
// https://leetcode.com/problems/excel-sheet-column-number
//
// NOTES
//   * Use ASCII with an offset to determine decimal value.
func TitleToNumber(columnTitle string) (int, error) {

	// 1 <= columnTitle.length <= 7
	if len(columnTitle) < 1 || len(columnTitle) > 7 {
		return 0, errors.New("Error: 1 <= columnTitle.length <= 7")
	}

	// columnTitle must consist only of uppercase English letters
	if strings.ToUpper(columnTitle) != columnTitle {
		return 0, errors.New("Error: columnTitle must consist only of uppercase English letters.")
	}

	// A is 65 in decimal
	offset, columnNumber := 64, 0

	for i, c := range columnTitle {
		if i == len(columnTitle)-1 {
			columnNumber += int(c) - offset
		} else {
			columnNumber += (int(c) - offset) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
		}
	}

	return columnNumber, nil
}

func main() {}
