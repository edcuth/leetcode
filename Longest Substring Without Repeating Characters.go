//Longest Substring Without Repeating Characters
//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	mp := make(map[byte]int)
	i := 0
	for j := 0; j < n; j++ {
		if _, ok := mp[s[j]]; ok {
			if mp[s[j]] > i {
				i = mp[s[j]]
			}
		}
		if (j - i + 1) > ans {
			ans = (j - i + 1)
		}
		mp[s[j]] = j + 1
	}
	return ans
}