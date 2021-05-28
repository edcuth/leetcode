//172. Factorial Trailing Zeroes
//https://leetcode.com/problems/factorial-trailing-zeroes/
func trailingZeroes(n int) int {
	var i, res int = 5, 0

	for n >= i {
		res += n / i
		i *= 5
	}

	return res
}