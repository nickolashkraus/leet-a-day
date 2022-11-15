package pkg

// 125. Valid Palindrome
//
// https://leetcode.com/problems/valid-palindrome
//
// NOTES
//   - Use two pointers.
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
