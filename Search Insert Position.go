//35. Search Insert Position
//https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	if target == 0 {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return -1
}