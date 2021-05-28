//28. Implement strStr()
//https://leetcode.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	l := len(needle)
	for i := 0; i < len(haystack)-l+1; i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}