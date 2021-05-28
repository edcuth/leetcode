//209. Minimum Size Subarray Sum
//https://leetcode.com/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		sum := 0
		count := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			count++
			if sum >= target {
				if count == 1 {
					return 1
				} else if count < min {
					min = count
					break
				}
			}
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}