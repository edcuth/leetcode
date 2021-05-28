//45. Jump Game II
//https://leetcode.com/problems/jump-game-ii/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		return 1
	}
	jumps := 0
	i := 0
	for i <= len(nums)-1 {
		jumps++
		if i+nums[i] >= len(nums)-1 {
			return jumps
		}
		max := 0
		indexmax := 0
		if nums[i] == 1 {
			i++
		} else {
			for j := i + 1; j <= i+nums[i]; j++ {
				if j == len(nums) {
					return jumps
				} else if nums[j]+j > max {
					max = nums[j] + j
					indexmax = j
				}
			}
			fmt.Println(indexmax)
			i = indexmax
		}
	}
	return jumps
}