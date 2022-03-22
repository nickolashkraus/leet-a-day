package pkg

// 1. Two Sum
//
// https://leetcode.com/problems/two-sum
//
// NOTES
//   * Use a hash table for complement lookups.
func twoSum(nums []int, target int) []int {

	table := make(map[int]int)

	for i, num := range nums {
		compl, ok := table[target-num]
		// Cannot use 'compl', since its zero value (0) is a valid value.
		if ok {
			return []int{compl, i}
		} else {
			table[num] = i
		}
	}

	return []int{}
}
