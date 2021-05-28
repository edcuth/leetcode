//81. Search in Rotated Sorted Array II
//https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func search(nums []int, target int) bool {
	check := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		} else if nums[i] < target {
			check = true
		} else if nums[i] > target {
			if check == true {
				return false
			}
		}
	}
	return false
}