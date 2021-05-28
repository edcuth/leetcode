//74. Search a 2D Matrix
//https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0]) - 1
	if matrix[len(matrix)-1][n] < target {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if target >= matrix[i][0] && target <= matrix[i][n] {
			for j := 0; j <= n; j++ {
				fmt.Println(matrix[i][j])
				if target == matrix[i][j] {
					return true
				}
			}
			fmt.Println("first false")
			return false
		}
	}
	fmt.Println("second false")
	return false
}