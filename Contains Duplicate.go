//217. Contains Duplicate
//https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			return true
		} else {
			mp[n] = true
		}
	}
	return false
}