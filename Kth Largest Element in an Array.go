//215. Kth Largest Element in an Array
//https://leetcode.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}