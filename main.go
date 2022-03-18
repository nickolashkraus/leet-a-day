package main

import (
	"errors"
	"math"
	"strings"
	"unicode"
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

// 1805. Number of Different Integers in a String
//
// https://leetcode.com/problems/number-of-different-integers-in-a-string
//
// NOTES
//   * Use a hash table for integer lookups.
func NumDifferentIntegers(word string) (int, error) {

	// 1 <= word.length <= 1000
	if len(word) < 1 || len(word) > 1000 {
		return 0, errors.New("Error: 1 <= word.length <= 1000")
	}

	integers := make(map[string]bool)
	runes := []rune{}
	for _, c := range word {
		if unicode.IsDigit(c) {
			runes = append(runes, c)
		} else if len(runes) > 0 {
			integers = insertInterger(integers, runes)
			runes = []rune{}
		}
	}
	if len(runes) > 0 {
		integers = insertInterger(integers, runes)
	}
	return len(integers), nil
}

// Insert integer in map 'integers', if it is not in the map
func insertInterger(integers map[string]bool, runes []rune) map[string]bool {
	integer := strings.TrimLeft(string(runes), "0")
	elem, _ := integers[integer]
	if !elem {
		integers[integer] = true
	}
	return integers
}

// 1652. Defuse the Bomb
//
// https://leetcode.com/problems/defuse-the-bomb
//
// NOTES
//   * Use modulo to find the correct index: i = i % n, where n is the length
//     of the array.
//
//   Example:
//
//   [0, 1, 2 ... 0, 1, 2]
//    ^  ^  ^     ^  ^  ^
//    0  1  2     3  4  5
//
//   The 3rd element would be 3 % 3 = arr[0] = 0.
func Decrypt(code []int, k int) []int {

	// n == code.length and 1 <= n <= 100, therefore allocate an array of size
	// 100 of type int.
	decrypted := [100]int{}

	for curr, _ := range code {
		// If k > 0, replace the ith number with the sum of the next k numbers.
		// If k < 0, replace the ith number with the sum of the previous k numbers.
		// If k == 0, replace the ith number with 0.
		if k > 0 {
			val := 0
			for i := 1; i <= k; i++ {
				val += code[(curr+i)%len(code)]
			}
			decrypted[curr] = val
		} else if k < 0 {
			val := 0
			for i := 1; i <= k*-1; i++ {
				if curr-i < 0 {
					val += code[len(code)-abs(curr-i)%len(code)]
				} else {
					val += code[abs(curr-i)%len(code)]
				}
			}
			decrypted[curr] = val
		} else {
			decrypted[curr] = 0
		}
	}

	return decrypted[0:len(code)]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
}
