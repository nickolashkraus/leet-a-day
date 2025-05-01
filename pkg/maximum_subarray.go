package pkg

import "math"

// 53. Maximum Subarray
//
// https://leetcode.com/problems/maximum-subarray
//
// NOTES
//   * Apply a greedy algorithm.
//
//     See the Wikipedia article:
//     https://en.wikipedia.org/wiki/Greedy_algorithm
//
//     See also:
//     https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm
func maxSubArray(nums []int) int {
	max := math.MinInt // max subarray
	curr := 0          // current subarray
	for _, elem := range nums {
		// Any subarray whose sum is positive is worth keeping.
		if (curr + elem) > 0 {
			curr += elem
		} else {
			curr = 0
		}
		if curr > max {
			max = curr
		}
	}
	return max
}
