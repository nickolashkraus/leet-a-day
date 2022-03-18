package main

// 1652. Defuse the Bomb
//
// https://leetcode.com/problems/defuse-the-bomb
//
// NOTES
//   * Use modulo to find the correct index:
//
//       index = index % n
//
//     where n is the length of the array.
//
//     Example:
//
//     [A, B, C ... A, B, C]
//      ^  ^  ^     ^  ^  ^
//      0  1  2     3  4  5
//
//     The 3rd element would be 3 % 3 = 0 or arr[0] = A.
func decrypt(code []int, k int) []int {

	codeLength := len(code)
	// n == code.length and 1 <= n <= 100, therefore allocate an array of size
	// 100 of type int.
	decrypted := [100]int{}

	// With 'range', if you only want the index, you can omit the second
	// variable.
	for curr := range code {
		// If k > 0, replace the ith number with the sum of the next k numbers.
		// If k < 0, replace the ith number with the sum of the previous k numbers.
		// If k == 0, replace the ith number with 0.
		if k > 0 {
			val := 0
			for i := 1; i <= k; i++ {
				val += code[(curr+i)%codeLength]
			}
			decrypted[curr] = val
		} else if k < 0 {
			val := 0
			for i := 1; i <= k*-1; i++ {
				if curr-i < 0 {
					val += code[codeLength-abs(curr-i)%codeLength]
				} else {
					val += code[abs(curr-i)%codeLength]
				}
			}
			decrypted[curr] = val
		} else {
			decrypted[curr] = 0
		}
	}

	// A slice is formed from the array by specifying two indices, a low and high
	// bound, separated by a colon.
	return decrypted[0:codeLength]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
