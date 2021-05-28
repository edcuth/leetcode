//96. Unique Binary Search Trees
//https://leetcode.com/problems/unique-binary-search-trees/
func numTrees(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			l := j - 1
			r := i - j
			res[i] += res[l] * res[r]
		}
	}
	fmt.Println(res)
	return res[n]
}