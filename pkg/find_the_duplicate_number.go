package pkg

// 287. Find the Duplicate Number
//
// https://leetcode.com/problems/find-the-duplicate-number
func findDuplicate(nums []int) int {
	tortoise, hare := nums[0], nums[0]
	// In phase 1, hare is twice as fast as tortoise:
	//   * hare = nums[nums[hare]]
	//   * tortoise = nums[tortoise]
	//
	// Since the hare goes faster, it would be the first to enter and run around
	// the cycle. At some point, the tortoise enters the cycle as well, and since
	// it's moving slower the hare catches up to the tortoise at some
	// intersection point.
	for true {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}
	// In phase 2, we give the tortoise a second chance by slowing down the hare,
	// so that it now moves at the speed of tortoise:
	//   * hare = nums[hare]
	//   * tortoise = nums[tortoise]
	//
	// The tortoise is back at the starting position, and the hare starts from
	// the intersection point. The tortoise and the (slowed down) hare will meet
	// at the entrance of the cycle.
	tortoise = nums[0]
	// NOTE: Check that index 0 is not the start of the cycle.
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}
	return tortoise
}
