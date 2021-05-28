//189. Rotate Array
//https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	if len(nums) != 0 {
		copy(nums, append(nums[len(nums)-k%len(nums):], nums[0:len(nums)-k%len(nums)]...))
	}
}