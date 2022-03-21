package pkg

import (
	"strings"
	"unicode"
)

// 1805. Number of Different Integers in a String
//
// https://leetcode.com/problems/number-of-different-integers-in-a-string
//
// NOTES
//   * Use a hash table for integer lookups.
func numDifferentIntegers(word string) int {

	integers := make(map[string]bool)
	runes := []rune{}

	for _, c := range word {
		if unicode.IsDigit(c) {
			runes = append(runes, c)
		} else if len(runes) > 0 {
			integers = insertInteger(integers, runes)
			runes = []rune{}
		}
	}

	if len(runes) > 0 {
		integers = insertInteger(integers, runes)
	}

	return len(integers)
}

// Insert integer (as string) in map 'integers', if it is not in the map
func insertInteger(integers map[string]bool, runes []rune) map[string]bool {

	integer := strings.TrimLeft(string(runes), "0")

	// Test that a key is present with a two-value assignment:
	//
	//   elem, ok = m[key]
	//
	// If key is in m, ok is true. If not, ok is false.
	//
	// If key is not in the map, then elem is the zero value for the map's
	// element type.
	elem, _ := integers[integer]
	if !elem {
		integers[integer] = true
	}

	return integers
}
