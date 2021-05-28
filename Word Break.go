//139. Word Break
//https://leetcode.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	dict := make(map[string]struct{})

	for _, val := range wordDict {
		dict[val] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				substr := s[j:i]
				if _, ok := dict[substr]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(dp)-1]
}