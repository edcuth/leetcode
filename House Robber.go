//198. House Robber
//https://leetcode.com/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev1, prev2 := 0, 0
	for _, num := range nums {
		temp := prev1
		if (prev2 + num) > prev1 {
			prev1 = prev2 + num
		}
		prev2 = temp
	}
	return prev1
}