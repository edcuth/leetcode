//1. Two Sum
//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            if target - nums[i] == nums[j] {
                return arr := [i, j]
            }
        }
    }
    return [0]
}