package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var m = map[int]int{}

	for i := range nums {
		complement := target - nums[i]
		if first, ok := m[complement]; ok {
			return []int{first, i}
		}

		m[nums[i]] = i
	}

	return nil
}

// @lc code=end
