//231. Power of Two
//https://leetcode.com/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if n == 0 {
		return false
	}
	for n%2 == 0 {
		if n == 2 {
			return true
		}
		n = n / 2
	}
	return false
}