//169. Majority Element
//https://leetcode.com/problems/majority-element/
func majorityElement(nums []int) int {
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return nums[(len(nums)/2)-1]
	}
	return nums[len(nums)/2]
}