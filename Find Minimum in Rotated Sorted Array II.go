//154. Find Minimum in Rotated Sorted Array II
//https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] || len(nums) == 1 {
		return nums[0]
	}
	if nums[0] == nums[len(nums)-1] {
		for i := 0; i < len(nums); i++ {
			if nums[i] < nums[0] {
				return nums[i]
			}
		}
	}
	i, j := 0, len(nums)-1
	n := nums[0]
	for j-i != 1 {
		if nums[(i+j)/2] >= n {
			i = (i + j) / 2
		} else if nums[(i+j)/2] < n {
			j = (i + j) / 2
		}
	}
	if nums[i] > nums[j] {
		return nums[j]
	}
	return nums[i]
}