//34. Find First and Last Position of Element in Sorted Array
//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result[0] = i
			break
		}
	}
	if result[0] == -1 {
		return result
	}
	if nums[len(nums)-1] == target {
		result[1] = len(nums) - 1
		return result
	}
	for i := result[0]; i < len(nums); i++ {
		if nums[i] != target {
			result[1] = i - 1
			break
		}
	}
	return result
}