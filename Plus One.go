//66. Plus One
//https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {

	n := len(digits) - 1
	carry := 1 // adding 1 to the number
	for i := n; i >= 0; i-- {
		digit := digits[i]
		if digit == 9 && carry == 1 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] = digits[i] + carry
			carry = 0
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}