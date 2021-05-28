//216. Combination Sum III
//https://leetcode.com/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	if k < 1 || k > 9 {
		return [][]int{}
	}

	r, _ := dfs(1, k, n)
	return r
}

func dfs(start, count, sumary int) ([][]int, bool) {
	if sumary <= 0 {
		return nil, false
	}
	if count == 1 {
		if start <= sumary && sumary <= 9 {
			return [][]int{{sumary}}, true
		}
		return nil, false
	}
	result := [][]int{}
	for i := start; i <= 10-count; i++ {
		next, ok := dfs(i+1, count-1, sumary-i)
		if ok {
			for _, n := range next {
				result = append(result, append([]int{i}, n...))
			}
		}
	}
	return result, len(result) > 0
}