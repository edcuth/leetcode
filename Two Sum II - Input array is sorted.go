//167. Two Sum II - Input array is sorted
//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	sum := numbers[i] + numbers[j]
	for sum != target {
		if sum > target {
			j--
		} else {
			i++
		}
		sum = numbers[i] + numbers[j]
	}
	return []int{i + 1, j + 1}
}