//136. Single Number
//https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}