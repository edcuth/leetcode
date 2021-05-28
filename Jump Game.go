//55. Jump Game
//https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if len(nums) == 2 && nums[0] != 0 {
		return true
	}
	i := 0
	for i <= len(nums)-1 {
		if i+nums[i] >= len(nums)-1 {
			return true
		}
		max := 0
		indexmax := 0
		if nums[i] == 0 {
			return false
		} else if nums[i] == 1 {
			i++
		} else {
			for j := i + 1; j <= i+nums[i]; j++ {
				if j == len(nums) {
					return true
				} else if nums[j]+j > max {
					max = nums[j] + j
					indexmax = j
				}
			}
			i = indexmax
		}
	}
	return true
}