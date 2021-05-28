//69. Sqrt(x)
//https://leetcode.com/problems/sqrtx/
func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		} else if i*i == x {
			return i
		}
	}
	return x
}