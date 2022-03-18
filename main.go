package main

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

	for curr := range code {
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

func main() {}
