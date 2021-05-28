//140. Word Break II
//https://leetcode.com/problems/word-break-ii/
func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	res := make([]string, 0)
	dp := make([][]string, n)
	m := make(map[string]bool)

	for i := 0; i < n; i++ {
		dp[i] = make([]string, 0)
	}

	for _, word := range wordDict {
		m[word] = true
	}

	for r := 0; r < n; r++ {
		for l := 0; l <= r; l++ {
			cur := s[l : r+1]
			if m[cur] {
				if l < 1 || len(dp[l-1]) > 0 {
					dp[r] = append(dp[r], cur)
				}
			}
		}
	}
	if len(dp[n-1]) == 0 {
		return []string{}

	}
	dfs(s, dp, "", n-1, &res)

	return res
}

func dfs(s string, dp [][]string, cur string, index int, res *[]string) {
	if index < 0 {
		*res = append(*res, strings.TrimSpace(cur))
		return
	}

	for i := 0; i < len(dp[index]); i++ {
		nextIndex := index - len(dp[index][i])
		next := dp[index][i] + " " + cur

		dfs(s, dp, next, nextIndex, res)

	}
}