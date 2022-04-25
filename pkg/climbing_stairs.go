package pkg

// 70. Climbing Stairs
//
// https://leetcode.com/problems/climbing-stairs
//
// NOTES
//   * Use dynamic programming.
//
//     See the Wikipedia article:
//     https://en.wikipedia.org/wiki/Dynamic_programming
//
//     See the Wikipedia article:
//     https://en.wikipedia.org/wiki/Memoization

// One can reach the ith step in one of two ways:
//   * Taking a step of 1 from the (i−1)th step.
//   * Taking a step of 2 from the (i−2)th step.
//
// So, the total number of ways to reach the ith step is equal to the sum of
// the ways to reach the (i−1)th step and the sum of the ways to reach (i−2)th
// step.
func climbStairs(n int) int {
	// 1 <= n <= 45
	dp := [46]int{}
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
		} else if i == 2 {
			dp[i] = 2
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}
